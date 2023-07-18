package main

import "fmt"

func main() {
	type NoTKP string
	type Married bool

	var noKtpNaufal NoTKP = "121231231231"
	var marriedStatus Married = false

	fmt.Println(noKtpNaufal)
	fmt.Println(marriedStatus)
}
