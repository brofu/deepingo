package slice

import (
	"fmt"
)

type Test struct {
	A string
	B int
}

func StringSliceAddresses() {

	a := make([]string, 0, 2)
	fmt.Printf("a: %p, %p\n", a, &a)
	a = append(a, "a1")
	fmt.Printf("a: %p, %p\n", a, &a)

	b := make([]string, 2, 2)
	fmt.Printf("b: %p, %p\n", b, &b)
	b[0] = "b0"
	fmt.Printf("b: %p, %p\n", b, &b)
	b = append(b, "b2")
	fmt.Printf("b: %p, %p\n", b, &b)

}

func StringAddresses() {
	a := "test"
	fmt.Printf("a: %p\n", &a)
}

func IntegerSliceAddresses() {

	a := make([]int, 2, 2)
	fmt.Printf("a: %p, %p\n", a, &a)
	a[0] = 1
	fmt.Printf("a: %p, %p\n", a, &a)
	a = append(a, 3)
	fmt.Printf("a: %p, %p\n", a, &a)

}

func MapSliceAddresses() {

	a := make([]map[int]string, 2, 2)
	fmt.Printf("a: %p, %p\n", a, &a)
	a[0] = map[int]string{0: "0"}
	a[1] = map[int]string{1: "1"}
	fmt.Printf("a: %p, %p\n", a, &a)

}

func StructSliceAddress() {

	a := make([]Test, 2, 2)
	b := make([]*Test, 2, 2)

	s1 := Test{
		A: "s1",
		B: 1,
	}

	s2 := Test{
		A: "s2",
		B: 2,
	}

	a[0], a[1] = s1, s2
	b[0], b[1] = &s1, &s2

	fmt.Printf("%p, %p\n", a, &a)
	fmt.Printf("%p, %p\n", b, &b)
}
