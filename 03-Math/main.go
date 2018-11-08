package main

import "fmt"

func main() {
	// NOTE: For more complex math use the math pkg.

	int1, int2 := 9, 4

	// Addition
	fmt.Println("9 + 4 =", int1+int2)
	// Subtraction
	fmt.Println("9 - 4 =", int1-int2)
	// Multiplication
	fmt.Println("9 * 4 =", int1*int2)
	// Division
	fmt.Println("9 / 4 =", int1/int2)
	fmt.Println("f64: 9 / 4 =", float64(int1)/float64(int2))
	// Modulus
	fmt.Println("9 % 4 =", int1%int2)
}
