package engine

import (
	"testing"

	"github.com/Masterminds/engine/form"
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
	}

	out, err := e.Render("form.html.tpl", f)
	if err != nil {
		t.Errorf("Failed render of form.html.tpl: %s", err)
	}
	t.Log(out)
}
