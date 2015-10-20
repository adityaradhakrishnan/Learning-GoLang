package main

import "fmt"
import "math"

func main() {
	
	i := 1

	for i <= 10 {
		fmt.Println(i, math.Pow(2,float64(i)), math.Pow(2,1/float64(i)))
		i = i + 1
	}

	// A more standard initial condition, iterate to final condition loop

	for j := 3; j <= 9; j++ {
		fmt.Println(j*j)
	}

	k := 1

	for {
		if k > 8 {
			fmt.Println("k equals", k, "and the loop stops NOW!")
			break
		} else if k == 5 {
			fmt.Println("k equals", k, "?! That's just swell!")
			k++
		} else {
			fmt.Println("k equals", k, "and the loop goes around again...")
			k++
		}
	}
}