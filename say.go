
package main

import "fmt"

func main() {
	var name string
	fmt.Print("I love people; \n What is your name?")
	fmt.Scan(&name)
	
	fmt.Println("Hello, ", name, "How are you? ")
}