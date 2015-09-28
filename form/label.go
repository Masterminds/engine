package form

import (
//"golang.org/x/net/html"
//"golang.org/x/net/html/atom"
)

// Label describes a label for a field.
//
// Because labels are almost always applied to specific fields, and because
// the rules for attaching a label to a field vary by field type, more often
// than not you should favor a field's Label property over adding a
// Label element directly.
type Label struct {
	HTML
	For, Form string
	//Content   html.Node
	Text string
}

// NewLabel creates a new label.
func NewLabel(forName, text string) *Label {
	return &Label{
		For:  forName,
		Text: text,
	}
}
