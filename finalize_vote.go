package main

import (
	"fmt"
	"sort"
)

func FinalizeVoting() {
	sort.Slice(Field, func(i, j int) bool {
		return Field[i].Number <= Field[j].Number
	})

	fmt.Println("Sie wiefolgt abgestimmt:")
	if Voted_list {
		fmt.Println("Sie haben ihr Listenkreuz bei", Voted_party, "gesetzt und diese Kandidat*innen von der Listenwahl ausgeschlossen:")

		for c := range Field {
			if Field[c].Crossed_out {
				fmt.Println(Field[c].Number, Field[c].Name)
			}
		}
	}

	fmt.Println("")
	fmt.Println("Sie haben diesen Kandidat*innen Personenstimmen gegeben")
	for c := range Field {
		if Field[c].Votes > 0 {
			fmt.Println(Field[c].Number, Field[c].Name, Field[c].Votes)
		}
	}

	fmt.Println("Drücken Sie [W], um ihre stimmen so auszudrucken und zu wählen.")
	fmt.Println("Drücken Sie [A], um den Vorgang abzubrechen.")
	var i string
	fmt.Scanln(&i)

	if i == "W" || i == "w" {
		Votes_left = 0
	}

}
