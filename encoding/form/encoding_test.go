package form

import (
	"fmt"
	"net/url"
	"reflect"
	"testing"
)

type ExampleAddress struct {
	Street string `form:"street"`
	City   string `form:"city"`
	State  string `form:"state"`
	Zip    string `form:"zip"`
}

type ExampleForm struct {
	// Retrieve these files from the form.
	FirstName string `form:"first_name"`
	LastName  string `form:"last_name"`
	// Retrieve these fields, but omit empty address
	HomeAddress *ExampleAddress `form:"+,omitempty"`
	// This uses prefix to modify the field name on the address. So 'street'
	// becomes 'mail_street'.
	MailingAddress *ExampleAddress `form:"+,omitempty,prefix=mail_"`
	// Ignore this field
	Processed bool `form:"-"`
}

func ExampleUnmarshal() {
	v := url.Values{}
	v.Set("first_name", "Matt")
	v.Set("street", "1234 Tape Dr")
	v.Set("mail_street", "4321 Disk Dr")

	ef := &ExampleForm{}
	if err := Unmarshal(v, ef); err != nil {
		panic(err)
	}

	fmt.Printf("Form: %v", v)
}

/*
func TestUnmarshal(t *testing.T) {
	v := url.Values{}
	v.Set("first_name", "Matt")
	v.Set("street", "1234 Tape Dr")
	v.Set("mail_street", "4321 Disk Dr")

	ef := &ExampleForm{}
	if err := Unmarshal(v, ef); err != nil {
		t.Fatal(err)
	}

	if ef.FirstName != "Matt" {
		t.Errorf("Expected Matt, got %q", ef.FirstName)
	}

	if ef.LastName != "" {
		t.Errorf("Expected empty string, got %q", ef.LastName)
	}

	if ef.StreetAddress.Street != "1234 Tape Dr" {
		t.Errorf("Unexpected mailing address: %q", ef.MailingAddress.Street)
	}

	if ef.MailingAddress.Street != "4321 Disk Dr" {
		t.Errorf("Unexpected mailing address: %q", ef.MailingAddress.Street)
	}
}
*/

func TestParseTag(t *testing.T) {
	tests := []struct {
		name   string
		tag    string
		expect tag
	}{
		{
			name:   "name only",
			tag:    "first_name",
			expect: tag{name: "first_name"},
		},
		{
			name:   "name, omitempty",
			tag:    "first_name,omitempty",
			expect: tag{name: "first_name", omit: true},
		},
		{
			name:   "ignore",
			tag:    "-",
			expect: tag{ignore: true},
		},
		{
			name:   "christmas tree",
			tag:    "name,prefix=pre_,suffix=suf_,omitempty",
			expect: tag{name: "name", prefix: "pre_", suffix: "suf_", omit: true},
		},
		{
			name:   "group",
			tag:    "+,prefix=pre_,suffix=suf_,omitempty",
			expect: tag{group: true, prefix: "pre_", suffix: "suf_", omit: true},
		},
	}

	for _, tt := range tests {
		got := parseTag(tt.tag)
		expect := tt.expect
		if got.name != expect.name {
			t.Errorf("%s expected %q, got %q", tt.name, expect.name, got.name)
		}
		if got.prefix != expect.prefix {
			t.Errorf("%s expected %q, got %q", tt.name, expect.prefix, got.prefix)
		}
		if got.suffix != expect.suffix {
			t.Errorf("%s expected %q, got %q", tt.name, expect.suffix, got.suffix)
		}
		if got.group != expect.group {
			t.Errorf("%s expected %t got %t", tt.name, expect.group, got.group)
		}
		if got.ignore != expect.ignore {
			t.Errorf("%s expected %t got %t", tt.name, expect.ignore, got.ignore)
		}
		if got.omit != expect.omit {
			t.Errorf("%s expected %t got %t", tt.name, expect.omit, got.omit)
		}
	}
}

func TestAssignToMap(t *testing.T) {
	m := map[string]string{}
	mv := reflect.ValueOf(m)

	if err := assignToMap(mv, "test", []string{"first"}); err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	if m["test"] != "first" {
		t.Errorf("Expected test key to have 'first', got '%v'", m["test"])
	}

	if err := assignToMap(mv, "test2", []string{"first", "second"}); err == nil {
		t.Errorf("Expeced an error assigning multiple values to single value. (%v)", m["test2"])
	} else if err.Error() != "foo" {
		t.Errorf("Unexpected error: %s (multi-val)", err)
	}
}

type intStruct struct {
	I   int
	I8  int8
	I16 int16
	I32 int32
	I64 int64
}

func TestAssignToInt(t *testing.T) {

	is := &intStruct{8, 8, 8, 8, 8}

	tests := []struct {
		name string
		rv   reflect.Value
		src  string
	}{
		{"int", reflect.ValueOf(&is.I), "64"},
		{"int8", reflect.ValueOf(&is.I8), "8"},
		{"int16", reflect.ValueOf(&is.I16), "16"},
		{"int32", reflect.ValueOf(&is.I32), "32"},
		{"int64", reflect.ValueOf(&is.I64), "64"},
	}

	for _, tt := range tests {
		if err := assignToInt(tt.rv, tt.src); err != nil {
			t.Fatal(err)
		}
		if got := fmt.Sprintf("%v", reflect.Indirect(tt.rv).Interface()); got != tt.src {
			t.Errorf("Expected %q, got %q", tt.src, got)
		}
	}
}

type AssignmentTestStruct struct {
	FirstName string `form:"first_name"`
	LastName  string
	Year      uint32
	Speed     float64
	IsUseless bool
}

func TestAssignToStruct(t *testing.T) {
	ats := &AssignmentTestStruct{}
	rats := reflect.Indirect(reflect.ValueOf(ats))

	if err := assignToStruct(rats, "first_name", []string{"Matt"}); err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	if ats.FirstName != "Matt" {
		t.Errorf("Expected ats.FirstName to be Matt, got %q", ats.FirstName)
	}

	if err := assignToStruct(rats, "LastName", []string{"Butcher"}); err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	if ats.LastName != "Butcher" {
		t.Errorf("Expected ats.LastName to be Butcher, got %q", ats.LastName)
	}

	if err := assignToStruct(rats, "Year", []string{"1999"}); err != nil {
		t.Errorf("Unexpected error assigning Year: %s", err)
	}
	if ats.Year != 1999 {
		t.Errorf("Expected Year=1999, got %d", ats.Year)
	}
	if err := assignToStruct(rats, "Speed", []string{"1.23"}); err != nil {
		t.Errorf("Unexpected error assigning Year: %s", err)
	}
	if ats.Speed != 1.23 {
		t.Errorf("Expected speed=1.23, got %d", ats.Speed)
	}
	if err := assignToStruct(rats, "IsUseless", []string{"true"}); err != nil {
		t.Errorf("Unexpected error assigning Year: %s", err)
	}
	if !ats.IsUseless {
		t.Error("Expected IsUseless to be true")
	}
}
