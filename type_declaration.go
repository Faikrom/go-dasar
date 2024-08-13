package main

import "fmt"

func main() {
	type noKTP string

	var KTPfitroh noKTP = "111111111111"
	var contoh = "222222222222"
	var contohKTP noKTP = noKTP(contoh)

	fmt.Println(KTPfitroh)
	fmt.Println(contohKTP)

}
