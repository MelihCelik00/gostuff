package main

import (
	"fmt"
	
	"example.com/greetings"
)

func main(){
	msg:= greetings.Hello("Melih")
	fmt.Println(msg)
}