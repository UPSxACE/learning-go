package main // Declare a main package (a package is a way to group functions, and it's made up of all the files in the same directory).

import "fmt" // Import the popular fmt package, which contains functions for formatting text, including printing to the console. 
             // This package is one of the standard library packages you got when you installed Go.

import "rsc.io/quote/v4"

/*
go mod tidy ensures that the go.mod file matches the source code in the module.
It adds any missing module requirements necessary to build the current module’s packages and dependencies,
and it removes requirements on modules that don’t provide any relevant packages. It also adds any missing entries to go.sum and
removes unnecessary entries.
*/

// Implement a main function to print a message to the console. A main function executes by default when you run the main package.
func main(){
	fmt.Println("Hello World!")
	fmt.Println(quote.Go())
}