package main

import (
	"fmt"
)

func runFirst() {
	fmt.Println("first")
}

func runLast() {
	fmt.Println("second")
}

func main() {
	defer runFirst()
	runLast()
}