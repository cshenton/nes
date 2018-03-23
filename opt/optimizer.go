package opt

// Evaler can evaluate a parameter and return an unscaled fitness.
type Evaler interface {
	Eval(z []float64) (f float64)
}

// Optimizer defines a 1d sample based optimizer.
type Optimizer interface {
	// Samples a paramater candidate from the search distribution.
	Sample() (z []float64)
	// Updates the search distribution given parameter, fitness.
	Update(z []float64, f float64)
	// Mean returns the average parameter in the search distribution.
	Mean() (z []float64)
}

// Optimize optimizes function e with optimizer o, over n samples, and returns
// the parameter and fitness that is reached.
func Optimize(e Evaler, o Optimizer, n int) (z []float64, f float64) {
	for i := 0; i < n; i++ {
		z = o.Sample()
		f = e.Eval(z)
		o.Update(z, f)
	}

	z = o.Mean()
	f = e.Eval(z)

	return z, f
}
