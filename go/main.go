package main

import (
	"fmt"
	"go-algorithm/data_structure"
)

func main() {
	fmt.Println("Hello, golang")
	stack := data_structure.NewStack()
	p := stack.Pop()
	fmt.Printf("p: %v \n", p)
	fmt.Printf("p: %v \n", stack.Len())
}
