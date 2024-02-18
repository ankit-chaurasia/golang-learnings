package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

// Defining Struct
type person struct {
	firstName string
	lastName  string
	// contact   contactInfo
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94555,
		},
	}
	jim.print()
	jim.updateName("New FirstName")
	jim.print()

	jimPointer := &jim
	jimPointer.updateNameWithPointer("New Firstname with Pointer")
	jimPointer.print()

	jim.updateNameWithPointer("Firstname with shortcut")
	jim.print()

	// // Declaring struct
	// // Method 1: but this approach will create issue if someone changes the order of struct fields
	// alex1 := person{"Alex1", "Anderson1"}
	// fmt.Println(alex1)
	// // Method 2: Recommded approach
	// alex2 := person{firstName: "Alex2", lastName: "Anderson2"}
	// fmt.Println(alex2)

	// // Method 3
	// var alex3 person

	// fmt.Println("Alex3 with default values", alex3)
	// fmt.Printf("%+v", alex3)
	// alex3.firstName = "Alex3"
	// alex3.lastName = "Anderson3"
	// fmt.Println("Alex3 after updating values", alex3)
}

// Pass by reference
func (pointerToPerson *person) updateNameWithPointer(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

// It will take p variable as a copy so updating any field will not update the original passed in person struct
func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
	fmt.Println()
}
