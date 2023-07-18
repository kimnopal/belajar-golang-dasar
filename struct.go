package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHi(name string) {
	fmt.Println("Hello", name, "My Name is", customer.Name)
}

func (a Customer) sayHuuu() {
	fmt.Println("Huuuuuu from", a.Name)
}

func main() {
	var naufal Customer
	naufal.Name = "Naufal"
	naufal.Address = "Indonesia"
	naufal.Age = 20

	naufal.sayHi("Hakim")
	naufal.sayHuuu()

	// fmt.Println(naufal.Name)
	// fmt.Println(naufal.Address)
	// fmt.Println(naufal.Age)

	// hakim := Customer{
	// 	Name:    "Hakim",
	// 	Address: "Indonesia",
	// 	Age:     21,
	// }
	// fmt.Println(hakim)

	// kim := Customer{"Kim", "Indonesia", 22}
	// fmt.Println(kim)
}
