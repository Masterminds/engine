package engine

import (
	"strings"
	"testing"
)

func TestNoDirs(t *testing.T) {

	if _, err := New("/no/such/path"); err == nil {
		t.Error("Expected failure when given bogus path")
	}

	// Canary
	if _, err := New("testdata"); err != nil {
		t.Errorf("Expected success, got '%s'", err)
	}

	if _, err := New("testdata", "/no/such/path"); err == nil {
		t.Error("Expected failure when given bogus path")
	}

}

func TestLegalName(t *testing.T) {
	data := map[string]bool{
		"..":           false,
		".":            true,
		"foo/..":       false,
		"foo/.":        true,
		"foo":          true,
		"/foo/":        true,
		"/foo":         true,
		"/../":         false,
		"../../../":    false,
		"./.././../":   false,
		"foo/bar baz/": true,
	}

	for d, expect := range data {
		if legalName(d) != expect {
			t.Errorf("Expected %s to be %v", d, expect)
		}
	}
}
func TestClean(t *testing.T) {
	data := map[string]string{
		"":             ".",
		"..":           "..",
		".":            ".",
		"foo/..":       ".",
		"foo/.":        "foo",
		"foo":          "foo",
		"/foo/":        "/foo",
		"/foo":         "/foo",
		"/../":         "/",
		"../../../":    "../../..",
		"./.././../":   "../..",
		"foo/bar baz/": "foo/bar baz",
		"foo/bar/..":   "foo",
		"./testdata":   "testdata",
	}

	for d, expect := range data {
		if o := clean(d); o != expect {
			t.Errorf("Expected %s to be %s, got %s", d, expect, o)
		}
	}
}

func TestNew(t *testing.T) {
	e, err := New("./testdata")
	if err != nil {
		t.Errorf("Failed parse of testdata: %s", err)
	}
	dirs := e.Dirs()
	if len(dirs) != 1 {
		t.Errorf("Expected only one template directory.")
	}

	if dirs[0] != "testdata" {
		t.Errorf("Expected 'testdata', got '%s'", dirs[0])
	}
}

func TestRender(t *testing.T) {
	e, err := New("testdata")
	if err != nil {
		t.Errorf("Failed parse of testdata: %s", err)
	}

	paths := e.Paths()
	if len(paths) != 1 {
		t.Errorf("Expected 1 template. Found %d", len(paths))
		for p := range paths {
			println(p)
		}
	}

	out, err := e.Render("simple.tpl", "test")
	if err != nil {
		t.Errorf("Failed render: %s", err)
	}

	out = strings.TrimSpace(out)
	if out != "test" {
		t.Errorf("Expected 'test', got '%s'", out)
	}
}

func TestCascadingRender(t *testing.T) {
	e, err := New("testdata/override", "testdata/base", "testdata")
	if err != nil {
		t.Errorf("Failed parse of testdata: %s", err)
	}

	out, err := e.Render("simple.tpl", "test")
	if err != nil {
		t.Errorf("Failed render: %s", err)
	}

	out = strings.TrimSpace(out)
	if out != "OVERRIDE:test" {
		t.Errorf("Expected 'OVERRIDE:test', got '%s'", out)
	}

	out, err = e.Render("onlybase.tpl", "test")
	if err != nil {
		t.Errorf("Could not render onlybase: %s", err)
	}

	out = strings.TrimSpace(out)
	if out != "onlybase:test" {
		t.Errorf("Expected 'onlybase:test', got '%s'", out)
	}
}

func TestAsset(t *testing.T) {
	e, err := New("testdata/override", "testdata/base")
	if err != nil {
		t.Errorf("Failed parse of testdata: %s", err)
	}

	a, err := e.Asset("asset.dat")
	if err != nil {
		t.Errorf("Could not find asset.dat: %s", err)
	}

	expect := "testdata/base/asset.dat"
	if a != expect {
		t.Errorf("Expected %s, got %s", expect, a)
	}

	// Make sure we can't render a non-template.
	if _, err := e.Render("asset.dat", 42); err == nil {
		t.Errorf("Should not be able to render asset.dat")
	}

}
