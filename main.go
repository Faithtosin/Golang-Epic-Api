package main

import (
	"fmt"
)

const (
	A int = 1
	B     = 3.14
	C     = "Hi!"
)

// this is a comment
func main() {
	//var student1 string = "John" //type is string
	//var student2 = "Jane"        //type is inferred
	//x := 2                       //type is inferred
	fmt.Println("Hello World!")

	var a string
	const t int = 500
	a = "3"
	c, d := 7, "World!"
	fmt.Printf("%v is of type %T \n", d, a)
	fmt.Printf("%b", c)

	var arr1 = [3]int{1, 2, 3}
	arr2 := [5]int{4, 5, 6, 7, 8}

	fmt.Println(len(arr1))
	fmt.Println(arr2)
}
