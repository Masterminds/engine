package form

import (
	"errors"
	"fmt"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

// Validator indicates whether a form field is valid.
type Validator func(field string, value []string) bool

// FieldUnmarshaler unmarshals a specific field.
type FieldUnmarshaler func(field string, value []string) (interface{}, error)

func Unmarshal(v url.Values, o interface{}) error {
	val := reflect.ValueOf(o)
	if val.Kind() != reflect.Ptr || val.IsNil() {
		return errors.New("unmarshal requires a pointer to a receiver")
	}

	return walk(val, v)
}

func walk(val reflect.Value, v url.Values) error {
	// Loop through values, top-down specificity
	for key, vals := range v {
		findIn(val, key, vals)
	}
	return nil
}

func findIn(rv reflect.Value, key string, values []string) error {
	switch rv.Kind() {
	case reflect.Map:
		// The map must take string keys.
		if _, ok := rv.Interface().(map[string]interface{}); ok {
			return assignToMap(rv, key, values)
		}
	case reflect.Struct:
		// Look for struct field named 'key'.
		return assignToStruct(rv, key, values)
	}
	return fmt.Errorf("object %s cannot be used to store values", rv.Type().Name())
}

func assignToMap(rv reflect.Value, key string, values []string) error {
	var err error
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("Failed map assignment: %v\n", e)
			// FIXME: can't modify err in recover.
			err = fmt.Errorf("failed map assignment: %s", e)
		}
	}()
	// FIXME: There must be a way to find the destination type of a map and
	// appropriately convert to it.
	switch l := len(values); {
	case l == 1:
		rv.SetMapIndex(reflect.ValueOf(key), reflect.ValueOf(values[0]))
	case l > 1:
		rv.SetMapIndex(reflect.ValueOf(key), reflect.ValueOf(values))
	}
	return err
}

func assignToStruct(rv reflect.Value, key string, values []string) error {
	rt := rv.Type()
	// Look for a Field on struct that matches the key name.
	for i := 0; i < rt.NumField(); i++ {
		f := rt.Field(i)
		tag := parseTag(f.Tag.Get("form"))
		if !tag.ignore && tag.name == "" {
			tag.name = f.Name
		}
		if tag.name == key {
			fmt.Printf("Assigning %s = %q\n", key, values[0])
			assignToStructField(rv.FieldByName(f.Name), values)
			return nil
		}
	}
	fmt.Printf("Skipped key %q", key)
	return nil
}

func assignToStructField(rv reflect.Value, val []string) error {
	// Basically, we need to convert from a string to the appropriate underlying
	// kind, then assign.
	switch rv.Kind() {
	case reflect.String:
		rv.Set(reflect.ValueOf(val[0]))
		return nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		vv := "0"
		if len(val) > 0 {
			vv = val[0]
		}
		return assignToInt(rv, vv)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		vv := "0"
		if len(val) > 0 {
			vv = val[0]
		}
		return assignToUint(rv, vv)
	case reflect.Float32, reflect.Float64:
		vv := "0"
		if len(val) > 0 {
			vv = val[0]
		}
		return assignToFloat(rv, vv)
	case reflect.Bool:
		b, err := strconv.ParseBool(val[0])
		reflect.Indirect(rv).Set(reflect.ValueOf(b))
		return err
	default:
		return fmt.Errorf("Unsupported kind")
	}
}

func assignToInt(rv reflect.Value, val string) error {
	rvv := reflect.Indirect(rv)
	if !rvv.CanSet() {
		return fmt.Errorf("cannot set %q (%s)", rv.Type().Name(), rv.Kind().String())
	}
	ival, err := strconv.ParseInt(val, 0, 0)
	if err != nil {
		return err
	}
	rvv.SetInt(ival)
	return nil
}
func assignToUint(rv reflect.Value, val string) error {
	rvv := reflect.Indirect(rv)
	if !rvv.CanSet() {
		return fmt.Errorf("cannot set %q (%s)", rv.Type().Name(), rv.Kind().String())
	}
	ival, err := strconv.ParseUint(val, 0, 0)
	if err != nil {
		return err
	}
	rvv.SetUint(ival)
	return nil
}
func assignToFloat(rv reflect.Value, val string) error {
	rvv := reflect.Indirect(rv)
	if !rvv.CanSet() {
		return fmt.Errorf("cannot set %q (%s)", rv.Type().Name(), rv.Kind().String())
	}
	ival, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return err
	}
	rvv.SetFloat(ival)
	return nil
}

func parseTag(str string) *tag {
	parts := strings.Split(str, ",")
	if len(parts) == 1 && parts[0] == "" {
		return &tag{}
	}
	t := &tag{}
	switch n := parts[0]; n {
	case "+":
		t.group = true
	case "-":
		t.ignore = true
	default:
		t.name = n
	}

	for _, p := range parts[1:] {
		switch {
		case p == "omitempty":
			t.omit = true
		case strings.HasPrefix(p, "prefix="):
			t.prefix = strings.TrimPrefix(p, "prefix=")
		case strings.HasPrefix(p, "suffix="):
			t.suffix = strings.TrimPrefix(p, "suffix=")
		}
	}
	return t
}

// tag represents a 'form' tag.
//
//	Name string `form:name`
//	Date time.Time `form:date,omitempty`
//	Address *Address `form:+,omitempty,prefix=addr_
type tag struct {
	name           string
	prefix, suffix string //prefix=, suffix=
	omit           bool   // omitempty
	ignore         bool   // -
	group          bool   // +
	validator      Validator
	unmarshaler    FieldUnmarshaler
}
