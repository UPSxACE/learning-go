package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returns a greeting for the named person.
/*
This function takes a name parameter whose type is string. The function also returns a string.
In Go, a function whose name starts with a capital letter can be called by a function not in the same package.
This is known in Go as an exported name.
*/
func Hello(name string) (string, error){

	if(name == ""){
		return "", errors.New("empty name")
	}
	
	// Return a greeting that embeds the name in a message.
	msg := fmt.Sprintf(randomBaseText(), name)
	/*
	In Go, the := operator is a shortcut for declaring and initializing a variable in one line (Go uses the value on the right to determine the variable's type). Taking the long way, you might have written this as:

	var message string
	message = fmt.Sprintf("Hi, %v. Welcome!", name)
	*/
	return msg, nil
}

func randomBaseText() string{
	// A slice of message formats.
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	// Return a randomly selected message format by specifying
    // a random index for the slice of formats.
    return formats[rand.Intn(len(formats))]
}

