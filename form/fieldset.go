package form

type FieldSet struct {
	GlobalAttributes
	Form, Name string
	Disabled   bool
	Fields     []Field

	// This is not an attribute, but we should auto-generate the results.
	Legend string
}
