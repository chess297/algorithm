package learn_go

import "fmt"

type Variable struct {
	// items []int
}

func (v *Variable) log() {
	fmt.Printf("Hello, %s Algorithm ~", GO_NAME)
}

func NewVariable() Variable {
	v := Variable{}
	v.log()

	return v
}
