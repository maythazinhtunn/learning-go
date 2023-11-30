package main

import (
	"fmt"
)

func swap(a string, b string) (string, string, error) {
	return b, a, nil
}

func main() {
	str2, str1, err := swap("Hello", "World")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(str2, str1)
}
