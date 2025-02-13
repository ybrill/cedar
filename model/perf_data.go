package model

import (
	"context"
	"time"

	"github.com/evergreen-ci/cedar"
	"github.com/mongodb/anser/bsonutil"
	"github.com/mongodb/grip"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const DefaultVer = 2

type MetricType string

const (
	MetricTypeMean         MetricType = "mean"
	MetricTypeMedian       MetricType = "median"
	MetricTypeMax          MetricType = "max"
	MetricTypeMin          MetricType = "min"
	MetricTypeSum          MetricType = "sum"
	MetricTypeStdDev       MetricType = "standard-deviation"
	MetricTypePercentile99 MetricType = "percentile-99th"
	MetricTypePercentile90 MetricType = "percentile-90th"
	MetricTypePercentile95 MetricType = "percentile-95th"
	MetricTypePercentile80 MetricType = "percentile-80th"
	MetricTypePercentile50 MetricType = "percentile-50th"
	MetricTypeThroughput   MetricType = "throughput"
	MetricTypeLatency      MetricType = "latency"
)

func (t MetricType) Validate() error {
	switch t {
	case MetricTypeMax, MetricTypeMean, MetricTypeMedian, MetricTypeMin, MetricTypeStdDev:
		return nil
	case MetricTypePercentile50, MetricTypePercentile80, MetricTypePercentile95, MetricTypePercentile99:
		return nil
	default:
		return errors.Errorf("'%s' is not a valid metric type", t)
	}
}

// PerfRollupValue describes a single "rollup", see PerfRollups for more
// information.
type PerfRollupValue struct {
	Name          string      `bson:"name"`
	Value         interface{} `bson:"val"`
	Version       int         `bson:"version"`
	MetricType    MetricType  `bson:"type"`
	UserSubmitted bool        `bson:"user"`
}

var (
	perfRollupValueNameKey          = bsonutil.MustHaveTag(PerfRollupValue{}, "Name")
	perfRollupValueValueKey         = bsonutil.MustHaveTag(PerfRollupValue{}, "Value")
	perfRollupValueVersionKey       = bsonutil.MustHaveTag(PerfRollupValue{}, "Version")
	perfRollupValueMetricTypeKey    = bsonutil.MustHaveTag(PerfRollupValue{}, "MetricType")
	perfRollupValueUserSubmittedKey = bsonutil.MustHaveTag(PerfRollupValue{}, "UserSubmitted")
)

// PerfRollups describes the "rolled up", or calculated metrics from time
// series data collected in a given performance test, of a performance result.
type PerfRollups struct {
	Stats       []PerfRollupValue `bson:"stats"`
	ProcessedAt time.Time         `bson:"processed_at"`

	dirty bool // nolint
	id    string
	env   cedar.Environment
}

var (
	perfRollupsStatsKey       = bsonutil.MustHaveTag(PerfRollups{}, "Stats")
	perfRollupsProcessedAtKey = bsonutil.MustHaveTag(PerfRollups{}, "ProcessedAt")
)

func (v *PerfRollupValue) getIntLong() (int64, error) {
	if val, ok := v.Value.(int64); ok {
		return val, nil
	} else if val, ok := v.Value.(int32); ok {
		return int64(val), nil
	} else if val, ok := v.Value.(int); ok {
		return int64(val), nil
	}
	return 0, errors.Errorf("mismatched type for name %s", v.Name)
}

func (v *PerfRollupValue) getFloat() (float64, error) {
	if val, ok := v.Value.(float64); ok {
		return val, nil
	}
	return 0, errors.Errorf("mismatched type for name %s", v.Name)
}

func (r *PerfRollups) Setup(env cedar.Environment) {
	r.env = env
}

// Add attempts to append a rollup to an existing set of rollups in a
// performance result.
func (r *PerfRollups) Add(ctx context.Context, rollup PerfRollupValue) error {
	if r.id == "" {
		return errors.New("rollups missing id")
	}
	if r.env == nil {
		return errors.New("cannot add rollups with a nil env")

	}

	collection := r.env.GetDB().Collection(perfResultCollection)
	processedAt := time.Now()
	updated, err := tryUpdate(ctx, collection, r.id, rollup, processedAt)
	if !updated && err == nil {
		_, err = collection.UpdateOne(ctx,
			bson.M{perfIDKey: r.id},
			bson.M{
				"$push": bson.M{
					bsonutil.GetDottedKeyName(perfRollupsKey, perfRollupsStatsKey): rollup,
				},
				"$set": bson.M{
					bsonutil.GetDottedKeyName(perfRollupsKey, perfRollupsProcessedAtKey): processedAt,
				},
			},
		)
	}
	if err != nil {
		return errors.Wrap(err, "problem adding rollup")
	}

	for i := range r.Stats {
		if r.Stats[i].Name == rollup.Name {
			r.Stats[i].Version = rollup.Version
			r.Stats[i].Value = rollup.Value
			r.Stats[i].UserSubmitted = rollup.UserSubmitted
			r.Stats[i].MetricType = rollup.MetricType
			return nil
		}
	}
	r.Stats = append(r.Stats, rollup)
	r.ProcessedAt = processedAt
	return nil
}

func tryUpdate(ctx context.Context, collection *mongo.Collection, id string, r PerfRollupValue, processedAt time.Time) (bool, error) {
	updateResult, err := collection.UpdateOne(ctx,
		bson.M{
			perfIDKey: id,
			bsonutil.GetDottedKeyName(perfRollupsKey, perfRollupsStatsKey, perfRollupValueNameKey): r.Name,
		},

		bson.M{
			"$set": bson.M{
				bsonutil.GetDottedKeyName(perfRollupsKey, perfRollupsStatsKey, "$[elem]"): r,
				bsonutil.GetDottedKeyName(perfRollupsKey, perfRollupsProcessedAtKey):      processedAt,
			},
		}, options.Update().SetArrayFilters(options.ArrayFilters{
			Filters: []interface{}{
				bson.M{
					"$and": []bson.M{
						{
							bsonutil.GetDottedKeyName("elem", perfRollupValueNameKey): bson.M{"$eq": r.Name},
						},
						{
							bsonutil.GetDottedKeyName("elem", perfRollupValueVersionKey): bson.M{"$lte": r.Version},
						},
					},
				},
			},
		}),
	)
	if err != nil {
		return false, errors.WithStack(err)
	}

	return updateResult.MatchedCount == 1, errors.WithStack(err)
}

func (r *PerfRollups) GetInt(name string) (int, error) {
	for _, rollup := range r.Stats {
		if rollup.Name == name {
			if val, ok := rollup.Value.(int); ok {
				return val, nil
			} else if val, ok := rollup.Value.(int32); ok {
				return int(val), nil
			} else {
				return 0, errors.Errorf("mismatched type for name %s", name)
			}
		}
	}
	return 0, errors.Errorf("name %s does not exist", name)
}

func (r *PerfRollups) GetInt32(name string) (int32, error) {
	val, err := r.GetInt(name)
	return int32(val), err
}

func (r *PerfRollups) GetInt64(name string) (int64, error) {
	for _, rollup := range r.Stats {
		if rollup.Name == name {
			return rollup.getIntLong()
		}
	}
	return 0, errors.Errorf("name %s does not exist", name)
}

func (r *PerfRollups) GetFloat(name string) (float64, error) {
	for _, rollup := range r.Stats {
		if rollup.Name == name {
			return rollup.getFloat()
		}
	}
	return 0, errors.Errorf("name %s does not exist", name)
}

func (r *PerfRollups) Map() map[string]int64 {
	result := make(map[string]int64)
	for _, rollup := range r.Stats {
		val, err := rollup.getIntLong()
		if err == nil {
			result[rollup.Name] = val
		}
	}
	return result
}

func (r *PerfRollups) MapFloat() map[string]float64 {
	result := make(map[string]float64)
	for _, rollup := range r.Stats {
		if val, err := rollup.getFloat(); err == nil {
			result[rollup.Name] = val
		} else if val, err := rollup.getIntLong(); err == nil {
			result[rollup.Name] = float64(val)
		}
	}
	return result
}

// MergeRollups merges rollups to existing rollups in a performance result. The
// environment should not be nil.
func (r *PerformanceResult) MergeRollups(ctx context.Context, rollups []PerfRollupValue) error {
	if r.env == nil {
		return errors.New("cannot merge rollups with a nil env")
	}

	catcher := grip.NewBasicCatcher()

	r.Rollups.id = r.ID
	r.Rollups.Setup(r.env)

	for _, rollup := range rollups {
		catcher.Add(r.Rollups.Add(ctx, rollup))
	}

	return catcher.Resolve()
}
