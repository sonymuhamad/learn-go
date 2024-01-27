package main

import "fmt"

func testFunction(str string, blabla ...int) {
	fmt.Println(str)
	fmt.Println(blabla)
}

func main() {

	person := map[int8]string{
		5: "Lima",
		6: "Enem",
	}

	testmap := map[string]string{}

	person[7] = "Tujuh"

	fmt.Println(person)
	fmt.Println(testmap)

	testFunction("blabla", 5, 2, 1, 2, 5, 2)

}
