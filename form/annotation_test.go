package form

import (
	"testing"
)

type Person struct {
	FirstName  string `form:"firstname,text"`
	LastName   string `form:",text"`
	Email      string `form:"email,email"`
	Age        int    `form:"age,number"`
	Occupation string `form:"occupation,select"`
	Address    *Address
}
type Address struct {
	Street string
	Zip    string
	City   string
	State  string
}

func TestParseAnnotation(t *testing.T) {
	a := parseAnnotation(`form:"a, b, c, d"`)
	if a.Name != "a" {
		t.Errorf("Expected a, got %q", a.Name)
	}
	if a.FieldType != "b" {
		t.Errorf("Expected b, got %q", a.FieldType)
	}
	if a.Generator != "c" {
		t.Errorf("Expected c")
	}
	if a.Validator != "d" {
		t.Errorf("Expected d")
	}
}

func TestGenerate(t *testing.T) {
	if _, err := Generate("hello"); err == nil {
		t.Error("Expected Generate to fail when given a non-struct")
	}

	data := &Person{
		FirstName: "David",
		LastName:  "Hume",
		Email:     "skeptic@example.com",
		Address: &Address{
			Street: "123 Missing Blue",
		},
	}

	_, err := Generate(data)
	if err != nil {
		t.Errorf("Expected form: %s", err)
	}

}
