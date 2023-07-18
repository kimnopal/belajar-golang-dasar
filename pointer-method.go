package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	naufal := Man{"Naufal"}
	naufal.Married()

	fmt.Println(naufal.Name)
}
