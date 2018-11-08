package main

import "fmt"

func main() {
	// INT
	myInt := 1       // 32 bit or 64 bit depending on OS
	var i8 int8 = -3 // same for 16, 32, 64
	var u8 uint8 = 3 // same for 16, 32, 64

	// FLOAT
	var f32 float32 = 3.2
	var f64 float64 = 6.4

	// COMPLEX
	cmplx64 := complex(f32, f32)
	var cmplx128 complex128 = 8 + 7i

	// STRING
	str := "My String"

	// BOOL
	myBool := true

	// Casting/converting
	var itof32 float64 = float64(myInt)

	fmt.Println(myInt, i8, u8, f32, f64, cmplx64, cmplx128, str, myBool, itof32)
}
