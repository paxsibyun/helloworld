package main

import (
	"fmt"
	"github.com/paxsibyun/stringutil"
)

func main() {
	fmt.Println("Hello World")
	val := stringutil.Reverse("ABCDEF")
	fmt.Println(val)
}

