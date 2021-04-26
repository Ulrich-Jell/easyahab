package main

import (
	"fmt"
)

var Votes_left int
var Voted_list bool
var Voted_party string

func main() {
	Connecttoserver()

	Voted_list = false
	Voted_party = ""

	for Votes_left = 16; Votes_left > 0; {
		println("Sie haben noch", Votes_left, "Stimmen.")
		fmt.Println("Drücken Sie [P], um Personenstimmen zu vergeben.")
		fmt.Println("Drücken Sie [S], um ein*e Kandidat*in von der Listenwahl auszuschließen")
		fmt.Println("Drücken Sie [L], um ein Listenkreuz zu setzen.")
		fmt.Println("Drücken Sie [A], um die Abstimmung zu abzuschließen.")
		var i2 string
		fmt.Scanln(&i2)

		//vote for a person
		if i2 == "P" || i2 == "p" {
			PersonVote()

			//cross candidate out
		} else if i2 == "S" || i2 == "s" {
			CrossCandidateOut()

			// List Voting
		} else if i2 == "L" || i2 == "l" {
			ListVote()

			//Finish Voting
		} else if i2 == "A" || i2 == "a" {
			fmt.Println("### Finishing not implemented yet")
			FinalizeVoting()

		}

	}
	PrintOut()
	fmt.Println("Vielen Dank, dass Sie den easyahab Wahlassistent zur vereinfachung der Wahl genutzt haben!")
	fmt.Println("Wir hoffen, dass das Wahlergebnis Ihren Erwartungen entspricht und wünschen Ihnen einen schönen Tag!")

}
