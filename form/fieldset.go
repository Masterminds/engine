package form

type FieldSet struct {
	HTML
	Form, Name string
	Disabled   bool
	Fields     []Field

	// This is not an attribute, but we should auto-generate the results.
	Legend string
}
