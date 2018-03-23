# Natural Evolution Strategies
Replication of ES optimizers in [Natural Evolution Strategies](http://www.jmlr.org/papers/volume15/wierstra14a/wierstra14a.pdf)


## What is NES?

Evolutionary Strategies are gradient based evolutionary algorithms for optimising
over real-valued parameter spaces. Natural evolutionary strategies are an
adjustment to the vanilla approach (which just computes an empirical gradient
update) which scale the gradient by the fisher information matrix of the search
distribution.

This


## What is this repo?

This is an implementation of the vanilla and natural ES algorithms in Go, to both
replicate the results from the paper, and for use in other projects.


## To Do

- Interface for 1d ES style optimizers
- Benchmark interface / package (quadratic, sphere 10, rosenbrock 10)
- Vanilla ES algorithm, run benchmarks
- Canonical NES algorithm, run benchmarks (eta in {0.1, 0.5})
- NES: Fitness shaping
- NES: Adaptation sampling
- Multivariate NES interface
- ...


## API scratch

so we have search distributions
then the algos which are just methods for updating these distributions
opt params are z []float64, search params are p []float64, fitnesses are f float64
distributions expose Param() methods which return search params
also expose Sample(), Mean(), Adjust(), Gradients(z)

Optimizers wrap (and are instantiated with) a search distribution, they're the
ones that batch 'generations' and stagger updates.

Okay, so we have dist.Distribution interface, dist.DiagNormal, dist.NewDiagNormal(loc, scale []float64)

Then opt.Optimizer, opt.Plain{}, opt.NewPlain(d dist.Distribution), opt.Natural{},
opt.NewNatural(d dist.Distribution, s Shaper)

Then easy to go

```go
type foo struct {}

func(f foo) Eval(z float64) { return z*z }
```

As for fitness shaping, this is like a pre-processing rule, so maybe interface?
And adaptation sampling? Again this is like an intermediate step?