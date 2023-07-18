package main

import "fmt"

func getFullName2() (firstName string, lastName string) {
	firstName = "Naufal"
	lastName = "Hakim"

	return
}

func main() {
	a, b := getFullName2()
	fmt.Println(a)
	fmt.Println(b)
}
