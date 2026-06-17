package main

import "fmt"

func isAdult(age int) bool {
	if age >= 18 {
		fmt.Println("you are eligible to drive")
		return true
	} else {
		fmt.Println("No you can't!")
		return false
	}
}

func main() {
	fmt.Println(isAdult(1))
}
