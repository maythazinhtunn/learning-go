package main

import "fmt"

func main() {
	var username string
	var password string
	fmt.Print("Enter username:")
	fmt.Scanln(&username)
	fmt.Print("Enter password:")
	fmt.Scanln(&password)

	if username == "admin" && password == "password123" {
		fmt.Println("Login Successful!")
	} else {
		fmt.Println("Login Fail!")
	}
	// hw2
	fmt.Print("Enter three numbers:")
	var num1, num2, num3 int
	fmt.Scan(&num1)
	fmt.Scan(&num2)
	fmt.Scan(&num3)
	fmt.Print("Largest number:")
	if num1 > num2 && num1 > num3 {
		fmt.Printf("%d", num1)
	} else if num2 > num1 && num2 > num3 {
		fmt.Printf("%d", num2)
	} else if num3 > num1 && num3 > num2 {
		fmt.Printf("%d", num3)
	}
}
