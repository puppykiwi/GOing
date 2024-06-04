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

func genRand(length int) string{
	var passwd string = ""
	var maxVar = 126
	var minVar = 48

	
	for i := 0; i < length; i++ {
		var randchar = rand.Intn(maxVar - minVar) + minVar
		//fmt.Println(randchar)
		passwd = passwd + string(rune(randchar))
		//fmt.Println(passwd)
	}
	//fmt.Println("Final : ", passwd)
	return (passwd)

}

func main(){
	var maxLength = getParams()
	// fmt.Print(e) //debug
	passwd := genRand(maxLength)
	fmt.Println("Your password is:", passwd,"\n")
}