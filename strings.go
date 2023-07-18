package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Naufal Hakim", "Naufal"))
	fmt.Println(strings.Contains("Naufal Hakim", "Nopal"))

	fmt.Println(strings.Split("Naufal Hakim", " "))

	fmt.Println(strings.ToLower("Naufal Hakim"))
	fmt.Println(strings.ToUpper("Naufal Hakim"))
	fmt.Println(strings.ToTitle("Naufal Hakim"))

	fmt.Println(strings.Trim("     Naufal Hakim     ", " "))
	fmt.Println((strings.ReplaceAll("Naufal Nopal Naufal", "Naufal", "Hakim")))
}
