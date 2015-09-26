package form

import (
//"golang.org/x/net/html"
//"golang.org/x/net/html/atom"
)

type Label struct {
	HTML
	For, Form string
	//Content   html.Node
	Text string
}

func NewLabel(forName, text string) *Label {
	return &Label{
		For:  forName,
		Text: text,
	}
}
