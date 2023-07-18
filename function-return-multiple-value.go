package main

import "fmt"

func getFullName() (string, string) {
	return "Naufal", "Hakim"
}

func main() {
	firstName, _ := getFullName()
	fmt.Println(firstName)
	// fmt.Println(lastName)
}
