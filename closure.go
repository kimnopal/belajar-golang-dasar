package main

import "fmt"

func main() {
	name := "Naufal"
	counter := 0

	increment := func() {
		name := "Hakim"
		fmt.Println("Increment")
		counter++
		fmt.Println(name)
	}

	increment()
	increment()

	fmt.Println(counter)
	fmt.Println(name)
}
