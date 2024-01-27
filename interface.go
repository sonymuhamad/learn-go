package main

import "fmt"

type HasName interface {
	getName() string
}

type Client struct {
	name string
}

func (client Client) getName() {
	fmt.Println("Hello my name is", client.name)
}

func main() {

	person := Client{
		name: "Sony",
	}

	person.getName()

}
