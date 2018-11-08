package main

import "fmt"
import "strings"

// func name(args) returnType
func add(x int, y int) int {
	return x + y
}

// NOTE: This shorter arg list works as well
// func add(x, y int) int {
// 	return x + y
// }

// Functions can have multiple return values
func swap(x, y string) (string, string) {
	return y, x
}

// Functions can have named return values captured by the return statement
func separate(str string) (x, y string) {
	split := strings.Split(str, " ")
	x = split[0]
	y = split[1]
	return
}

func main() {
	fmt.Println(add(12, 34))
	fmt.Println(swap("Hello", "World"))
	fmt.Println(swap(separate("Hello, World!")))
}
