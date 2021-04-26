package main

import (
	"fmt"
	"sort"
)

func PersonVote() {
	fmt.Println("Wen möchten Sie Wählen? Bitte geben Sie den Listenplatz an.")
	var cast int
	fmt.Scanln(&cast)

	sort.Slice(Field, func(i, j int) bool {
		return Field[i].Number <= Field[j].Number
	})

	needle := cast
	idx := sort.Search(len(Field), func(i int) bool {
		return int(Field[i].Number) >= needle
	})

	if Field[idx].Number == needle && Field[idx].Crossed_out {
		fmt.Println("Sie haben", Field[idx].Name, "bereits von der Listenwahl ausgeschlossen. Was möchten Sie tun?")
		fmt.Println("Drücken Sie [F], um", Field[idx].Name, "eine Stimme zu geben und wieder bei der Listenwahl zu berücksichtigen.")
		fmt.Println("Drücken Sie [A], um den Vorgang abzubrechen.")
		var vi2 string
		fmt.Scanln(&vi2)

		if vi2 == "F" || vi2 == "f" {
			Field[idx].Crossed_out = false
			Field[idx].Votes += 1
			Votes_left -= 1
		}

	} else if Field[idx].Number == needle && Field[idx].Votes < 3 {
		Field[idx].Votes += 1
		fmt.Println("Sie haben", Field[idx].Name, "eine Stimme gegeben.")
		fmt.Println(Field[idx].Name, "hat jetzt", Field[idx].Votes, "Stimmen")
		Votes_left -= 1

	} else if Field[idx].Number == needle && Field[idx].Votes == 3 {
		fmt.Println(Field[idx].Name, "hat bereits drei Stimmen. Was möchten Sie tun?")
		fmt.Println("Drücken Sie [0], um die Stimmen von", Field[idx].Votes, "auf 0 zu setzen.")
		fmt.Println("Drücken Sie [A] um den Vorgang abzubrechen.")
		var vi1 string
		fmt.Scanln(&vi1)

		if vi1 == "0" {
			Field[idx].Votes = 0
			Votes_left += 3
		}

	} else {
		fmt.Println("Found noting: ", idx)
	}

}

func CrossCandidateOut() {
	fmt.Println("Wen möchten Sie von der Listenwahl ausschließen? Bitte geben Sie den Listenplatz an.")
	var cast int
	fmt.Scanln(&cast)

	sort.Slice(Field, func(i, j int) bool {
		return Field[i].Number <= Field[j].Number
	})

	needle := cast
	idx := sort.Search(len(Field), func(i int) bool {
		return int(Field[i].Number) >= needle
	})
	if Field[idx].Number == needle && Field[idx].Votes > 0 {
		fmt.Println("Sie haben", Field[idx].Name, "bereits", Field[idx].Votes, "gegeben. Was möchten Sie tun?")
		fmt.Println("Drücken Sie [F], um fortzufahren und", Field[idx].Name, "von der Listenwahl auszuschließen.")
		fmt.Println("Sie können anschließend", Field[idx].Votes, "Stimmen neu vergeben.")
		fmt.Println("Drücken Sie [A], um den Vorgang abzubrechen.")
		var cast2 string
		fmt.Scanln(&cast2)

		if cast2 == "F" || cast2 == "f" {
			Field[idx].Crossed_out = true
			Votes_left += Field[idx].Votes
		}
	} else if Field[idx].Number == needle {
		Field[idx].Crossed_out = !Field[idx].Crossed_out
		fmt.Println(Field[idx])
	}
}

func ListVote() {
	fmt.Println("Welche Partei oder Wähler*innengruppe möchten Sie wählen? Bitte geben Sie das Kürzel an.")
	var cast string
	fmt.Scanln(&cast)

	sort.Slice(Lists, func(i, j int) bool {
		return Lists[i].Shorthand <= Lists[j].Shorthand
	})

	needle := cast
	idx := sort.Search(len(Lists), func(i int) bool {
		return string(Lists[i].Shorthand) >= needle
	})

	if Lists[idx].Shorthand == needle && !Voted_list {
		Lists[idx].List_vote = true
		Voted_list = true
		Voted_party = Lists[idx].Shorthand
		fmt.Println("Sie haben ein Listenkreuz bei", Lists[idx].Name, "gesetzt.")

	} else if Lists[idx].Shorthand == needle && Voted_list {
		fmt.Println("Sie haben bereits bei", Voted_party, "ein Listenkreuz gesetzt. Was möchten Sie tun?")
		fmt.Println("Drücken Sie [F], um das Kreuz bei", Voted_party, "zu löschen und statt dessen bei", Lists[idx].Shorthand, "zu setzen.")
		fmt.Println("Drücken Sie [A], um den Vorgang abzubrechen.")
		var lv string
		fmt.Scanln(&lv)

		if lv == "F" || lv == "f" {

			Lists[idx].List_vote = true
			t := Lists[idx].Shorthand

			idx2 := sort.Search(len(Lists), func(i int) bool {
				return string(Lists[i].Shorthand) >= Voted_party
			})
			Lists[idx2].List_vote = false
			Voted_party = t
		}
	}

}
