package form

import (
	"fmt"
	"reflect"
	"strings"

	"golang.org/x/net/html"
)

// attr is a utility for appending attributes.
func attr(a []html.Attribute, n, v string) []html.Attribute {
	return append(a, html.Attribute{Key: n, Val: v})
}

// toAttr tries to turn whatever is passed in into an HTML attribute.
func toAttr(o interface{}) (html.Attribute, bool) {
	v := reflect.Indirect(reflect.ValueOf(o))
	if v.IsNil() {
		return html.Attribute{}, false
	}
	if val, ok := aval(o); ok {
		t := v.Type()
		return html.Attribute{Key: strings.ToLower(t.Name()), Val: val}, true
	}
	return html.Attribute{}, false
}

// aval converts a value to an attribute string value
//
// The key is the lowercased name of the thing.
//
// The values are calculated as follows:
//
// 	- Bools become "true" and "false" strings
// 	- Numeric types are converted to strings
// 	- Strings are omitted if the len() == 0, or encoded as-is
//
// Anything else returns an empty string, and a false .
func aval(o interface{}) (string, bool) {
	v := reflect.Indirect(reflect.ValueOf(o))

	switch v.Kind() {
	case reflect.Bool:
		if v.Bool() == true {
			return "true", true
		}
		return "false", true
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return fmt.Sprintf("%d", o), true
	case reflect.Float32, reflect.Float64:
		return fmt.Sprintf("%f", o), true
	case reflect.String:
		o := o.(string)
		if len(o) == 0 {
			return o, false
		}
		return o, true
	}
	println("Did not match type.")

	return "", false
}

// structToAttrs converts fields on a struct into attributes.
// Given a struct and a list of field names, convert to HTML attributes.
func structToAttrs(s interface{}, names ...string) []html.Attribute {

	c := make(map[string]struct{}, len(names))
	for _, n := range names {
		c[n] = struct{}{}
	}

	v := reflect.Indirect(reflect.ValueOf(s))
	if v.Kind() != reflect.Struct {
		return []html.Attribute{}
	}
	nf := v.NumField()

	a := []html.Attribute{}
	for i := 0; i < nf; i++ {
		fv := v.Field(i)
		fn := v.Type().Field(i).Name
		if _, ok := c[fn]; ok {
			if v, ok := aval(fv.Interface()); ok {
				a = attr(a, strings.ToLower(fn), v)
			}
		}
	}
	return a
}
