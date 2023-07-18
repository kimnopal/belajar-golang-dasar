package main

import "fmt"

func main() {
	var name = "Naufal"

	if name == "Naufal" {
		fmt.Println("Hello Naufal")
	} else if name == "Hakim" {
		fmt.Println("Hello Hakim")
	} else if name == "Kim" {
		fmt.Println("Hello Kim")

	} else {
		fmt.Println("Hi, kenalan donk")
	}

	if length := len(name); length > 5 {
		fmt.Println("Terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
