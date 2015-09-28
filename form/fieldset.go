package form

// FieldSet describes a set of form fields.
type FieldSet struct {
	HTML
	Form, Name string
	Disabled   bool
	Fields     []Field

	// This is not an attribute, but we should auto-generate the results.
	Legend string
}

// Divs are generic containers for fields.
//
// Because divs are frequently used to segment forms, we support them
// explicitly.
type Div struct {
	HTML
	Fields []Field
}
