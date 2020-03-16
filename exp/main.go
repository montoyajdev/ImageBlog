package main

import (
	"fmt"
)

type Dog struct{}

func (d Dog) Speak() {
	fmt.Println("woof")
}

type Husky struct {
	Speaker
}

type Speaker interface {
	Speak()
}

func main() {
	h := Husky{Dog{}}
	h.Speak()
}
