package testdata

import (
	"math"
	"time"
)

type Repeated struct {
	TimestampList []time.Time      `bigquery:"timestamp_list"`
	StrList       []string         `bigquery:"str_list"`
	BooleanList   []bool           `bigquery:"boolean_list"`
	IntegerList   []int64          `bigquery:"integer_list"`
	FloatList     []float64        `bigquery:"float_list"`
	RecordList    []RepeatedRecord `bigquery:"record_list"`
}

type RepeatedRecord struct {
	Len         int64   `bigquery:"len"`
	IntegerList []int64 `bigquery:"integer_list"`
}

func init() {
	now := time.Now()

	TestCaseRepeated = []TestCase{
		{
			Name: "repeated",
			V1: &Repeated{
				TimestampList: []time.Time{now},
				StrList: []string{
					"repeated-1",
					"repeated-2",
				},
				BooleanList: []bool{true, false},
				IntegerList: []int64{1, 2, 3},
				FloatList:   []float64{0.1, 0.2, 0.3},
				RecordList: []RepeatedRecord{
					{Len: 1, IntegerList: []int64{math.MinInt64}},
					{Len: 3, IntegerList: []int64{1, math.MaxInt64, 3}},
				},
			},
			V2: &RepeatedProto{
				TimestampList: []int64{now.Unix() * 1000000},
				StrList: []string{
					"repeated-1",
					"repeated-2",
				},
				BooleanList: []bool{true, false},
				IntegerList: []int64{1, 2, 3},
				FloatList:   []float64{0.1, 0.2, 0.3},
				RecordList: []*RepeatedRecordProto{
					{Len: 1, IntegerList: []int64{math.MinInt64}},
					{Len: 3, IntegerList: []int64{1, math.MaxInt64, 3}},
				},
			},
		},
	}
}
