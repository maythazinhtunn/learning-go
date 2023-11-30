package main

import "fmt"

func myfunc(strSlice *[]string) {
	(*strSlice)[5] = "F"
	(*strSlice)[7] = "H"
	(*strSlice)[9] = "J"
}
func main() {
	strSlice := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	myfunc(&strSlice)
	fmt.Println(strSlice)
}
