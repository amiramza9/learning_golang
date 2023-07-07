package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now().UTC()
	fmt.Println(now.Clock())
	fmt.Println(now.YearDay())
	fmt.Println(now.Month())
	fmt.Println(now.Unix())
	fmt.Println(now)
	fmt.Println(now)
	fmt.Println(now)
	fmt.Println(time.Now())
}
