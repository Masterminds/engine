package form

type Button struct {
	HTML
	Autofocus, Disabled           bool
	Form, Menu, Name, Type, Value string
}

func NewButton(name, val string) *Button {
	return &Button{Name: name, Value: val}
}
