package main

import "fmt"

func Ups(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "Ups"
	}
}

func main() {
	var data1 interface{} = Ups(1)
	fmt.Println(data1)

	var data2 interface{} = Ups(2)
	fmt.Println(data2)

	var data3 interface{} = Ups(3)
	fmt.Println(data3)
}
