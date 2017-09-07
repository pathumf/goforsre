package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroprt(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroprt(&i)
	fmt.Println("zeroval:", i)

	fmt.Println("zeroval:", &i)

}
