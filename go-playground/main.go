package main

import (
	"errors"
	"fmt"
)

func validatePort(port int) (int, error) {
	if port < 1 || port > 65535 {
		return 0, errors.New("port out of range")
	} else {
		return port, nil
	}
}

func main() {
	p, err := validatePort(8080)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("valid port:", p)
	}
	d, err := validatePort(0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("valid port:", d)
	}
}
