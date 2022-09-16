package main

import (
	"fmt"
	"time"
)

const barkSound = "woof"
const dogName = "Lea"
const maxBellyRubs uint = 5
const numberOfBarks uint = 50

var remainingBarks uint = 49

func main() {

	greetUser()
	userName := gatherUserName()

	for {
		userBellyRubs := askForRubs()
		rubs := verifyRubs(userBellyRubs)
		rubBelly(rubs)
		bark(rubs)

		remainingBarks -= rubs

		if remainingBarks == 0 {
			fmt.Printf("Sorry %s, I cannot bark any more, I sleep now. Bye.", userName)
			break
		}

		displayRemainingBarks()
	}

}

func greetUser() {
	fmt.Println("Hello World, I am", dogName)
	fmt.Printf("I have %v loud barks at my disposal, %s! ... and %v are left. \n", numberOfBarks, barkSound, remainingBarks)
}

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

func verifyRubs(bellyRubs uint) uint {
	if bellyRubs > maxBellyRubs {
		fmt.Printf("This is way more than your neighbours can take! You can only rub the belly %v times at once. \n", maxBellyRubs)
		bellyRubs = maxBellyRubs
	}

	if bellyRubs > remainingBarks {
		fmt.Printf("I can't take that much belly rubbing now, you may rub %v times. \n", remainingBarks)
		bellyRubs = remainingBarks
	}

	return bellyRubs
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
