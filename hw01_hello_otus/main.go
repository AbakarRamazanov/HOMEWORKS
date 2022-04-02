package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	text := "Hello, OTUS!"
	fmt.Print(stringutil.Reverse(text))
}
