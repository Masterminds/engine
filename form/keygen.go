package form

// Keygen describes the keygen form field type.
type Keygen struct {
	HTML
	Challenge, Form, KeyType, Name string
	Autofocus, Disabled            bool
}
