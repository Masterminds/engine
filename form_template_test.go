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

	out, err := e.Render("form.html.tpl", f)
	if err != nil {
		t.Errorf("Failed render of form.html.tpl: %s", err)
	}
	t.Log(out)
}
