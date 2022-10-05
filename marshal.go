package bqproto

import (
	"fmt"
	"math"
	"math/big"
	"reflect"
	"time"

	"cloud.google.com/go/bigquery"
	"google.golang.org/protobuf/encoding/protowire"
)

const (
	buildStructSize = 256 * 1024
	buildValueSize  = 128
	buildlSliceSize = 10 * buildValueSize
)

var (
	typeBytes = reflect.TypeOf([]byte(nil))
	typeTime  = reflect.TypeOf(time.Time{})
	typeRat   = reflect.TypeOf(big.Rat{})
)

func MarshalForSchema(schema bigquery.Schema, v any, opts ...Option) ([]byte, error) {
	opt := newOption(opts...)

	for _, f := range opts {
		f(&opt)
	}

	if m, ok := v.(map[string]any); ok {
		return opt.marshalMapForSchema(schema, m)
	}

	return opt.marshalStructForSchema(schema, -1, reflect.ValueOf(v))
}

func (opt *option) marshalMapForSchema(schema bigquery.Schema, v map[string]any) ([]byte, error) {
	return nil, fmt.Errorf("not implemented")
}

//nolint:gocognit
func (opt *option) marshalStructForSchema(schema bigquery.Schema, protonum protowire.Number, value reflect.Value) ([]byte, error) {
	for value.Kind() == reflect.Pointer || value.Kind() == reflect.Interface {
		if value.IsNil() {
			return nil, fmt.Errorf("struct is nil")
		}

		value = value.Elem()
	}

	if value.Kind() != reflect.Struct {
		return nil, fmt.Errorf("should be struct type: %s", value.Kind())
	}

	if value.IsZero() {
		return nil, fmt.Errorf("struct is empty")
	}

	tags := detectTags(value.Type(), opt.tagKey)
	buf := make([]byte, 0, buildStructSize)

	for i, fieldSchema := range schema {
		n := protowire.Number(i + 1)
		t, ok := tags[fieldSchema.Name]

		if !ok {
			if fieldSchema.Required {
				return nil, fmt.Errorf("%w: %q", ErrMissingRequired, fieldSchema.Name)
			}

			continue
		}

		var (
			v   = value.Field(t.index)
			b   []byte
			err error
		)

		switch v.Kind() { //nolint:exhaustive
		case reflect.Interface, reflect.Map, reflect.Pointer:
			if v.IsNil() {
				if fieldSchema.Required {
					return nil, fmt.Errorf("%w: %q", ErrMissingRequired, fieldSchema.Name)
				}

				continue
			}
		}

		for v.Kind() == reflect.Pointer || v.Kind() == reflect.Interface {
			v = v.Elem()
		}

		switch {
		case fieldSchema.Repeated:
			b, err = opt.marshalSliceForFieldSchema(fieldSchema, n, v)
		case fieldSchema.Type == bigquery.RecordFieldType:
			b, err = opt.marshalStructForSchema(fieldSchema.Schema, n, v)
		default:
			b, err = opt.marshalValueForFieldSchema(fieldSchema.Type, n, v, false)
		}

		if err != nil {
			return nil, fmt.Errorf("marshal error for %q: %w", fieldSchema.Name, err)
		}

		buf = append(buf, b...)
	}

	if protonum >= 0 {
		b := make([]byte, 0, 2)
		b = protowire.AppendTag(b, protonum, protowire.BytesType)
		b = protowire.AppendVarint(b, uint64(len(buf)))
		buf = append(b, buf...)
	}

	return buf, nil
}

//nolint:gocognit
func (opt *option) marshalValueForFieldSchema(fieldType bigquery.FieldType, n protowire.Number, v reflect.Value, packing bool) ([]byte, error) {
	buf := make([]byte, 0, buildValueSize)

	switch fieldType {
	case bigquery.StringFieldType:
		if packing {
			return nil, fmt.Errorf("cannot pack for string")
		}

		if v.Kind() != reflect.String {
			return nil, fmt.Errorf("should be string type: %s", v.Type())
		}

		buf = protowire.AppendTag(buf, n, protowire.BytesType)
		buf = protowire.AppendString(buf, v.String())

	case bigquery.BytesFieldType:
		if packing {
			return nil, fmt.Errorf("cannot pack for bytes")
		}

		if v.Type() != typeBytes {
			return nil, fmt.Errorf("should be []byte type: %s", v.Type())
		}

		v.Bytes()

		buf = protowire.AppendTag(buf, n, protowire.BytesType)
		buf = protowire.AppendBytes(buf, v.Bytes())

	case bigquery.IntegerFieldType:
		var num uint64

		switch {
		case v.CanInt():
			num = uint64(v.Int())
		case v.CanUint():
			num = v.Uint()
		default:
			return nil, fmt.Errorf("should be integer type: %s", v.Type())
		}

		if !packing {
			buf = protowire.AppendTag(buf, n, protowire.VarintType)
		}

		buf = protowire.AppendVarint(buf, num)

	case bigquery.FloatFieldType:
		if !v.CanFloat() {
			return nil, fmt.Errorf("should be float type: %s", v.Type())
		}

		if !packing {
			buf = protowire.AppendTag(buf, n, protowire.Fixed64Type)
		}

		buf = protowire.AppendFixed64(buf, math.Float64bits(v.Float()))

	case bigquery.BooleanFieldType:
		if v.Kind() != reflect.Bool {
			return nil, fmt.Errorf("should be bool type: %s", v.Type())
		}

		if !packing {
			buf = protowire.AppendTag(buf, n, protowire.VarintType)
		}

		buf = protowire.AppendVarint(buf, protowire.EncodeBool(v.Bool()))

	case bigquery.TimestampFieldType:
		ts := uint64(0)

		switch {
		case v.Type() == typeTime:
			ts = uint64(v.Interface().(time.Time).Unix())
		case v.CanUint():
			ts = v.Uint()
		case v.CanInt():
			ts = uint64(v.Int())
		default:
			return nil, fmt.Errorf("should be time.Time, int or uint type: %s", v.Type())
		}

		if !packing {
			buf = protowire.AppendTag(buf, n, protowire.VarintType)
		}

		buf = protowire.AppendVarint(buf, ts*1000000)

	case bigquery.NumericFieldType:
		if packing {
			return nil, fmt.Errorf("cannot pack for numeric")
		}

		if v.Type() != typeRat {
			return nil, fmt.Errorf("should be big.Rat type: %s", v.Type())
		}

		rat := v.Interface().(big.Rat)

		if rat.Cmp(&MaxDecimal) == 1 {
			return nil, fmt.Errorf("too large value: %s", &rat)
		}

		if rat.Cmp(&MinDecimal) == -1 {
			return nil, fmt.Errorf("too small value: %s", &rat)
		}

		buf = protowire.AppendTag(buf, n, protowire.BytesType)
		buf = protowire.AppendBytes(buf, encodeDecimal(&rat))

	case bigquery.DateFieldType:
		if !v.CanInt() {
			return nil, fmt.Errorf("should be int type: %s", v.Type())
		}

		if !packing {
			buf = protowire.AppendTag(buf, n, protowire.VarintType)
		}

		buf = protowire.AppendVarint(buf, uint64(v.Int()))

	case bigquery.TimeFieldType:
		if v.Type() != typeTime {
			return nil, fmt.Errorf("should be string or time.Time")
		}

		t := v.Interface().(time.Time)

		if !packing {
			buf = protowire.AppendTag(buf, n, protowire.VarintType)
		}

		if opt.fractionalTime {
			buf = protowire.AppendVarint(buf, encodeTimeFraction(t))
		} else {
			buf = protowire.AppendVarint(buf, encodeTime(t))
		}

	case bigquery.DateTimeFieldType:
		if v.Type() != typeTime {
			return nil, fmt.Errorf("should be string or time.Time")
		}

		t := v.Interface().(time.Time)

		if !packing {
			buf = protowire.AppendTag(buf, n, protowire.VarintType)
		}

		if opt.fractionalTime {
			buf = protowire.AppendVarint(buf, encodeDatetimeFraction(t))
		} else {
			buf = protowire.AppendVarint(buf, encodeDatetime(t))
		}

	case bigquery.GeographyFieldType,
		bigquery.BigNumericFieldType,
		bigquery.IntervalFieldType,
		bigquery.JSONFieldType:
		return nil, fmt.Errorf("not implemented type: %s", fieldType)

	case bigquery.RecordFieldType:
		fallthrough

	default:
		return nil, fmt.Errorf("invalid schema type: %s", fieldType)
	}

	return buf, nil
}

func (opt *option) marshalSliceForFieldSchema(fieldSchema *bigquery.FieldSchema, n protowire.Number, value reflect.Value) ([]byte, error) {
	if !fieldSchema.Repeated {
		return nil, fmt.Errorf("not repeated schema")
	}

	if value.Kind() != reflect.Slice {
		return nil, fmt.Errorf("should be slice: %s", value.Kind())
	}

	if value.Len() == 0 {
		return []byte{}, nil
	}

	var (
		sliceType reflect.Type
		packing   bool
	)

	buf := make([]byte, 0, value.Len()*buildValueSize)

	for i := 0; i < value.Len(); i++ {
		v := value.Index(i)

		for v.Kind() == reflect.Pointer || v.Kind() == reflect.Interface {
			v = v.Elem()
		}

		if sliceType == nil {
			sliceType = v.Type()
			packing = shouldPack(v)
		} else if sliceType != v.Type() {
			return nil, fmt.Errorf("should not have multiple type in slice: a=%s b=%s", sliceType, v.Type())
		}

		var (
			b   []byte
			err error
		)

		switch {
		default:
			fallthrough
		case v.Type() == typeTime, v.Type() == typeRat:
			b, err = opt.marshalValueForFieldSchema(fieldSchema.Type, n, v, packing)
		case v.Kind() == reflect.Struct:
			b, err = opt.marshalStructForSchema(fieldSchema.Schema, n, v)
		}

		if err != nil {
			return nil, err
		}

		buf = append(buf, b...)
	}

	if packing {
		b := make([]byte, 0, 2)
		b = protowire.AppendTag(b, n, protowire.BytesType)
		b = protowire.AppendVarint(b, uint64(len(buf)))
		buf = append(b, buf...)
	}

	return buf, nil
}

func shouldPack(v reflect.Value) bool {
	switch {
	case v.CanInt(), v.CanUint(), v.CanFloat(), v.Type() == typeTime, v.Kind() == reflect.Bool:
		return true

	default:
		return false
	}
}
