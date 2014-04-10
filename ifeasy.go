package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		fmt.Println(v)
		return v

	}
	return lim
}
func main() {
	fmt.Println(
		pow(2, 2, 10),
		pow(3, 3, 20),
		pow(3, 4, 30),
	)
}
