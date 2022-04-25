package main

import "fmt"

func example1() {
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

func example2() {
	var v int = 100
	var pt1 *int = &v
	var pt2 **int = &pt1

	fmt.Println("The value of variable v is = ", v)

	*pt1 = 200

	fmt.Println("Value stored in v after changing pt1 = ", v)

	**pt2 = 300

	fmt.Println("Value stored in v after changing pt2 = ", v)
}

// see: https://www.geeksforgeeks.org/go-pointer-to-pointer-double-pointer/
func main() {
	example1()
	example2()
}
