package dist

// Distribution is a differentiable search distribution.
type Distribution interface {
	// Sample samples from the search distribution.
	Sample() (z []float64)
	// Params returns all the search distribution params in a flat slice.
	Params() (p []float64)
	// SearchGrads computes the partial derivatives of the loglikelihood w.r.t.
	// the distribution parameters.
	SearchGrads(z []float64) (pg []float64)
}
