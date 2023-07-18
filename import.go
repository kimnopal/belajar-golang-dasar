package main

import (
	"belajar-golang-dasar/helper"
	"fmt"
)

func main() {
	helper.SayHello("Naufal")
	// helper.sayGoodbye("Naufal") // error
	fmt.Println(helper.Application)
	// fmt.Println(helper.version) // error
}
