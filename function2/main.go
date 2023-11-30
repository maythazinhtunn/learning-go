package main

import (
	"errors"
	"fmt"
)

func calculateSquare(num int) (int, error) {
	if num < 0 {
		return 0, errors.New("input must be a non-negative integer")
	}
	return num * num, nil
}
func main() {
	result, err := calculateSquare(-5)
	if err != nil {
		fmt.Println("ERROR", err)
		return
	}
	fmt.Println(result)
}
