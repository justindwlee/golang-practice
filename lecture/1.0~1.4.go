package main

import (
	"fmt"
	"strings"
)

//go is a strongly typed lang
func multiply(a ,b int) int{
    return a * b
}

//returning multiple variables
func lenAndUpper(name string) (int, string) {
    return len(name), strings.ToUpper(name)
}

//variadic parameter
func repeatMe(words ...string) {
    fmt.Println(words)
}

//naked return / defer
func lenAndUpper2(name string) (length int, uppercase string) {
    defer fmt.Println("I'm done")
    length = len(name)
    uppercase = strings.ToUpper(name)
    return
}

