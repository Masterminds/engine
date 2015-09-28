package form

// Progress defines a progress meter form element type.
type Progress struct {
	HTML
	Value, Max float64
}

// Meter defines a generic meter form element type.
type Meter struct {
	HTML
	Value, Min, Max, Low, High, Optimum float64
}
