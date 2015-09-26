package form

type Select struct {
	HTML
	Autofocus, Disabled, Multiple, Required bool
	Form, Name                              string
	Size                                    uint64
	Options                                 []OptionItem
	Label                                   string
}

// DataList is a hidden option list used by other fields.
type DataList struct {
	HTML
	Options []Option
}

type OptionItem interface{}

type OptGroup struct {
	HTML
	Label    string
	Disabled bool
	Options  []Option
}

type Option struct {
	HTML
	Disabled, Selected bool
	// A label is the user-visible text, while the value is what is
	// sent to the server. Label may be rendered as phrasing content.
	Label, Value string
}
