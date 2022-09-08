package main

import "fmt"

func main() {
	// var dogName = "Lea" // is equal to -->
	dogName := "Lea"             // create a variable and assign a value, no constants and no possibility to explicitly declare a type
	const numberOfBarks = 50     // total number of barks, type defined implicitly
	var remainingBarks uint = 50 // explicit specification mostly used when type is to be different than the default, uint is a non-negative number
	const barkSound = "woof"

	fmt.Println("Hello World, I am", dogName)
	fmt.Printf("I have %v loud barks at my disposal, %v, and %v are left. \n", numberOfBarks, barkSound, remainingBarks)

	var userName string
	var bellyRubs int
	fmt.Printf("dogName is a %T, remainingBarks is type %T and userName is %T. \n", dogName, remainingBarks, userName)

	// Scan take user input as a parameter
	// a pointer directs to a memory address where a variable is stored, concept from C, C++ (not in java f.ex.)
	fmt.Println(&userName)
	// direct Scan to the memory address
	fmt.Println("Enter your first name: ")
	fmt.Scan(&userName)

	//var userProvokeBarks int
	fmt.Println("How many belly rubs do you want to offer? You may offer between 0 and 5 rubs at a time.")
	fmt.Scan(&bellyRubs)
	fmt.Printf("Hey %v, thanks for offering %v belly rubs.", userName, bellyRubs)

	// userProvokeBarks = 2

	// need to define type explicitly when declaring without assignment
}
