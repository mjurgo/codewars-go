package main

import (
	"codewars/kyu6"
	"fmt"
)

func main() {
	fmt.Println(kyu6.FindUniq([]float32{1.0, 1.0, 1.0, 2.0, 1.0, 1.0}))
	fmt.Println(kyu6.FindUniq([]float32{0, 0, 0.55, 0, 0}))
}
