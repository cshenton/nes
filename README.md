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
