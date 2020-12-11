package main

import "fmt"

func main() {
	fmt.Println("Hello world")

	foo()

	for i := 1; i < 21; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
	bale()
	fmt.Println("exit the code")
}

func foo() {
	fmt.Println("I am in foo")
}

func bale() {
	fmt.Println("I am in bar")
}
