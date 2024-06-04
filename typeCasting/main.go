// Go program to illustrate how to concatenate strings
// Using Sprintf function
package main

import (
	"fmt"
	"strconv"
)

func main() {

	// Creating and initializing strings
	str1 := "Tutorial"
	str2 := "of"
	str3 := "Go"
	str4 := "Language"

	result1 := str3 + " for " + str4

	result := fmt.Sprintf("%s %s %s %s", str1,
		str2, str3, str4)

	fmt.Println(result, result1)

	str4 = "12345678"

	num, err := strconv.Atoi(str4)

	fmt.Println(num, err)

}
