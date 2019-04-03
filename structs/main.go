package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

type contactInfo struct {
	phoneNo int
	email   string
}

func main() {

	//alex := person{firstName: "alex", lastName: "Anderson"}
	//fmt.Println(alex)

	// var alex person

	// alex.firstName = "alex"
	// alex.lastName = "Anderson"
	// fmt.Printf("%+v\n", alex)

	var alex person

	alex.firstName = "Alex"
	alex.lastName = "Anderson"
	alex.contact.email = "alex.anderson@gmail.com"
	alex.contact.phoneNo = 9876543210

	alex2 := person{firstName: "alex",
		lastName: "Anderson",
		contact: contactInfo{
			email:   "alex.anderson2@gmail.com",
			phoneNo: 9876543210,
		},
	}

	fmt.Printf("%+v", alex)
	fmt.Printf("%+v", alex2)
}
