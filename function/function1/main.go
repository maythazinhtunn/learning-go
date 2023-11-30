package main

import "fmt"

func greet(name string) (string, error) {
	return "Hello " + name + "! Welcome!", nil
}
func main() {
	greet_msg, err := greet("John")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(greet_msg)
}
