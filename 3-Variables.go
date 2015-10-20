package main

import "fmt"
import "math"

func main() {
	
	var Name string = "Aditya Radhakrishnan"
	fmt.Println(Name)

	var IntA, IntB int = 1, 2
	fmt.Println(IntA, IntB)
	fmt.Println(IntA + IntB)

	var BoolCheck  = true
	var BoolCheckB = (IntA == 1) 

	fmt.Println(BoolCheck, BoolCheckB)

	// Variables that are declared without
	// initialization are zero-valued.

	var A int
	var B string
	var C bool

	// var D float gives the error:
	// "undefined: float"
	// 
	// Instead, use:

	var D float32
	var E float64

	fmt.Println(A,B,C,D,E)

	// Syntactical shorthand for declaring and
	// initializing a variable - very Pythonic

	Text := "Short!"
	fmt.Println(Text)

	Pi     := 3.1415926535
	Radius := 3.0
	Area   := Pi * math.Pow(Radius,2)

	fmt.Println(Pi*Radius, Area)

}