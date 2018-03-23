package opt

import (
	"fmt"

	"github.com/cshenton/nes/dist"
)

// Plain is a plain evolutionary strategies optimizer.
type Plain struct {
	Population dist.Distribution // Population distribution
	Rate       float64           // Learning rate
	Size       int               // Generation size

	update []float64 // Parameter update (adjusted incrementally)
	count  int       // Generation progress counter
}

// NewPlain creates a new plain ES optimizer, and returns an error if the
// learning rate or population are invalid.
func NewPlain(pop dist.Distribution, rate float64, size int) (p *Plain, err error) {
	if rate <= 0 || rate >= 1 {
		err = fmt.Errorf("learning rate must be in (0, 1), but was %v", rate)
		return nil, err
	}
	if size <= 0 {
		err = fmt.Errorf("generation size must be positive, but was %v", size)
		return nil, err
	}

	p = &Plain{
		Population: pop,
		Rate:       rate,
		Size:       size,
		update:     make([]float64, len(pop.Params())),
	}
	return p, nil
}

// Sample returns a parameter samples from the current search distribution.
func (p *Plain) Sample() (z []float64) { return p.Population.Sample() }

// Update updates the search distribution in response to an observed fitness.
// This will just accumulate gradients until the generation size is reached at
// which point it applies and resets them.
func (p *Plain) Update(z []float64, f float64) {
	s := p.Population.SearchGrads(z)
	for i := range p.update {
		p.update[i] += s[i] * f / float64(p.Size)
	}
	p.count++

	if p.count >= p.Size {
		p.Population.Apply(p.update)
		p.update = make([]float64, len(p.Population.Params()))
		p.count = 0
	}
}

// Mean returns the mean of the search distribution.
func (p *Plain) Mean() (z []float64) { return p.Population.Mean() }
