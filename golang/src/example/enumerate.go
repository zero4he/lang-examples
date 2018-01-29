package main

import "fmt"

// 枚举
const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

// 枚举
const (
	A = "A"
	B = "B"
	C = 3
	D int = 4
)

func main() {
	fmt.Print(A, B, C)
}
