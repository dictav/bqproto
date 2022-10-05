package bqproto_test

import (
	"errors"
	"fmt"
	"os"
	"testing"

	"cloud.google.com/go/bigquery"
	"github.com/dictav/bqproto"
	"github.com/dictav/bqproto/testdata"
)

func TestSchema_ValidateSchemaCompativility(t *testing.T) {
	t.Parallel()

	var empty testdata.Basic

	testCases := []struct {
		name string
		in   any
		out  error
	}{
		{
			name: "basic",
			in:   empty,
			out:  nil,
		},
		{
			name: "lack",
			in: struct {
				Str string `bigquery:"str"`
			}{},
			out: nil,
		},
		{
			name: "missing requires",
			in: struct {
				Boolean bool `bigquery:"boolean"`
			}{},
			out: bqproto.ErrMissingRequired,
		},
		{
			name: "hindrance",
			in: struct {
				Str  string `bigquery:"str"`
				Str2 string `bigquery:"str2"`
			}{},
			out: bqproto.ErrExistsHindrance,
		},
	}

	schema, err := schemaFromJSON("basic")
	if err != nil {
		t.Fatal(err)
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := bqproto.ValidateSchemaCompatibility(schema, tc.in)
			if !errors.Is(got, tc.out) {
				t.Errorf("want %v, got %v", tc.out, got)
			}
		})
	}
}

func schemaFromJSON(key string) (bigquery.Schema, error) {
	file := fmt.Sprintf("testdata/%s_schema.json", key)

	json, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	schema, err := bigquery.SchemaFromJSON(json)
	if err != nil {
		return nil, err
	}

	return schema, nil
}
