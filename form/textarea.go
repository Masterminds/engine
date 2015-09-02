package form

type TextArea struct {
	GlobalAttributes
	Autocomplete, Dirname, Form, Name, Placeholder, Wrap string
	Autofocus, Disabled, ReadOnly, Required              bool
	Cols, MaxLength, MinLength, Rows                     uint64
}
