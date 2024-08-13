package main

import "fmt"

type Addres struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia(address *Addres) {
	address.Country = "Malaysia"
}

func main() {
	alamat := Addres{
		City:     "Bekasi",
		Province: "Jawa Barat",
		Country:  "Indonesia",
	}
	ChangeCountryToIndonesia(&alamat)
	fmt.Println(alamat)
}
