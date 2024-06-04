//Toddler Steps

package main

import "fmt"
import "math/rand"

func getParams() int {

	var length int
	fmt.Println("* Ultra Secure Password generator *")
	fmt.Print("How long would you like your passwd, kind hooman: ")
	fmt.Scan(&length)

	return (length)
}

