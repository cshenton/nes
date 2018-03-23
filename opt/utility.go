package opt

import (
	"math"
	"sort"
)

// utilSlice is a generic sorting interface that preserves indices.
type utilSlice struct {
	sort.Interface
	perm []int
}

func (s utilSlice) Swap(i, j int) {
	s.Interface.Swap(i, j)
	s.perm[i], s.perm[j] = s.perm[j], s.perm[i]
}

// utilities computes rank normalized utilities given an arbitrary slice of fitnesses.
func utilities(f []float64, size int) (u []float64) {
	// Make the permutation slice
	perm := make([]int, len(f))
	for i := range perm {
		perm[i] = i
	}
	us := utilSlice{
		Interface: sort.Float64Slice(f),
		perm:      perm,
	}

	// Sort to get permutations
	sort.Sort(us)
	perm = us.perm

	// Calculate utilities
	s := float64(size)
	u = make([]float64, len(f))
	for i := range perm {
		u[perm[i]] = float64(i) // whatever was in position perm[i] is rank i
	}
	for i := range u {
		u[i] = math.Max(0, math.Log(s/2+1)-math.Log(u[i]))
	}
	uSum := 0.0
	for i := range u {
		uSum += u[i]
	}
	for i := range u {
		u[i] = u[i]/uSum - 1/s
	}

	return u
}
