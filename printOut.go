package main

import (
	"fmt"
	"log"
	"os"
)

func PrintOut() {
	f, err := os.Create("ballot.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	fmt.Fprintln(f, "Wahlzettel für die Kommunalwahl der Stadt Lummerland vom 31.04.2002")
	fmt.Fprintln(f, "Wahllokal 173", "Antoniussheim")
	fmt.Fprintln(f, "_______________________________________________________________")
	fmt.Fprintln(f, "Per listenkreuz gewählte Partei oder Wählergruppe:", Voted_party)
	fmt.Fprintln(f, "")
	fmt.Fprintln(f, "Von dieser Liste ausgenommene Kandidat*innen:")

	for c := range Field {
		if Field[c].Crossed_out {
			fmt.Fprintln(f, Field[c].Number, Field[c].Name)
		}
	}
	fmt.Fprintln(f, "")

	fmt.Fprintln(f, "Durch Personenwahl gewählte Kandidat*innen:")
	for c := range Field {
		if Field[c].Votes > 0 {
			fmt.Fprintln(f, Field[c].Number, Field[c].Name, Field[c].Votes)
		}
	}

}
