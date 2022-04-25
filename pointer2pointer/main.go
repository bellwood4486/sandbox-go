package main

import "fmt"

// see: https://www.geeksforgeeks.org/go-pointer-to-pointer-double-pointer/
func main() {
	var v int = 100

	var pt1 *int = &v

	var pt2 **int = &pt1

	fmt.Println("The Value of Variable v is = ", v)
	fmt.Println("Address of variable v is = ", &v)

	fmt.Println("The Value of pt1 is = ", pt1)
	fmt.Println("Address of pt1 is = ", &pt1)

	fmt.Println("The Value of pt2 is = ", pt2)

	fmt.Println("Value at the address of pt2 is or *pt2 = ", *pt2)

	fmt.Println("*(Value at the address of pt2 is) or **pt2 = ", **pt2)
}
