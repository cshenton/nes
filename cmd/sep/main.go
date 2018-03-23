package main

import (
	"fmt"
	"log"
	"time"

	"github.com/cshenton/nes/opt"
)

func main() {
	rate := 0.5
	size := 10
	n := 1000

	o, err := opt.NewSeparable(rate, size, 1)
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
		if i%(n/10) == 0 {
			fmt.Printf("n: %v \t loc: %6.3f \t scale: %6.3f \n", i, o.Loc[0], o.Scale[0])
		}
	}
	fmt.Printf("%v iterations completed in %v \n", n, time.Now().Sub(t))
}
