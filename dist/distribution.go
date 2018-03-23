package dist

// Distribution is a differentiable search distribution.
type Distribution interface {
	// Sample samples from the search distribution.
	Sample() (z []float64)
	Params() (p []float64)
	Gradients(z []float64) (pg []float64)
}
