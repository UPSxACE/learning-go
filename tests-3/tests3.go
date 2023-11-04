package main

import "fmt"

func main(){
	// for loop
	for i:= 0; i < 10; i++{
		fmt.Println(i+1)
	}

	fmt.Println("End for loop")

	i, j := 0, false
	// while loop
	for j == false {
		fmt.Println(i)
		i++
		if i >= 3 {
			j = true
		}
	}

	fmt.Println("End while loop")

	number := 2

	switch number {
		case 1: fmt.Println("One")
		case 2: fmt.Println("Two")
		case 3: fmt.Println("Three")
		default: fmt.Println("No")
	}

	fmt.Println("End switch")

	array := [5]int{1,2,3}

	for _, num := range array {
		fmt.Println(num)
	}

	fmt.Println("End of array iteration with a for loop")

	slicedArray := array[1:4]

	fmt.Println(slicedArray)
	fmt.Println("End of sliced array")

	manualSlice := make([]int, 1, 2);
	fmt.Println("#1:",manualSlice)
	manualSlice = manualSlice[:2] // reslicing to have length 2
	//manualSlice = manualSlice[:3] // cannot reslice to 3 because that goes beyond its capacity
	fmt.Println("#2:",manualSlice)
	manualSlice[0], manualSlice[1] = 3, 4
	fmt.Println("#3:",manualSlice)

	updatedSlice := append(manualSlice, 10)
	fmt.Println("updated:",updatedSlice)

	numMap := make(map[string]int)

	numMap["one"] = 1;
	numMap["two"] = 2;

	fmt.Println(numMap)
	fmt.Println(len(numMap))

	numMap2 := map[string]string {
		"key": "value",
	}

	fmt.Println(numMap2)

	confusingMap := map[string]map[string]string {
		"map1": {
			"key": "value",
		},
		"map2": {
			"key": "value",
		},
	}

	fmt.Println(confusingMap)

	fmt.Println("End of slices and maps")

	defer fmt.Println("First defer")
	defer fmt.Println("Second defer")
	defer fmt.Println("Hello")
	fmt.Println("Bye")


	



	defer handlePanic()

	

	dog1 := Dog{"Cookie", func()string{
		return "*bark*"
	}}

	fmt.Println(dog1)
	fmt.Println(dog1.bark())

	dog1.whoDog()
	dog1.whoDog()

	rectangle := Rectangle{50,100}
	areaOfShape := getAreaOfShape(&rectangle)
	fmt.Println(areaOfShape)

	fmt.Println("End")
	
}

func handlePanic(){
	x := recover()
	if x != nil {
		// panic happened
		fmt.Println("YOOOOO")
		fmt.Println("The error was:", x)
	}
}

func (doggo *Dog) whoDog(){
	fmt.Println("Dog name is", doggo.name)
	fmt.Println(doggo.bark())

	doggo.name = "New name test"
}

type Dog struct {
	name string;
	bark func() string;
}

type Shape interface {
	area() float64
}

type Rectangle struct  {
	height float64
	width float64
}

func (rect *Rectangle) area() float64{
	return rect.height * rect.width
}

func getAreaOfShape (shape Shape) float64{
	return shape.area();
}