package main
import "fmt"
import "os"

func main(){
	arguments := os.Args
	args := arguments[1:]
	fmt.Println("Slice: ", args)
	for _, arg := range args {
		fmt.Println(arg)
	}
}

