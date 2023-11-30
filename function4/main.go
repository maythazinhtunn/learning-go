package main

import (
	"errors"
	"fmt"
)

func devideWithRemainder(num1, num2 int) (int, int, error) {
	if num2 == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	}
	quotient := num1 / num2
	remainder := num1 % num2

	return quotient, remainder, nil
}
func main() {
	resQuotient, resRemainder, err := devideWithRemainder(15, 10)
	if err != nil {
		fmt.Println("ERROR", err)
		return
	}
	fmt.Println("Quotient:", resQuotient)
	fmt.Println("Remainder:", resRemainder)
}
