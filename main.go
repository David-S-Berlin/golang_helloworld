package main

import (
	"fmt"
)

func main() {
	a := 8.0
	b := 4.0
	c := 3.0
	d := 5.0
	fmt.Println("Addition:\t", (a + b))
	fmt.Println("Subtraction:\t", (a - b))
	fmt.Println("Multiplication:\t", (a * b))
	fmt.Println("Multiplication:\t", (a * a * b))
	fmt.Println("Division:\t", (a / b / a))
	fmt.Println("Multiplication:\t", (a * a * b * a))
	fmt.Println("Addition:\t", (a + b + c + d + a + a + d + a))
	fmt.Println("Multiplication:\t", (a * b * c * d * a))
}
