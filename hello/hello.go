package main

import (
	"fmt"

	"github.com/MelihCelik00/gostuff/greet/greetings"
)

func main(){
	msg:= greetings.Hello("Gladys")
	fmt.Println(msg)
}