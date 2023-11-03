package main

import (
	"fmt"
	"log"
	"upsxace/greetings"
)

func main(){
 	// Set properties of the predefined Logger, including
    // the log entry prefix and a flag to disable printing
    // the time, source file, and line number.
    log.SetPrefix("greetings: ")
    log.SetFlags(0)

    // A slice of names.
    names := []string {"Ace", "Eduardo", "Mark"}

    // Request greeting messages for the names.
	messages, err := greetings.HelloMultiple(names)
	// If an error was returned, print it to the console and
    // exit the program.
    if err != nil {
        log.Fatal(err)
    }

    // If no error was returned, print the returned messages
    // to the console.
    for _, message := range messages {
        fmt.Println(message)
    }
}

// go mod edit -replace upsxace/greetings=../greetings
// go mod tidy