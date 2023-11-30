package main

import (
	"errors"
	"fmt"
)

func sumAndDifference(num1, num2 int) (int, int, error) {
	sum := num1 + num2
	difference := num1 - num2
	if (num1 > 0 && num2 < 0 && difference < 0) || (num1 < 0 && num2 > 0 && difference > 0) {
		return 0, 0, errors.New("integer overflow in differrnce")
	}
	return sum, difference, nil
}
func main() {
	resultSum, resultDiff, err := sumAndDifference(1, -30)
	if err != nil {
		fmt.Println("ERROR", err)
		return
	}
	fmt.Println("Sum:", resultSum)
	fmt.Println("Difference:", resultDiff)
}
