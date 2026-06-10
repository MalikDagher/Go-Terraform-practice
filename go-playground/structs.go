package main

import "fmt"

type Instance struct {
	ID    string
	CPU   int
	Mem   int
	State string
}

func main() {
	inst := Instance{
		ID:    "i-1234567",
		CPU:   2,
		Mem:   4,
		State: "running",
	}

	fmt.Println(inst.ID)
}
