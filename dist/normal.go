package dist

import (
	"errors"
	"fmt"
	"math/rand"
)

// DiagNormal is a diagonal multinormal distribution.
type DiagNormal struct {
	Loc   []float64
	Scale []float64
}

// NewDiagNormal returns a new diagonal normal, or an error if the input
// loc, scale vectors are invalid.
func NewDiagNormal(loc, scale []float64) (n *DiagNormal, err error) {
	if len(loc) == 0 {
		err = errors.New("loc, scale, must have non zero length")
		return nil, err
	}
	if len(loc) != len(scale) {
		err = fmt.Errorf("loc, scale must be same length, but were len %v and %v", len(loc), len(scale))
		return nil, err
	}
	for i := range scale {
		if scale[i] <= 0 {
			err := fmt.Errorf("scale must be strictly positive, but was %v at pos %v", scale[i], i)
			return nil, err
		}
	}

	n = &DiagNormal{
		Loc:   loc,
		Scale: scale,
	}
	return n, nil
}

// Sample draws from the diagonal normal distribution.
func (n *DiagNormal) Sample() (z []float64) {
	z = make([]float64, len(n.Loc))

	for i := range z {
		z[i] = n.Loc[i] + n.Scale[i]*rand.NormFloat64()
	}

	return z
}
