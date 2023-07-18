package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	var address1 Address = Address{"Purbalingga", "Jawa Tengah", "Indonesia"}
	var address2 *Address = &address1
	var address3 *Address = &address1

	address2.City = "Kebumen"

	*address2 = Address{"Malang", "Jawa Timur", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	var address4 *Address = new(Address)
	address4.City = "Jakarta"
	fmt.Println(address4)

	var alamat = Address{
		City:     "Purbalingga",
		Province: "Jawa Tengah",
		Country:  "",
	}

	var alamatPointer *Address = &alamat
	ChangeCountryToIndonesia(alamatPointer)
	fmt.Println(alamat)
}
