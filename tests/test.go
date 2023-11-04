package main

import (
	"fmt"
)

func sum(a, b string) string {
	return "hello"
}

func nakedReturnStatement(a,b int)(x,y int){
	x = a + b
	y = a * b
	return
}

var u int
var v string
var z bool
var i,o,p bool

var variableWithInitializer, var2, var3 = 1, true, []string{"a","b"}

var (
	varA string = "yo"
	varB int = 1343
)

var floatNum float64 = 5;


func main() {
	canOnlyDeclareThisWayInsideFunctions := 123456

	converted := float64(523)

	const Truth = true

	ac,ad := 1,2

	fmt.Println(ac,ad)

	fmt.Println("Go rules?", Truth)

	fmt.Println(sum("afsaf","asgasg"))
	fmt.Println(nakedReturnStatement(1,2))
	fmt.Println(variableWithInitializer, var2, var3 )
	fmt.Println(canOnlyDeclareThisWayInsideFunctions)
	fmt.Println(converted)
}


/*
Go basic types are:

bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128


The int, uint, and uintptr types are usually 32 bits wide on 32-bit systems and 64 bits wide on 64-bit systems.
When you need an integer value you should use int unless you have a specific reason to use a sized or unsigned
integer type.

*/


/*

Variables declared without an explicit initial value are given their zero value.

The zero value is:

0 for numeric types,
false for the boolean type, and
"" (the empty string) for strings.

*/