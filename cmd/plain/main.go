package main

import (
	"fmt"
	"log"
	"math"
	"time"

	"github.com/cshenton/nes/dist"
	"github.com/cshenton/nes/opt"
)

type quad struct{}

func (q *quad) Eval(z []float64) (f float64) { return math.Pow(z[0], 2) }

func main() {
	lr := 0.001
	popSize := 10
	n := 100000

	d, err := dist.NewDiagNormal([]float64{0}, []float64{100})
	if err != nil {
		log.Fatal(err)
	}

	o, err := opt.NewPlain(d, lr, popSize)
	if err != nil {
		log.Fatal(err)
	}

	var (
		z []float64
		f float64
	)

	eval := func(z []float64) float64 { return -z[0] * z[0] }

	t := time.Now()
	for i := 0; i < n; i++ {
		z = o.Sample()
		f = eval(z)
		o.Update(z, f)
		if i%(n/50) == 0 {
			fmt.Printf("n: %v \t loc: %6.3f \t scale: %6.3f \n", i, d.Loc[0], d.Scale[0])
		}
	}
	fmt.Printf("%v iterations completed in %v \n", n, time.Now().Sub(t))
}
