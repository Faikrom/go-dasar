package main

import "fmt"

func sayHai(name string, umur int, alamat string) string {
	return fmt.Sprintf("Hai, nama saya %s, umur saya %d tahun, alamat saya di %s", name, umur, alamat)
}

func main() {
	var identitas string = sayHai("Fitroh", 21, "Tangerang")
	fmt.Println(identitas)
}
