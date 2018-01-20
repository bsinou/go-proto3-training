package simple

import (
	"fmt"
)

func fizzbuzz(n int) {
	for i := 1; i <= n; i++ {
		switch {
		case i%3 == 0 && i%5 == 0: // equivalent to i%15 == 0 :
			fmt.Println("buzzfizz")
		case i%3 == 0:
			fmt.Println("buzz")
		case i%5 == 0:
			fmt.Println("fizz")
		default:
			fmt.Println(i)
		}
	}
}
