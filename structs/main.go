package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	brad := person{
		firstName: "Brad",
		lastName:  "Schmitt",
		contactInfo: contactInfo{
			email:   "bschmitt@someemail.com",
			zipCode: 99999,
		
	}

	// bradPointer := &brad
	// shortcut below.  receiver uses pointer because *person is specified.
	brad.updateName("Bradley")
	brad.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

// Receiver is a pointer that points at a person
// Address -> Value with *address
// Value -> Address with &value
func (pointerToPerson *person) updateName(newFirstName string) {
	// An operator that manipulates the pointer that is being referenced
	(*pointerToPerson).firstName = newFirstName
}
