package form

import (
//"golang.org/x/net/html"
//"golang.org/x/net/html/atom"
)

// Password provides a field for obscured text.
type Password Input

// Text provides a single text entry line.
// See TextArea for multi-line input.
type Text Input

// Submit provides a button pre-wired for submission.
type Submit Input

// Tel provides a telephone number field.
type Tel Input

// URL provides a URL field.
type URL Input

// Email provides an Email address field.
type Email Input

// Date provides a date form entry field.
type Date Input

// Time provides a time form entry field.
type Time Input

// Number provides a numeric field.
type Number Input

// Range provides a number range field.
type Range Input

// Color provides a color picker.
type Color Input

// Checkbox provides a single checkbox.
type Checkbox Input

// Radio provides a single radio button.
type Radio Input

// File provides a file upload field.
type File Input

// Image provides a button that is painted with an image.
type Image Input

// Reset provides a button that is pre-wired to reset the form.
type Reset Input

// Button provides a generic button.
// This is deprecated in favor of the Button type.
type ButtonInput Input

// Hidden provides a field that will not be displayed.
type Hidden Input

// Input defines a generic untyped form field.
//
// It should not generally be used directly.
type Input struct {
	HTML
	Accept, Alt, Autocomplete, Dirname, Form, List, InputMode, Max, Min, MaxLength string
	Name, Pattern, Placeholder, Src, Step, Value                                   string
	Autofocus, Checked, Disabled, Multiple, ReadOnly, Required                     bool
	Height, Width, Size                                                            uint64

	// Technically, this is not an attribute of an Input field, but we put it here
	// to simplify the process of labeling fields.
	Label string
}

// Field describes any form element.
type Field interface{}
