package form

import (
	"fmt"
	"testing"

	"golang.org/x/net/html"
)

func TestForm(t *testing.T) {
	f := New("name", "/foo")
	f.Target = "target"
	f.Class = []string{"foo", "bar"}
	f.Title = "title"

	node := f.Element()

	expectAttrs(t, node, map[string]string{
		"name":   "name",
		"id":     "name",
		"class":  "foo bar",
		"target": "target",
		"title":  "title",
	})

}

func TestAsValues(t *testing.T) {
	f := Form{
		Name: "test",
		Fields: []Field{
			&Password{Name: "pass", Value: "foo"},
			&Checkbox{Name: "check", Value: "a", Checked: true},
			&Checkbox{Name: "check", Value: "b", Checked: true},
			&Select{
				Name: "choose", Options: []OptionItem{
					&OptGroup{
						Options: []*Option{
							&Option{Value: "first", Selected: true},
						},
					},
					&Option{Value: "second", Selected: false},
				},
			},
			&Button{Name: "button", Value: "submit"},
		},
	}

	m := f.AsValues()
	if len(*m) != 4 {
		t.Errorf("Expected 4 items, got %d", len(*m))
	}
	if m.Get("pass") != "foo" {
		t.Errorf("Wrong pass value %s", m.Get("pass"))
	}

	if m.Get("choose") != "first" {
		t.Errorf("expected 'first', got %q", m.Get("choose"))
	}

	vals := *m
	check := vals["check"]
	if len(check) != 2 {
		t.Errorf("expected two checked values.")
	}
}
func ExampleAsValues() {

	// A form with a group of checkboxes and a select list.
	f := Form{
		Name: "test",
		Fields: []Field{
			&Checkbox{Name: "check", Value: "a", Checked: true},
			&Checkbox{Name: "check", Value: "b", Checked: true},
			&Select{
				Name: "choose", Options: []OptionItem{
					&Option{Value: "first", Selected: true},
					&Option{Value: "second", Selected: false},
				},
			},
		},
	}

	values := f.AsValues()

	// Get the value of the select list.
	fmt.Println(values.Get("choose"))

	// Get the values of the check boxes.
	valmap := *values
	check := valmap["check"]
	fmt.Printf("%d checked", len(check))
	// Output:
	// first
	// 2 checked
}

func TestFormFields(t *testing.T) {
	f := Form{
		Name:   "test",
		Action: "/action",
		Fields: []Field{
			&Div{HTML: HTML{Class: []string{"test"}}},
			String("test"),
			&Button{Name: "Button"},
			&ButtonInput{Name: "buttonInput"},
			&Password{Name: "password", Size: 10},
			&Text{Name: "text", Size: 5},
			&Submit{Name: "submit"},
			&Tel{Name: "telephone"},
			&URL{Name: "url"},
			&Email{Name: "Email"},
			&Date{Name: "date"},
			&Time{Name: "time"},
			&Number{Name: "number"},
			&Range{Name: "range"},
			&Color{Name: "color"},
			&Checkbox{Name: "checkbox", Value: "one"},
			&Checkbox{Name: "checkbox", Value: "two"},
			&Radio{Name: "radio", Value: "on"},
			&Radio{Name: "radio", Value: "off"},
			&File{Name: "file"},
			&Image{Name: "image"},
			&Reset{Name: "reset"},
			&Hidden{Name: "hidden", Value: "Shhh"},
			&FieldSet{
				Name:   "fields",
				Legend: "Use me!",
				Fields: []Field{
					&Radio{Name: "radio2", Value: "on"},
					&Radio{Name: "radio2", Value: "off"},
				},
			},
			&Keygen{Name: "keygen", Disabled: true},
			&Output{Name: "output"},
			&Progress{Value: 0.0, Max: 1.0},
			&Meter{Min: 0.0, Max: 1.0},
			&Select{
				Name: "select",
				Options: []OptionItem{
					&Option{Value: "Option 1"},
					&OptGroup{
						Label: "Opt Group",
						Options: []*Option{
							&Option{Value: "Option 2"},
							&Option{Value: "Option 3"},
						},
					},
				},
			},
			TextArea{Name: "textarea", Rows: 4, Cols: 40},
		},
	}

	// TODO: Replace with a better test.
	for _, i := range f.Fields {
		if i == nil {
			t.Error("Surprised to find a nil field.")
		}
	}
}

func expectAttrs(t *testing.T, n *html.Node, e map[string]string) {
	attrMap := make(map[string]string)
	for _, a := range n.Attr {
		t.Logf("Adding attr %s: %s", a.Key, a.Val)
		attrMap[a.Key] = a.Val
	}
	for k, v := range e {
		if found, ok := attrMap[k]; !ok {
			t.Errorf("Expected to find attribute '%s'.", k)
			t.Logf("Node: %v", n)
		} else if found != v {
			t.Errorf("Expected %s to be '%s', got '%s'", k, v, found)
		}
	}
}
