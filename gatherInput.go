package main

import (
	"fmt"
)

func gatherUserName() string {
	var userName string

	fmt.Println("What's your first name?")
	fmt.Scan(&userName)

	fmt.Printf("Hi %v! Nice to meet you. \n", userName)

	return userName
}

func askForRubs() uint {
	var bellyRubs uint

	fmt.Printf("How many belly rubs would you like to offer? You may offer between 0 and %v rubs at a time. \n", maxBellyRubs)
	fmt.Scan(&bellyRubs)
	fmt.Printf("You've offered %v belly rubs.\n", bellyRubs)

	return bellyRubs
}
