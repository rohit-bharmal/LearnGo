package main

import "fmt"

// struct is a collection of fields which can be of different types and related to each other
type Person struct {
	Name        string
	Age         int
	ContactInfo ContactInfo // this is a nested struct
}

// ContactInfo is a struct that contains the contact information of a person
type ContactInfo struct {
	Email string
	Phone string
}

// you can create a new person by using the Person struct
func main() {
	person1 := Person{Name: "John", Age: 20}
	fmt.Println(person1.Name, person1.Age)

	// you can also create a new person by using the Person struct and the ContactInfo struct
	person2 := Person{Name: "Jane", Age: 21, ContactInfo: ContactInfo{Email: "jane@example.com", Phone: "1234567890"}}
	fmt.Println(person2.Name, person2.Age, person2.ContactInfo.Email, person2.ContactInfo.Phone)

	// you can also create a new person by using the new keyword
	person3 := new(Person)
	person3.Name = "Jim"
	person3.Age = 22
	fmt.Println(person3.Name, person3.Age)

	// you can also create a new person by using the & operator
	person4 := &Person{Name: "Jill", Age: 23}
	fmt.Println(person4.Name, person4.Age)

	// %+v is a format specifier that prints the struct in a human readable format and printf is a function that prints the string to the console
	fmt.Printf("%+v\n", person1)

	// update the name of the person using a pointer to a Person struct using a pointer receiver
	updateName(person4)
	fmt.Println(person4)
}

// updateName is a function that updates the name of a person using a pointer to a Person struct
func updateName(person *Person) {
	(*person).Name = "Rohit"
}
