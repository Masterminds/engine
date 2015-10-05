package form

import (
	"fmt"
	"net/url"
	"testing"
	"time"
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

	Reconcile(f, &v)

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

func TestFormHandler(t *testing.T) {
	fh := NewFormHandler(NewCache(), time.Minute)

	f := New("test", "test")
	f.Fields = []Field{
		&Password{Name: "p", Value: "secret"},
	}

	id, err := fh.Prepare(f)
	if err != nil {
		t.Errorf("Error preparing form: %s", err)
	}

	if len(id) < SecurityTokenLength {
		t.Errorf("Expected token length %d, got %d", SecurityTokenLength, len(id))
	}

	if ff, err := fh.Get(id); err != nil {
		t.Errorf("Could not get form %s: %s", id, err)
	} else if ff.Name != "test" {
		t.Errorf("Expected form named 'test', got %q", ff.Name)
	}

	vals := &url.Values{
		"p":             []string{"cheese"},
		SecureTokenName: []string{id},
	}

	ff, err := fh.Retrieve(vals)
	if err != nil {
		t.Errorf("Failed to retrieve form: %s", err)
	}

	cheese := ff.Fields[0].(*Password).Value
	if cheese != "cheese" {
		t.Errorf("Expected cheese, got %q", cheese)
	}
}

func Example() {

	// Create a new form:
	f := New("My Form", "/submit")

	// Add some fields
	f.Fields = []Field{
		&Text{Name: "username"},
		&Password{Name: "password"},
	}

	// Prepare the form using the default form handler.
	id, err := DefaultFormHandler.Prepare(f)
	if err != nil {
		fmt.Printf("Error preparing form: %s", err)
		return
	}

	// Render the form to the user agent. Typlically you do this
	// with a template.

	// When the form is submitted, it will return something like this:
	vals := &url.Values{
		"username":      []string{"matt"},
		"password":      []string{"secret"},
		SecureTokenName: []string{id},
	}

	// We can pass in the values and retrieve the form, with the values
	// all entered on the appropriate element.
	ff, err := DefaultFormHandler.Retrieve(vals)
	if err != nil {
		fmt.Printf("Failed to retrieve form: %s", err)
		return
	}

	// Now we can access the fields directly
	user := ff.Fields[0].(*Text).Value
	pass := ff.Fields[1].(*Password).Value

	fmt.Println(user)
	fmt.Println(pass)
	// Output:
	// matt
	// secret
}
