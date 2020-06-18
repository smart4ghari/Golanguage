package main

import "fmt"

// while declaring a variable outside the function you should use full format
var str string = "Hi there!"

func main() {
	//There are many ways to represent a variable
	var number = 9 //Method 1
	i := 23        //Method 2 most simple form
	var j int = 4  //Method 3
	fmt.Println(number, "This is the first method")
	fmt.Println(i, "This is the second method")
	fmt.Println(j, "This is the third method")
	fmt.Println(str)
}
