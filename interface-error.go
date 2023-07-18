package main

import (
	"errors"
	"fmt"
)

// type ErrorPembagi struct {
// 	message string
// }

// func (errorPembagi ErrorPembagi) Error() string {
// 	return errorPembagi.message
// }

// func Pembagi(nilai int, pembagi int) (int, error) {
// 	if pembagi == 0 {
// 		return 0, ErrorPembagi{message: "error cuy"}
// 	} else {
// 		result := nilai / pembagi
// 		return result, nil
// 	}
// }

func Pembagi(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi tidak boleh 0")
	} else {
		result := nilai / pembagi
		return result, nil
	}
}

func main() {
	hasil, err := Pembagi(100, 0)
	if err == nil {
		fmt.Println("Hasil", hasil)
	} else {
		fmt.Println(err.Error())
	}
}
