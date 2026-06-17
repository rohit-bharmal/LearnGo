package main

import "fmt"

func main() {
	// you can declare a variable with a type
	var name string = "John"
	var age int = 20
	var isStudent bool = true
	var height float64 = 1.75

	fmt.Println("Name:", name, "Age:", age, "Is Student:", isStudent, "Height:", height)

	// you can declare a variable without a type
	names := "John"
	ages := 20
	isStudents := true
	heights := 1.75
	fmt.Println("Name:", names, "Age:", ages, "Is Student:", isStudents, "Height:", heights)

}
