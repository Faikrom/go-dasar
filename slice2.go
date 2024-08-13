package main

import "fmt"

func main() {

	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}

	daySlice1 := days[5:] //sabtu minggu
	fmt.Println(daySlice1)

	daySlice1[0] = "Sabtu Baru"
	daySlice1[1] = "Minggu Baru"
	fmt.Println(daySlice1)
	fmt.Println(days)

	daySlice2 := append(daySlice1, "Libur Baru")
	daySlice2[0] = "Sabtu Lama"

	fmt.Println(daySlice1)
	fmt.Println(daySlice2)
	fmt.Println(days)

	fmt.Println("===========================")

	newSlice := make([]string, 2, 5)
	newSlice[0] = "Fitroh"
	newSlice[1] = "Ikrom"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	fmt.Println("===========================")

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

}
