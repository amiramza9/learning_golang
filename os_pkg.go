package main

import (
	"fmt"
	"os"
	"reflect"
)

func main() {
	fmt.Println(os.Getpid())
	fmt.Println(os.Hostname())
	fmt.Println(reflect.TypeOf(os.ReadFile("./hello.go") any))

	//fmt.Println(os.ReadFile("./hello.go"))

}
