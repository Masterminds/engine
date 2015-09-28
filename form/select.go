package form

// Select defines a selection list form element.
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

// OptionItem describes any item that can be a member of an options list.
//
// Select fields allow option items, while DataLists are more strict, and
// require Option types.
type OptionItem interface{}

// OptGroup describes a list of options.
type OptGroup struct {
	HTML
	Label    string
	Disabled bool
	Options  []Option
}

// Option describes an individual option in a selection, datalist, or optgroup.
type Option struct {
	HTML
	Disabled, Selected bool
	// A label is the user-visible text, while the value is what is
	// sent to the server. Label may be rendered as phrasing content.
	Label, Value string
}
