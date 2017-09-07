package main

import "fmt"

func puls(a int, b int) int {
	return a + b
}

func plusPuls(a, b, c int) int {
	return a + b + c
}

func main() {
	res := puls(1, 2)
	fmt.Println("1+2=", res)

	res = plusPuls(1, 2, 3)
	fmt.Println("1+2+3 =", res)
}
