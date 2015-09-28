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

// FormElement describes any form element capable of expression as an html.Node.
type FormElement interface {
	Element() *html.Node
}

type OptionalBool uint8

const (
	// Nothing set, results in inheriting parent settings.
	ONone OptionalBool = iota
	// True
	OTrue
	// False
	OFalse
)

const (
	// Left-to-right
	LTR = "ltr"
	// Right-to-left
	RTL = "rtl"
	// Determine based on UA
	Auto = "auto"
)

// HTML captures a group of attributes common across all HTML elements.
//
// These attributes are all defined as Global, ARIA and Data attributes in
// the HTML5 specification. Because all of these can be applied to any
// form content, they are exposed here.
//
// The allowed values for all of these are explained in the HTML5 spec.
// Because we strive more for expression in the browser than semantic
// correctness, here and elsewhere we rarely force a particular value to
// conform to the spec. Typically, typing is as close as we get to
// enforcement.
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

// EnsureId ensures that an HTML has an ID attribute.
func (g HTML) EnsureId(seed string) string {
	if len(g.Id) > 0 {
		return g.Id
	} else if len(seed) > 0 {
		return seed
	}
	// TODO: Should probably generate a random ID.
	return ""
}

// Attache attaches these attributes to an html.Node.
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

// String is for PCData that can be arbitarily embeded in a []Field list.
type String string

// New creates a new form with the Name and Action fields set.
func New(name, action string) *Form {
	return &Form{Name: name, Action: action}
}

// Form describes an HTML5 form.
//
// A form can encapsulate an arbitrary number of form.Field objects.
//
// Forms can be create with the New() function, or instantiated directly.
// Then are typically rendered through the form templating system.
type Form struct {
	HTML
	AcceptCharset, Enctype, Action, Method, Name, Target string
	Autocomplete, Novalidate                             bool
	Fields                                               []Field
}

// Add adds any number of fields to a form.
func (f *Form) Add(field ...Field) *Form {
	f.Fields = append(f.Fields, field...)
	return f
}

// Element retrieves the form as an html.Node of type ElementNode.
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
