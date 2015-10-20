// GoLang seems to require a specific
// package call for the main function.

package main

import "fmt"

func main() {
	fmt.Println("Go" + "Lang")
	fmt.Println("1 + 1 = ", 1+1)
	fmt.Println("7.0/3.0 = ", 7.0/3.0)
	fmt.Println("7/3 = ", 7/3)
	fmt.Println("7.0/3 = ", 7.0/3)
	fmt.Println("7/3.0 = ", 7/3.0)
	fmt.Println("You can make print statements across at least", 3, "items!")
}