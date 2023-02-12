package main

import (
	"credentity/generated"
	"fmt"
)

func main() {
	var g = &generated.Gateway{
		Name: "upi",
	}
	fmt.Println(g)
}
