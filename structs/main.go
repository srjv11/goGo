package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {

	//alex := person{firstName: "Rajeev", lastName: "Shivaram"}
	//fmt.Println(alex)

	var rajeev person

	rajeev.firstName = "Rajeev"
	rajeev.lastName = "Shivaram"

	fmt.Printf("%+v\n", rajeev)
}
