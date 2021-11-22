package main

import (
	"fmt"
)

var (
	ara [12][12]float32
	b   byte
	sum float32 = 0.0
	p   int     = 5
	q   int     = 6
)

func main() {
	fmt.Scan(&b)

	for i := 0; i < len(ara); i++ {
		for j := 0; j < len(ara); j++ {
			fmt.Scan(&ara[i][j])
		}
	}

	for i := 0; i < len(ara); i++ {
		for j := p; j < q; j++ {
			sum += ara[i][j]
		}

		p--
		q++
	}

	if b == 'S' {
		fmt.Printf("%.1f\n", sum)
	} else {
		sum = (sum / 30.0)
		fmt.Printf("%.1f\n", sum)
	}
}
