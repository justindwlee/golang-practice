package main

import (
	"fmt"
	accounts "test/banking"
)

func main(){
	account := accounts.NewAccount("nico")
	fmt.Println(*account)
} 