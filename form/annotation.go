package form

import (
	"errors"
	"reflect"
	"regexp"
	"strings"
)

func Generate(v interface{}) (*Form, error) {
	val := reflect.Indirect(reflect.ValueOf(v))
	if val.Kind() != reflect.Struct {
		return nil, errors.New("v must be a struct")
	}
	f := &Form{}

	t := val.Type()
	for i := 0; i < val.NumField(); i++ {
		fv := val.Field(i)
		sf := t.Field(i)
		ele, err := generateField(fv, &sf)
		if err != nil {
			return f, err
		}
		f.Fields = append(f.Fields, ele)
	}

	return f, nil
}

func generateField(v reflect.Value, sf *reflect.StructField) (Field, error) {
	return nil, nil
}

// atn represents an annotation.
//
// form="name,type,generator,validator"
type atn struct {
	Name      string
	FieldType string
	Generator string
	Validator string
}

var are = regexp.MustCompile(`\bform:"([\s\w,]*)"`)

func parseAnnotation(tag string) *atn {
	a := &atn{FieldType: "text"}

	matches := are.FindStringSubmatch(tag)
	if len(matches) != 2 || matches[1] == "" {
		return a
	}

	splode := strings.Split(matches[1], ",")
	sl := len(splode)
	if sl == 0 {
		return a
	}
	a.Name = strings.TrimSpace(splode[0])
	if sl == 1 {
		return a
	}
	ft := strings.TrimSpace(splode[1])
	if ft != "" {
		a.FieldType = ft
	}

	if sl == 2 {
		return a
	}

	gen := strings.TrimSpace(splode[2])
	if gen != "" {
		a.Generator = gen
	}

	if sl == 3 {
		return a
	}

	vd := strings.TrimSpace(splode[3])
	if vd != "" {
		a.Validator = vd
	}
	return a
}
