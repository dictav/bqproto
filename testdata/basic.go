package testdata

import (
	"math"
	"math/big"
	"time"
)

type Basic struct {
	Str        string    `bigquery:"str"`
	IgnoreStr  string    `bigquery:"-"`
	Boolean    bool      `bigquery:"boolean"`
	Integer    int64     `bigquery:"integer"`
	Float      float64   `bigquery:"float"`
	Binary     []byte    `bigquery:"binary"`
	Date       int32     `bigquery:"date"`
	Timestamp1 int64     `bigquery:"timestamp1"`
	Timestamp2 time.Time `bigquery:"timestamp2"`
	Decimal    big.Rat   `bigquery:"decimal"`
	Time       time.Time `bigquery:"time"`
	DateTime   time.Time `bigquery:"datetime"`
}

func init() {
	now := time.Now()
	datetime := time.Date(1, 1, 1, 0, 0, 1, 0, time.UTC)
	datetimen := uint64(74904230690816) //0001-01-01T00:00:01

	_ = datetime
	_ = datetimen

	TestCaseBasic = []TestCase{
		{
			Name: "basic",
			V1: &Basic{
				Str:        "日本語",
				Boolean:    true,
				Integer:    math.MaxInt64,
				Float:      math.Pi,
				Binary:     []byte("こんにちは、世界"),
				Date:       11016,
				Timestamp1: now.Unix(),
				Timestamp2: now,
				Decimal:    *big.NewRat(-1, 1000_000_000),
				Time:       datetime,
				DateTime:   datetime,
			},
			V2: &BasicProto{
				Str:        "日本語",
				Boolean:    true,
				Integer:    math.MaxInt64,
				Float:      math.Pi,
				Binary:     []byte("こんにちは、世界"),
				Date:       11016,
				Timestamp1: now.Unix() * 1000000,
				Timestamp2: now.Unix() * 1000000,
				Decimal:    []byte{255},
				Time:       datetimen & 0xFFFFFFFFF,
				Datetime:   datetimen,
			},
		},
	}
}
