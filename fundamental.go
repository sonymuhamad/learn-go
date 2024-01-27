package main

import (
	"fmt"
	"strings"
)

func main() {

	str := "1255049"
	num := 0

	splittedString := strings.Split(str, "")

	for _, value := range splittedString {
		fmt.Println(value)

	}

	for {

		if num > 10 {
			break
		}

		num++
	}

	fmt.Println(num)

}
