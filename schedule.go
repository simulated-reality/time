package time

// Schedule represents a schedule of an application on a platform.
type Schedule struct {
	Cores uint
	Tasks uint
	Span  float64

	Mapping []uint
	Order   []uint
	Start   []float64
	Finish  []float64
}
