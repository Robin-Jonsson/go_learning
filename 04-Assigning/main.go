package main

import "fmt"

func main() {
	// Single assign
	var v1 int = 1
	var v2 = 2
	v3 := 3

	const v4 int = 4
	const v5 = 5

	// Multi assign
	var v6, v7, v8 = 6, 7, 8
	v9, v10 := 9, 10

	var (
		v11 = 11
		v12 = 12
	)

	const (
		v13 = 13
		v14 = 14
	)

	fmt.Println(v1, v2, v3, v4, v5, v6, v7, v8, v9, v10, v11, v12, v13, v14)
}
