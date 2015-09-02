package form

type Progress struct {
	GlobalAttributes
	Value, Max float64
}

type Meter struct {
	GlobalAttributes
	Value, Min, Max, Low, High, Optimum float64
}
