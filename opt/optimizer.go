package uv

// Evaler can evaluate a parameter and return an unscaled fitness.
type Evaler interface {
	Eval(p float64) (f float64)
}

// Optimizer defines a 1d sample based optimizer.
type Optimizer interface {
	// Samples a paramater candidate from the search distribution.
	Sample() (x float64)
	// Updates the search distribution given parameter, fitness.
	Update(x, f float64)
	// Mean returns the average parameter in the search distribution.
	Mean() (x float64)
}

// Optimize optimizes function e with optimizer o, over n samples, and returns
// the parameter and fitness that is reached.
func Optimize(e Evaler, o Optimizer, n int) (x, f float64) {
	for i := 0; i < n; i++ {
		x = o.Sample()
		f = e.Eval(x)
		o.Update(x, f)
	}

	x = o.Mean()
	f = e.Eval(x)

	return x, f
}
