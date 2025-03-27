package main

import "go-algorithm/data_structure"

func main() {
	println("Hello, Golang Algorithm ~")
	a := data_structure.NewArray([]int{}, 10)
	a.DeleteLast()
}
