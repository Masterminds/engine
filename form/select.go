package form

type Select struct {
	GlobalAttributes
	Autofocus, Disabled, Multiple, Required bool
	Form, Name                              string
	Size                                    uint64
	Options                                 []OptionItem
}

type OptionItem interface{}

type OptGroup struct {
	GlobalAttributes
	Label    string
	Disabled bool
	Options  []Option
}

type Option struct {
	GlobalAttributes
	Disabled, Selected bool
	Label, Value       string
}
