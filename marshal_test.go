package bqproto_test

import (
	"bytes"
	"testing"

	"github.com/dictav/bqproto"
	"github.com/dictav/bqproto/testdata"
	"google.golang.org/protobuf/proto"
)

func TestMarshal_MarshalForSchema_Basic(t *testing.T) {
	t.Parallel()

	for _, tc := range testdata.TestCaseBasic {
		tc := tc

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			testMarshalForSchema(t, tc, "basic")
		})
	}
}

func TestMarshal_MarshalForSchema_Record(t *testing.T) {
	t.Parallel()

	for _, tc := range testdata.TestCaseRecord {
		tc := tc

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			testMarshalForSchema(t, tc, "record")
		})
	}
}

func TestMarshal_MarshalForSchema_Repeated(t *testing.T) {
	t.Parallel()

	for _, tc := range testdata.TestCaseRepeated {
		tc := tc

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			testMarshalForSchema(t, tc, "repeated")
		})
	}
}

func testMarshalForSchema(t *testing.T, tc testdata.TestCase, schemaName string) {
	t.Helper()

	schema, err := schemaFromJSON(schemaName)
	if err != nil {
		t.Fatal(err)
	}

	got, err := bqproto.MarshalForSchema(schema, &tc.V1)
	if err != nil {
		t.Fatal(err)
	}

	want, err := proto.Marshal(tc.V2)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(want, got) {
		t.Errorf("\nwant %v\n got %v", want, got)
	}
}
