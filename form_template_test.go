package engine

import (
	"strings"
	"testing"

	"github.com/Masterminds/engine/form"
	"golang.org/x/net/html"
)

func TestFormTemplate(t *testing.T) {
	e, err := New("_template")
	if err != nil {
		t.Errorf("Failed to load templates: %s", err)
	}

	f := form.New("name", "form-action")
	f.Id = "1234"
	f.Lang = "en-us"
	f.Title = "Login"
	f.Hidden = form.OFalse
	f.Class = []string{"foo", "bar", "baz"}
	f.Method = "POST"
	f.Autocomplete = true
	f.Fields = []form.Field{
		form.Button{
			GlobalAttributes: form.GlobalAttributes{Id: "button-1"},
			Name:             "button-1",
			Value:            "Push Me!",
		},
		form.FieldSet{
			Name:   "fset1",
			Legend: "Look, Ma! A Fieldset!",
			Fields: []form.Field{
				&form.Button{
					GlobalAttributes: form.GlobalAttributes{Id: "button-2"},
					Name:             "button-2",
					Value:            "Push Me Too!",
				},
			},
		},
		form.Keygen{
			Name:      "keygen",
			KeyType:   "rsa",
			Challenge: "How much would could a wood chuck chuck?",
		},
		form.Label{For: "ever", Text: "Don't Label Me"},
		form.Output{For: "your eyes only", Name: "Bond, James Bond"},
		form.Progress{Value: 0.5, Max: 1.0},
		form.Meter{Value: 0.5, Max: 1.0, Min: 0.2, Optimum: 0.7, Low: 0.1, High: 0.5},
		form.DataList{
			GlobalAttributes: form.GlobalAttributes{Id: "dl-1"},
			Options: []form.Option{
				{Value: "one", Label: "One"},
				{Value: "two", Label: "Two"},
			},
		},
		form.Select{
			Name:  "cookies",
			Label: "How many cookies?",
			Options: []form.OptionItem{
				form.Option{Value: "one", Label: "One"},
				form.Option{Value: "two", Label: "Two"},
				form.OptGroup{
					Label: "optgroup",
					Options: []form.Option{
						form.Option{Value: "three", Label: "Three"},
						form.Option{Value: "four", Label: "Four"},
					},
				},
			},
		},
		form.TextArea{
			Name:  "textarea",
			Cols:  80,
			Rows:  5,
			Value: "Default text",
		},
		form.Password{Name: "password", Label: "Enter Password"},
		form.Text{Name: "text"},
		form.Submit{Name: "submit"},
		form.Tel{Name: "tel"},
		form.URL{Name: "url"},
		form.Email{Name: "email"},
		form.Date{Name: "date"},
		form.Time{Name: "time"},
		form.Number{Name: "number"},
		form.Range{Name: "range"},
		form.Color{Name: "color"},
		form.Checkbox{Name: "checkbox"},
		form.Radio{Name: "radio"},
		form.File{Name: "file"},
		form.Image{Name: "image"},
		form.Reset{Name: "reset"},
		form.Hidden{Name: "hidden"},
	}

	out, err := e.Render("#form", f)
	if err != nil {
		t.Errorf("Failed render of form.html.tpl: %s", err)
	}
	t.Log(out)

	// Now we load the result with an HTML parser to ensure that it's
	// well-formed.
	read := strings.NewReader(out)
	if _, err := html.Parse(read); err != nil {
		t.Errorf("Failed to parse generated markup: %s", err)
	}

}
