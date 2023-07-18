package main

import "fmt"

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

func main() {
	firstName := "Naufal"
	sayHelloTo(firstName, "Hakim")
	sayHelloTo("Kim", "Nopal")
}
