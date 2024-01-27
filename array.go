package main

import "fmt"

func main() {

	var fruits = [4]string{"apple", "grape", "banana", "melon"}
	fmt.Println(fruits)

	slice := fruits[:3]
	fmt.Println(slice)

	slice3 := make([]string, 3, 5)

	copy(slice3, slice)

	slice[0] = "Rambutan"

	fmt.Println(slice3)

	slice3[0] = "Mangga"

	fmt.Println(slice3)
	fmt.Println(fruits)

}
