package main

import (
	"errors"
	"fmt"
)

func Pembagi(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi tidak boleh Nol")
	} else {
		result := nilai / pembagi
		return result, nil
	}
}

func main() {
	hasil, err := Pembagi(100, 0)
	if err == nil {
		fmt.Println("Hasil :", hasil)
	} else {
		fmt.Println("Error :", err.Error())
	}

	// hasil2, err2 := Pembagi(100, 0)
	// if err == nil {
	// 	fmt.Println("Hasil :", hasil2)
	// } else {
	// 	fmt.Println("Error :", err2.Error())
	// }
}
