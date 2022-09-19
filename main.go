package main

import (
	"fmt"
)

const barkSound = "woof"
const dogName = "Lea"
const maxBellyRubs uint = 5
const numberOfBarks uint = 50

var remainingBarks uint = numberOfBarks - 1

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
