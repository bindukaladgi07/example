package main

import "fmt"

func main() {
	n, _ := fmt.Println("Hello world, 34, true")
	fmt.Println(n)
	//fmt.Println(err)
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

// control flow:
// (1) sequence
// (2) loop; iterative
// (3) conditional