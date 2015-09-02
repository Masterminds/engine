package form

type Keygen struct {
	GlobalAttributes
	Challenge, Form, KeyType, Name string
	Autofocus, Disabled            bool
}
