package main

import "fmt"

func main(){
	x:= 5
	fmt.Println(&x)
	fmt.Println("before:", x)
	changeValue(&x)
	fmt.Println("after:", x)

	var y *int = &x
	fmt.Println("extra: ", y)
}

func changeValue(variable *int){
 *variable = 10
}