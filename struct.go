package main

import "fmt"

type Customer struct {
	name string
	age  int8
}

func (customer Customer) sayHello() {

	fmt.Println(customer.name)

}

func main() {

	var person Customer

	person.name = "Sony Muhamad"
	person.age = 17

	//person2 := Customer{
	//	name: "",
	//	age:  0,
	//}

	person.sayHello()

}
