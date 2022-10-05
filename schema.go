package bqproto

import (
	"fmt"
	"reflect"
	"strings"
	"sync"

	"cloud.google.com/go/bigquery"
)

var (
	ErrMissingRequired = fmt.Errorf("missing required field")
	ErrExistsHindrance = fmt.Errorf("exists hindrance field")
)

func ValidateSchemaCompatibility(tableSchema bigquery.Schema, v any, opts ...Option) error {
	opt := newOption(opts...)

	_, ok := v.(map[string]any)
	if ok {
		return nil
	}

	value := reflect.ValueOf(v)

	for value.Kind() == reflect.Pointer || value.Kind() == reflect.Interface {
		value = value.Elem()
	}

	if value.Kind() != reflect.Struct {
		return fmt.Errorf("should be struct type: %q", value.Type())
	}

	return opt.validateSchemaCompatibility(tableSchema, value.Type())
}

type bqTag struct {
	kind     reflect.Kind
	index    int
	nullable bool
}

type bqTags map[string]bqTag

var (
	cachedb = map[string]bqTags{}
	mux     sync.RWMutex
)

func (opt *option) validateSchemaCompatibility(tableSchema bigquery.Schema, typ reflect.Type) error {
	tags := detectTags(typ, opt.tagKey)

	for _, s := range tableSchema {
		tag, ok := tags[s.Name]

		if s.Required && (!ok || tag.nullable) {
			return fmt.Errorf("%w: %q", ErrMissingRequired, s.Name)
		}
	}

	for k := range tags {
		ok := false

		for _, s := range tableSchema {
			if k == s.Name {
				ok = true
				break
			}
		}

		if !ok {
			return fmt.Errorf("%w: %q", ErrExistsHindrance, k)
		}
	}

	return nil
}

func detectTags(typ reflect.Type, tagName string) bqTags {
	if typ.Kind() != reflect.Struct {
		return nil
	}

	key := fmt.Sprintf("%s.%s", typ.PkgPath(), typ.Name())

	mux.RLock()
	cache, ok := cachedb[key]
	mux.RUnlock()

	if ok {
		return cache
	}

	mux.Lock()
	defer mux.Unlock()

	cache = bqTags{}
	numLen := typ.NumField()

	for i := 0; i < numLen; i++ {
		f := typ.Field(i)

		tag, ok := f.Tag.Lookup(tagName)
		if !ok || tag == "" || tag == "-" {
			continue
		}

		props := strings.Split(tag, ",")
		name := props[0]
		nullable := false

		for _, v := range props[1:] {
			if v == "nullable" {
				nullable = true
			}
		}

		cache[name] = bqTag{
			index:    i,
			kind:     f.Type.Kind(),
			nullable: nullable,
		}
	}

	if key != "." {
		// don't cache anonymouse struct
		cachedb[key] = cache
	}

	return cache
}
