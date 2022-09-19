package main

import (
	"fmt"
	"math"
	"time"
)

func greetUser() {
	fmt.Println("Hello World, I am", dogName)
	fmt.Printf("I have %v loud barks at my disposal, %s! ... and %v are left. \n", numberOfBarks, barkSound, remainingBarks)
}

func verifyRubs(bellyRubs int) uint {
	var result uint = uint(math.Abs(float64(bellyRubs)))

	if result > maxBellyRubs {
		fmt.Printf("This is way more than your neighbours can take! You can only rub the belly %v times at once. \n", maxBellyRubs)
		result = maxBellyRubs
	}

	if result > remainingBarks {
		fmt.Printf("I can't take that much belly rubbing now, you may rub %v times. \n", remainingBarks)
		result = remainingBarks
	}

	return result
}

func rubBelly(bellyRubs uint) {

	fmt.Println("Rubbing the belly...")

	showLoader(bellyRubs)

	fmt.Println("\rRubbing complete. Initiating barking.")
	time.Sleep(500 * time.Millisecond)
}

func showLoader(times uint) {
	var loadingSigns = []string{"/", "-", "\\", "|"}
	for i := 0; i <= int(times); i++ {
		for _, sign := range loadingSigns {
			fmt.Printf("\r %v", sign)
			time.Sleep(150 * time.Millisecond)
		}
	}
}

func bark(times uint) {
	fmt.Print("\n")

	for i := 0; i < int(times); i++ {
		fmt.Printf("%s! ", barkSound)
		time.Sleep(150 * time.Millisecond)
	}

	fmt.Print("\n")
	time.Sleep(150 * time.Millisecond)
}

func displayRemainingBarks() {
	fmt.Printf("\nThat was some premium rubbing! I have %v barks left. \n\n", remainingBarks)
}
