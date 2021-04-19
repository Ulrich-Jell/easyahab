package main

import "fmt"

func main() {
	Connecttoserver()

	for votes_left := 16; votes_left > 0; {
		println("Sie haben noch", votes_left, "Stimmen.")
		fmt.Println("Drücken Sie [P], um Personenstimmen zu vergeben.")
		fmt.Println("Drücken Sie [S], um ein*e Kandidat*in von der Listenwahl auszuschließen")
		fmt.Println("Drücken Sie [L], um ein Listenkreuz zu setzen.")
		fmt.Println("Drücken Sie [V], um die Abstimmung zu abzuschließen.")
		var i2 string
		fmt.Scanln(&i2)

		//vote for a person
		if i2 == "P" || i2 == "p" {
			fmt.Println("### Person Voting not implemented yet###")

			//cross candidate out
		} else if i2 == "S" || i2 == "s" {
			fmt.Println("### Crossing Candidates out implemented yet###")

			// List Voting
		} else if i2 == "L" || i2 == "l" {
			fmt.Println("### List voting not implemented yet###")

			//Finish Voting
		} else if i2 == "A" || i2 == "a" {
			fmt.Println("### Finishing not implemented yet")
		}

	}
}
