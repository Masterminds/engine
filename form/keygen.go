package form

type Keygen struct {
	HTML
	Challenge, Form, KeyType, Name string
	Autofocus, Disabled            bool
}
