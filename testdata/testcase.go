package testdata

import "google.golang.org/protobuf/proto"

type TestCase struct {
	Name string
	V1   any
	V2   proto.Message
}

var (
	TestCaseBasic    []TestCase
	TestCaseRecord   []TestCase
	TestCaseRepeated []TestCase
)
