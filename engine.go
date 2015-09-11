package engine

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/Masterminds/sprig"
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
//
// For convenience, the engine supports an additional set of template
// functions as defined in Sprig:
// 	https://github.com/Masterminds/sprig
// These can be disabled by not passing the Sprig functions into NewEngine.
func New(paths ...string) (*Engine, error) {
	return NewEngine(paths, sprig.FuncMap(), []string{})
}

// NewEngine constructs a new *Engine.
// NewEngine provides more control over the template engine than New.
//
// - funcMap is passed to the template.
// - options are passed to the template.
func NewEngine(paths []string, funcs template.FuncMap, options []string) (*Engine, error) {
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
		dirs:   paths,
		cache:  make(map[string]map[string]bool, len(paths)),
		master: template.New("master"),
	}

	if len(funcs) > 0 {
		e.master.Funcs(funcs)
	}
	if len(options) > 0 {
		e.master.Option(options...)
	}

	return e, e.parse()
}

type Engine struct {
	// Order is important, so we keep dirs to maintain an ordering of themes.
	dirs []string

	cache  map[string]map[string]bool
	master *template.Template
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
		if t, ok := e.cache[d][n]; ok && t {
			var buf bytes.Buffer
			key := filepath.Join(d, n)
			err := e.master.ExecuteTemplate(&buf, key, data)
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

	// XXX: Should we allow .tpl files to be fetched as assets? Probably
	// not. For now, denying.
	if filepath.Ext(name) == ".tpl" {
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
			e.cache[d] = map[string]bool{}
			continue
		}

		e.cache[d] = make(map[string]bool, len(files))
		for _, f := range files {
			// Second half of cache key.
			r, err := filepath.Rel(d, f)
			if err != nil {
				return err
			}

			// TODO: Reading the file and then casting it to a string
			// doesn't feel like the right solution. But using ParseFiles
			// creates its own naming scheme, which doesn't work for us.
			data, err := ioutil.ReadFile(f)
			if err != nil {
				return err
			}

			// Assumption is that f is exactly the same as filepath.Join(d, r)
			if _, err := e.master.New(f).Parse(string(data)); err != nil {
				return err
			}

			e.cache[d][r] = true
		}
	}
	return nil
}
