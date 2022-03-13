package main

import (
	"fmt"
	"github.com/bellwood4486/sandbox-go/embed/P"
)

type T1 struct {
	F1 int
}

func (t T1) Func1() {
	fmt.Println("Func1")
}

func (t *T1) Func1P() {
	fmt.Println("Func1P")
}

type T2 struct {
	F2 int
}

func (t T2) Func2() {
	fmt.Println("Func2")
}

func (t *T2) Func2P() {
	fmt.Println("Func2P")
}

type S struct {
	T1
	*T2
	P.T3
	*P.T4
	x int
}

func main() {
	t1 := T1{F1: 1}
	t2 := T2{F2: 2}
	t3 := P.T3{F3: 3}
	t4 := P.T4{F4: 4}

	s := S{
		T1: t1,
		T2: &t2,
		T3: t3,
		T4: &t4,
		x:  10,
	}

	s.Func1()
	s.Func1P()
	s.Func2()
	s.Func2P()
	s.Func3()
	s.Func3P()
	s.Func4()
	s.Func4P()

	before := t1.F1
	s.F1 = s.F1 * 2
	fmt.Printf("t1.F1: before:%d, after:%d\n", before, t1.F1)

	before = t2.F2
	s.F2 = s.F2 * 2
	fmt.Printf("t2.F2: before:%d, after:%d\n", before, t2.F2)
}
