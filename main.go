package main

import (
	"fmt"
	"strings"
	"time"
)

const barkSound = "woof"

func main() {
	// var dogName = "Lea" // is equal to -->
	dogName := "Lea" // create a variable and assign a value, no constants and no possibility to explicitly declare a type
	// used for short lived variables, result of a function call or an iterator in a function

	const numberOfBarks = 50     // total number of barks, type defined implicitly
	var remainingBarks uint = 50 // explicit specification mostly used when type is to be different than the default, uint is a non-negative number

	fmt.Println("Hello World, I am", dogName)
	fmt.Printf("I have %v loud barks at my disposal, %s, and %v are left. \n", numberOfBarks, barkSound, remainingBarks)

	var userName string
	// is the same as var userName string = ""
	var bellyRubs uint
	var maxBellyRubs uint = 5
	// fmt.Printf("dogName is a %T, remainingBarks is type %T and userName is %T. \n", dogName, remainingBarks, userName)

	// Scan take user input as a parameter
	// a pointer directs to a memory address where a variable is stored, concept from C, C++ (not in java f.ex.)
	// fmt.Println(&userName)

	// direct Scan to the memory address
	fmt.Println("Enter your first name: ")
	fmt.Scan(&userName)

	fmt.Printf("Hi %v! Nice to meet you. \n", userName)
	fmt.Printf("How many belly rubs would you like to offer? You may offer between 0 and %v rubs at a time. \n", maxBellyRubs)
	fmt.Scan(&bellyRubs)
	fmt.Printf("You've offered %v belly rubs.\n", bellyRubs)

	if bellyRubs > 5 {
		fmt.Printf("This is way more than your neighbours can take! You can only rub the belly %v times at once. \n", maxBellyRubs)
		bellyRubs = maxBellyRubs
		remainingBarks -= maxBellyRubs
	}

	if bellyRubs > remainingBarks {
		fmt.Println("I can't take that much belly rubbing now, I need to sleep first.")
		bellyRubs = remainingBarks
		remainingBarks = 0
	}

	fmt.Println("Rubbing the belly...")
	var loadingRubbingSigns = []string{"/", "-", "\\", "|"}
	for i := 0; i <= int(bellyRubs); i++ {
		for _, sign := range loadingRubbingSigns {
			fmt.Printf("\r %v", sign)
			time.Sleep(150 * time.Millisecond)
		}
	}

	fmt.Println("\rRubbing complete. Initiating barking. ")
	time.Sleep(500 * time.Millisecond)

	var barks []string

	for i := 0; i < int(bellyRubs); i++ {
		barks = append(barks, barkSound)
	}

	barkingString := strings.Join(barks, ", ")
	fmt.Printf("\n%s! \n\n", barkingString)

}
