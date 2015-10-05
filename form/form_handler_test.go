package form

import (
	"net/url"
	"testing"
)

func TestReconcile(t *testing.T) {
	f := New("test", "test")
	f.Fields = []Field{
		&Text{Name: "hello", Value: "nobody"},
		&Hidden{Name: "goodbye", Value: "girl"},
		&Div{
			Fields: []Field{
				&TextArea{Name: "goodnight", Value: "moon"},
			},
		},
		&Radio{Name: "box", Value: "fun"},
		&Radio{Name: "box", Value: "boring"},
		&Checkbox{Name: "stooge", Value: "Larry"},
		&Checkbox{Name: "stooge", Value: "Moe"},
		&Checkbox{Name: "stooge", Value: "Samantha"},
	}

	v := url.Values{
		"hello":     []string{"world"},
		"goodnight": []string{"cow jumping over the moon"},
		"box":       []string{"boring"},
		"stooge":    []string{"Larry", "Moe"},
	}

	reconcile(f, &v)

	rval := f.Fields[0].(*Text).Value
	if rval != "world" {
		t.Errorf("Expected 'world', got %q", rval)
	}
	cow := f.Fields[2].(*Div).Fields[0].(*TextArea).Value
	if cow != "cow jumping over the moon" {
		t.Errorf("Expected a cow jumping over the moon, got %q", cow)
	}
	if f.Fields[3].(*Radio).Checked {
		t.Errorf("Expected %s to be unchecked.", f.Fields[3].(*Radio).Value)
	}
	if !f.Fields[4].(*Radio).Checked {
		t.Errorf("Expected %s to be checked.", f.Fields[4].(*Radio).Value)
	}
	if !f.Fields[5].(*Checkbox).Checked {
		t.Errorf("Expected %s to be checked.", f.Fields[5].(*Checkbox).Value)
	}
	if f.Fields[7].(*Checkbox).Checked {
		t.Errorf("Expected %s to be unchecked.", f.Fields[7].(*Checkbox).Value)
	}
}
