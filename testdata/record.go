package testdata

import "time"

type Record struct {
	Timestamp time.Time `bigquery:"ts"`
	ID        string    `bigquery:"id"`
	City      *City     `bigquery:"city,nullable"`
}

type City struct {
	Name   string `bigquery:"name"`
	Code   int    `bigquery:"code"`
	Office Office `bigquery:"office"`
}

type Office struct {
	Tel     string `bigquery:"tel"`
	Address string `bigquery:"address"`
}

func init() {
	now := time.Now()

	TestCaseRecord = []TestCase{
		{
			Name: "record",
			V1: &Record{
				Timestamp: now,
				ID:        "record-1",
				City: &City{
					Name: "練馬",
					Code: 131202,
					Office: Office{
						Tel:     "03-3993-1111",
						Address: "練馬区豊玉北6-12-1",
					},
				},
			},
			V2: &RecordProto{
				Ts: now.Unix() * 1000000,
				Id: "record-1",
				City: &CityProto{
					Name: "練馬",
					Code: 131202,
					Office: &OfficeProto{
						Tel:     "03-3993-1111",
						Address: "練馬区豊玉北6-12-1",
					},
				},
			},
		},
		{
			Name: "null",
			V1: &Record{
				Timestamp: now,
				ID:        "record-1",
			},
			V2: &RecordProto{
				Ts: now.Unix() * 1000000,
				Id: "record-1",
			},
		},
	}
}
