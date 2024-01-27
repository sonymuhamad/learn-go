package main

import (
	"errors"
	"fmt"
	"reflect"
)

func random() interface{} {

	return "Bla bla"

}

func main() {

	var customError = errors.New("Error bosku")

	fmt.Println(customError.Error())

	random := random()
	result := reflect.TypeOf(random)

	fmt.Println(result)

}
