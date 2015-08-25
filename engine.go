package engine

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"
)

// NoTemplateFound indicates that a desired template cannot be located.
var (
	NoTemplateFound = errors.New("no template found")
	NoAssetFound    = errors.New("no asset found")
	IllegalName     = errors.New("name contains illegal patterns")
)

// New creates a new *Template and processes the given directories.
//
// Each path should point to a "theme" directory. That directory is scanned
// for templates, which are compiled immediately and then cached.
//
// Paths are normalized, but the resulting normalized version cannot have
// a relative path. So the path "foo/bar/.." is fine, and will evaluate to
// "foo", but the path "foo/../.." is not okay, as it will evaluate to
// "..", which represents a potential security risk.
//
// Each path is scanned for files that end with the extension '.tpl'.
// Directories are not scanned recursively. Any other files or directories are
// ignored.
func New(paths ...string) (*Engine, error) {

	// First, we do a quick normalization of all paths.
	for i, d := range paths {
		d = filepath.Clean(d)

		// Clean will resolve '..' when possible. If any are left, it means the
		// relative path is above the given directory, and that's a security
		// problem.
		if !legalName(d) {
			return nil, IllegalName
		}

		if !dirExists(d) {
			return nil, fmt.Errorf("could not read directory '%s'", d)
		}
		paths[i] = d
	}

	e := &Engine{
		dirs:  paths,
		cache: make(map[string]map[string]*template.Template, len(paths)),
	}

	return e, e.parse()
}

type Engine struct {
	// Order is important, so we keep dirs to maintain an ordering of themes.
	dirs []string

	// cache is a cache of parsed templates.
	// Cache maps are of the form map[basePath][relPath]template.
	cache map[string]map[string]*template.Template
}

// Render looks for a template with the given name, then executes it with the given data.
//
// The 'name' parameter should be a relative template name (foo.tpl). This
// will look through all of the known templates and execute the first match
// found. Traversal order is the order in which the templates were added.
//
// The 'data' will be passed into the template unaltered.
//
// If the renderer cannot find a template, it returns NoTemplateFound. If
// the template cannot be rendered, it may return a different error.
func (e *Engine) Render(name string, data interface{}) (string, error) {
	n := filepath.Clean(name)
	for _, d := range e.dirs {
		if t, ok := e.cache[d][n]; ok {
			var buf bytes.Buffer
			err := t.ExecuteTemplate(&buf, n, data)
			return buf.String(), err
		}
	}
	return "", NoTemplateFound
}

// Asset returns the first matching asset path.
//
// An asset is a non-template file or directory in a theme directory. This
// function returns the string path of the first path that matches.
//
// An asset path is only returned if the asset exists and can be stat'ed.
func (e *Engine) Asset(name string) (string, error) {
	name = filepath.Clean(name)
	if !legalName(name) {
		return "", IllegalName
	}
	for _, d := range e.dirs {
		p := filepath.Join(d, name)
		if _, err := os.Stat(p); err == nil {
			return p, nil
		}
	}

	return "", NoAssetFound
}

// Dirs returns a list of directories that this Engine knows about.
//
// Directories are presented in their cleaned, but not absolute, form.
func (e *Engine) Dirs() []string {
	return e.dirs
}

// Paths returns all know template paths.
func (e *Engine) Paths() []string {
	res := make([]string, 0, len(e.dirs))
	for base, tt := range e.cache {
		for rel, _ := range tt {
			res = append(res, filepath.Join(base, rel))
		}
	}
	return res
}

func dirExists(d string) bool {
	fi, err := os.Stat(d)
	if err != nil {
		return false
	}
	return fi.IsDir()
}

func legalName(d string) bool {
	// I'm not sure how OS-independent this would turn out to be.
	// An alternative test might be to call Abs and make sure that this
	// path is "present" in Abs. That, however, would render "." illegal.
	return !strings.Contains(d, "..")
}

func clean(d string) string {
	return filepath.Clean(d)
}

func (e *Engine) parse() error {

	// XXX: It is assumed that e.dirs have already been normalized and
	// checked.
	for _, d := range e.dirs {

		ts := filepath.Join(d, "*.tpl")
		files, err := filepath.Glob(ts)
		if err != nil {
			// ErrBadPattern is the only error that
			// will return. files is nil if the pattern didn't turn up
			// anything.
			return err
		}

		// An dir with no templates is totally legit. This directory may
		// just contain other assets. So we add to the map and continue.
		if files == nil {
			e.cache[d] = map[string]*template.Template{}
			continue
		}

		e.cache[d] = make(map[string]*template.Template, len(files))
		for _, f := range files {
			// Second half of cache key.
			r, err := filepath.Rel(d, f)
			if err != nil {
				return err
			}

			t, err := template.New(f).ParseFiles(f)
			if err != nil {
				return err
			}

			e.cache[d][r] = t
		}
	}
	return nil
}
