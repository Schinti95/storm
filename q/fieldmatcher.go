package q

import (
	"errors"
	"reflect"
	"strings"
)

// ErrUnknownField is returned when an unknown field is passed.
var ErrUnknownField = errors.New("unknown field")

type fieldMatcherDelegate struct {
	FieldMatcher
	Field string
}

// NewFieldMatcher creates a Matcher for a given field.
func NewFieldMatcher(field string, fm FieldMatcher) Matcher {
	return fieldMatcherDelegate{Field: field, FieldMatcher: fm}
}

// FieldMatcher can be used in NewFieldMatcher as a simple way to create the
// most common Matcher: A Matcher that evaluates one field's value.
// For more complex scenarios, implement the Matcher interface directly.
type FieldMatcher interface {
	MatchField(v interface{}) (bool, error)
}

func (r fieldMatcherDelegate) Match(i interface{}) (bool, error) {
	v := reflect.Indirect(reflect.ValueOf(i))
	return r.MatchValue(&v)
}

func (r fieldMatcherDelegate) MatchValue(v *reflect.Value) (bool, error) {
	field_name := r.Field
	field_name_arr := strings.Split(field_name, ".")
	var field reflect.Value
	if len(field_name_arr) > 0 {
		field = *v
		for i := range field_name_arr {
			field = field.FieldByName(field_name_arr[i])
			if !field.IsValid() {
				return false, ErrUnknownField
			}
		}
	}
	return r.MatchField(field.Interface())
}
