package main

import (
	"fmt"
	"upsxace/greetings"
)

func main(){
	message := greetings.Hello("Ace")
	fmt.Println(message)
}

// go mod edit -replace upsxace/greetings=../greetings
// go mod tidy