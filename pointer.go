package main

import "fmt"

type Address struct {
	Country, City, Street string
}

func (address *Address) changeStreetToParman() {
	address.Street = "S. Parman"
}

func main() {

	address1 := Address{
		Country: "Indonesia",
		City:    "Jakarta",
		Street:  "Gatsu",
	}

	address2 := &address1

	address2.Street = "DI Panjaitan"

	fmt.Println(address1, "1")
	fmt.Println(address2, "2")

	// #############################################3

	fmt.Println("End Region")

	address5 := new(Address)
	address4 := address5

	address4.Country = "Singapore"
	address4.City = "Bandung"

	address4.changeStreetToParman()
	fmt.Println(address4)
}
