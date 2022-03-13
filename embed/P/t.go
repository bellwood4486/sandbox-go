package P

import "fmt"

type T3 struct {
	F3 int
}

func (t T3) Func3() {
	fmt.Println("Func3")
}

func (t *T3) Func3P() {
	fmt.Println("Func3P")
}

type T4 struct {
	F4 int
}

func (t T4) Func4() {
	fmt.Println("Func4")
}

func (t *T4) Func4P() {
	fmt.Println("Func4P")
}
