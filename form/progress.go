package form

type Progress struct {
	HTML
	Value, Max float64
}

type Meter struct {
	HTML
	Value, Min, Max, Low, High, Optimum float64
}
