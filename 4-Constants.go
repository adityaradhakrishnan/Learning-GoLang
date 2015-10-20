package main

import "fmt"
import "math"

const Name string = "Aditya Radhakrishnan"

func main() {
	fmt.Println(Name)

	const Pi = 3.1415926535

	fmt.Println(Pi * math.Pow(3,2))
	fmt.Println(Pi)
	fmt.Println(math.Sin(Pi))

	// Temporary casting to another type doesn't 
	// destroy the initial value.

	fmt.Println(float32(Pi))
	fmt.Println(Pi)
}