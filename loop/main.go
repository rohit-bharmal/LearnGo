package main

import "fmt"

func forLoop() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func whileLoop() {
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
}

func doWhileLoop() {
	i := 0
	for {
		fmt.Println(i)
		i++
		if i > 10 {
			break
		}
	}
}
func main() {
	forLoop()
	whileLoop()
	doWhileLoop()
}
