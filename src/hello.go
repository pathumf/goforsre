package main

import "fmt"

type sal struct {
	name     string
	greeting string
}

func greet(Sal sal) {
	fmt.Println(CreateMessage(Sal.name, Sal.greeting))
}

func CreateMessage(name, greeting string) string {
	return greeting + " " + name
}

func main() {
	s := sal{"Joe", "Hello"}
	greet(s)

}
