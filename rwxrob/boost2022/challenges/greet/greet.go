package greet

// to run use go install ./cmd/salut; salut or cmd/greet; greet

import "fmt"

func Hi(name string) {
	fmt.Printf("Hi, %v, nice to meet you!\n", name)
}

func Salut(name string) {
	fmt.Printf("Salut, %v, ravi de faire votre connaissance!\n", name)
}
