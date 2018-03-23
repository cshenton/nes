package opt

import (
	"fmt"
	"math"

	"github.com/cshenton/nes/dist"
)

// Separable is a separable natural evolution strategies (SNES) optimizer.
// Unlike Plain, it always uses a diagonal normal population.
type Separable struct {
	*dist.DiagNormal
	LocRate   float64 // Learning rate
	ScaleRate float64 // Learning rate
	Size      int     // Generation size

	draws [][]float64 // Normalized draws this generation
	fits  []float64   // Fitnesses this generation
	count int
}

// NewSeparable creates a new SNES optimizer.
func NewSeparable(rate float64, size int, dim int) (s *Separable, err error) {
	if rate <= 0 || rate >= 1 {
		err = fmt.Errorf("learning rate must be in (0, 1), but was %v", rate)
		return nil, err
	}
	if size <= 0 {
		err = fmt.Errorf("generation size must be positive, but was %v", size)
		return nil, err
	}
	if dim <= 0 {
		err = fmt.Errorf("dimension must be positive, but was %v", dim)
		return nil, err
	}
	loc := make([]float64, dim)
	scale := make([]float64, dim)
	for i := range scale {
		scale[i] = 1e6
	}
	d, _ := dist.NewDiagNormal(loc, scale)

	s = &Separable{
		DiagNormal: d,
		LocRate:    rate,
		ScaleRate:  rate,
		Size:       size,
		draws:      make([][]float64, size),
		fits:       make([]float64, size),
	}
	return s, nil
}

// Update updates the search distribution given parameter, fitness.
func (s *Separable) Update(z []float64, f float64) {
	zn := make([]float64, len(z))
	for i := range zn {
		zn[i] = (z[i] - s.Loc[i]) / s.Scale[i]
	}

	s.draws[s.count] = zn
	s.fits[s.count] = f
	s.count++

	if s.count >= s.Size {
		u := utilities(s.fits, s.Size)
		dLoc := make([]float64, len(s.Loc))
		dScale := make([]float64, len(s.Scale))

		for i := range u {
			for j := range dLoc {
				dLoc[j] += u[i] * s.draws[i][j]
				dScale[j] += u[i] * (math.Pow(s.draws[i][j], 2) - 1)
			}
		}

		for j := range dLoc {
			dLoc[j] = s.LocRate * s.Scale[j] * dLoc[j]
			dScale[j] = math.Exp(s.ScaleRate * dScale[j] / 2)
		}

		for j := range s.Loc {
			s.Loc[j] += dLoc[j]
			s.Scale[j] *= dScale[j]
		}
	}
}
