package main
import (
	"fmt"

)
 func main(){
	if true && false {
		fmt.Println("true && false")
	} else if false || false {
		fmt.Println("false || false")
	} else if false && false {
		fmt.Println("false && false")
	} else if true || false {
		fmt.Println("true || false means true")
	} else {
		fmt.Println("This is the else")	}
 }