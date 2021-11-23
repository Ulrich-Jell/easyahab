package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var Ballot *os.Process

func random(min int, max int) int {
	return rand.Intn(max-min) + min
}

func PrintOut() {
	f, err := os.Create("ballot.txt") //normally I would write the ballot to a pdf file, but I wasn't able to figure that out (,yet).

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	fmt.Fprintln(f, "Stimmzettel für die Kommunalwahl der Stadt Lummerland vom 31.04.2025")
	fmt.Fprintln(f, "Wahllokal 173", "Antoniusheim")
	fmt.Fprintln(f, "_______________________________________________________________")
	fmt.Fprintln(f, "Per listenkreuz gewählte Partei oder Wählergruppe:", VotedParty)
	fmt.Fprintln(f, "")
	fmt.Fprintln(f, "Von dieser Liste ausgenommene Kandidat*innen:")

	for c := range Field {
		if Field[c].CrossedOut {
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

	fmt.Fprintln(f, "")
	fmt.Fprintln(f, "")
	fmt.Fprintln(f, "")

	fmt.Fprintln(f, "Bitte schreiben Sie die folgende Zahl ab, um sicher zu stellen, dass ein Mensch diesen Stimmzettel erstellt hat.")

	rand.Seed(time.Now().UnixNano())
	randomNum := random(1000, 9999)
	fmt.Fprintln(f, randomNum)
	fmt.Fprintln(f, "")
	fmt.Fprintln(f, "")
	fmt.Fprintln(f, "")

	fmt.Fprintln(f, "_____  _____  _____  _____")

	showAnother(App)

	Ballot := exec.Command("kate", "ballot.txt") //how would I adress the default text editor?
	Ballot.Run()

	//Ballot, _ := os.StartProcess("/usr/bin/kate", []string{"kate", "ballot.txt"}, new(os.ProcAttr))
	//fmt.Println(Ballot)
	//time.Sleep(time.Second * 10)
	// Ballot.Kill()

	//I tried to automatically close the text editor, but it only works in this scope.
	//I don't know why

}

func showAnother(a fyne.App) {

	win := a.NewWindow("Abschluss")

	q := widget.NewButton("Schließen", func() {
		//Ballot.Kill()
		win.Close()
	})
	r := widget.NewButton("Wahl beenden", func() {
		Reset()
		//Ballot.Kill()
		win.Close()
	})

	win.SetContent(container.NewHBox(q, r))

	win.Show()

}
