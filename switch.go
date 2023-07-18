package main

import "fmt"

func main() {
	name := "Naufal"

	switch name {
	case "Naufal":
		fmt.Println("Hello Naufal")
		fmt.Println("Hello Naufal")
	case "Hakim":
		fmt.Println("Hello Hakim")
		fmt.Println("Hello Hakim")
	default:
		fmt.Println("Kenalan Donk")
	}

	// switch length := len(name); length > 5 {
	// case true:
	// 	fmt.Println("Nama terlalu panjang")
	// case false:
	// 	fmt.Println("Nama sudah benar")
	// }

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")
	}
}
