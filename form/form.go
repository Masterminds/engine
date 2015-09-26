/* Package form provides utilities for creating and accessing HTML forms.

This is an abstraction for defining forms in text, and accessing them
accordingly. It is partially inspired by Drupal's form library.

This generates HTML5 forms. http://www.w3.org/TR/html5/forms.html
*/
package form

import (
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

type FormElement interface {
	Element() *html.Node
}

type OptionalBool uint8

const (
	ONone OptionalBool = iota
	OTrue
	OFalse
)

const (
	LTR  = "ltr"
	RTL  = "rtl"
	Auto = "auto"
)

// Attributes common across all HTML element.
type HTML struct {
	Class                                                       []string
	AccessKey, Id, Dir, Lang, Style, TabIndex, Title, Translate string
	ContentEditable, Hidden                                     OptionalBool
	Role                                                        string

	// Data stores arbitrary attributes, such as data-* fields. It is up to
	// the implementation to know how to deal with these fields.
	Data map[string]string

	// Attributes prefixed "aria-"
	Aria map[string]string
}

func (g HTML) EnsureId(seed string) string {
	if len(g.Id) > 0 {
		return g.Id
	} else if len(seed) > 0 {
		return seed
	}
	// TODO: Should probably generate a random ID.
	return ""
}

func (g HTML) Attach(node *html.Node) {
	attrs := []html.Attribute{}
	if g.ContentEditable > 0 {
		if g.ContentEditable == OTrue {
			attrs = attr(attrs, "contenteditable", "true")
		} else {
			attrs = attr(attrs, "contenteditable", "false")
		}
	}
	if g.Hidden > 0 {
		if g.Hidden == OTrue {
			attrs = attr(attrs, "hidden", "true")
		} else {
			attrs = attr(attrs, "hidden", "false")
		}
	}

	if len(g.Data) > 0 {
		for k, v := range g.Data {
			attrs = attr(attrs, k, v)
		}
	}

	if len(g.Class) > 0 {
		v := strings.Join(g.Class, " ")
		attrs = attr(attrs, "class", v)
	}

	s := []string{"AccessKey", "Id", "Dir", "Lang", "Style", "TabIndex", "Title", "Translate"}
	attrs = append(attrs, structToAttrs(g, s...)...)

	node.Attr = append(node.Attr, attrs...)
}

// Divs appear frequently in forms, so we're punting and supporting them.
type Div struct {
	HTML
	Fields []Field
}

// String is for PCData that can be arbitarily embeded in a []Field list.
type String string

func New(name, action string) *Form {
	return &Form{Name: name, Action: action}
}

type Form struct {
	HTML
	AcceptCharset, Enctype, Action, Method, Name, Target string
	Autocomplete, Novalidate                             bool
	Fields                                               []Field
}

func (f *Form) Add(field ...Field) *Form {
	f.Fields = append(f.Fields, field...)
	return f
}

func (f *Form) Element() *html.Node {
	n := &html.Node{
		Type:     html.ElementNode,
		DataAtom: atom.Form,
		Data:     "form",
	}

	n.Attr = structToAttrs(f, "AcceptCharset", "Enctype", "Action", "Method", "Name", "Target")

	// We want to at least try to set an ID.
	f.HTML.Id = f.HTML.EnsureId(f.Name)
	f.HTML.Attach(n)

	return n
}
