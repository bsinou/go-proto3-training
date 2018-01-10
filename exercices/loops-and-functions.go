package main

import (
	"fmt"
	"math"
)

func Sqrt1(x float64) {

	fmt.Println("With 10 iteration: ")

	z := float64(1)
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println("Iteration #", i, ": z =", z)
	}

	fmt.Println("Result: ", z)

}

func Sqrt2(x float64) {
	fmt.Println("With maximal distance")
	z := float64(1)
	currErr := float64(1)
	errMax := float64(0.00001)
	fmt.Println("errMax: ", errMax)
	i := 0
	for currErr > errMax {
		z -= (z*z - x) / (2 * z)
		currErr = math.Abs((z*z - x))
		fmt.Println("Iteration #", i, ": z =", z, ", curr error: ", currErr)
		i++
	}
	fmt.Println("Result after ", i, " iterations: ", z)
}

func sqrttest() {
	Sqrt1(2)
	Sqrt2(2)
}
