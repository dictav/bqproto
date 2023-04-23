# BQProto

bqproto is the library for encoding Go Struct to Protocol Buffers message without a definition.

## Usage

```go
var schema bigquery.Schema

type Row struct {
	Str string  `bigquery:"str"`
	Int integer `bigquery:"int"`
}

func example() {
	row := Row {
		Str: "日本語",
		Int: 123,
	}
	
	b1, err := bqproto.MarshalForSchema(schema, &row)
	if err != nil {
		log.Fatal(err)
	}
}
```
