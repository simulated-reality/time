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

// Duration computes the difference between Finish and Start.
func (self *Schedule) Duration() []float64 {
	time := make([]float64, len(self.Start))
	for i := range time {
		time[i] = self.Finish[i] - self.Start[i]
	}
	return time
}
