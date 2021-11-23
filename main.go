package main

//the main problem with the layout is, that I could only use fyne's grid layout to get a usable result
//in fyne all collumns in a grid must have the same width
//fyne has a possibility to create costum layouts, but I could not figure out how

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var CandidateContainer *fyne.Container
var PartyContainer *fyne.Container
var HeadingContainer *fyne.Container
var Scroll *fyne.Container

var VotesLabel binding.Int
var App fyne.App

//All Colors refer to specific Elements of the Layout
//functions, that are named after colors create or refresh these elements in the GUI
//see Layout.png for a visual representation

func DrawCandidates(l int) {
	CandidateContainer.Objects = CandidateContainer.Objects[:0]

	CandidateContainer.Add(widget.NewLabel("Listenplatz"))
	CandidateContainer.Add(widget.NewLabel("Name"))
	CandidateContainer.Add(widget.NewLabel("Stimmen"))
	CandidateContainer.Add(widget.NewLabel("Simme geben"))
	CandidateContainer.Add(widget.NewLabel("Stimmen reduzieren"))
	CandidateContainer.Add(widget.NewLabel("Von der Listenwahl ausgeschlossen"))

	for i := range Field {
		if Field[i].List == l {
			n := Field[i].Number
			ns := strconv.Itoa(n)
			v := Field[i].Votes
			vs := strconv.Itoa(v)
			t := Field[i].Name

			CandidateContainer.Add(widget.NewLabel(ns))
			CandidateContainer.Add(widget.NewLabel(t))
			CandidateContainer.Add(widget.NewLabel(vs))
			CandidateContainer.Add(widget.NewButton("+", func() {
				PersonVote(n)
			}))
			CandidateContainer.Add(widget.NewButton("-", func() {
				ReduceVote(n)
			}))
			if Field[i].CrossedOut {
				CandidateContainer.Add(widget.NewButton("Ja", func() {
					CrossCandidateOut(n)
				}))
			} else {
				CandidateContainer.Add(widget.NewButton("Nein", func() {
					CrossCandidateOut(n)
				}))
			}
		}

	}

	if Lists[l-1].ListVote {
		CandidateContainer.Add(widget.NewButton("Liste wurde gewählt", func() {
			VoteList(Lists[l-1].Shorthand)
		}))
	} else {
		CandidateContainer.Add(widget.NewButton("Liste wählen", func() {
			fmt.Println(Lists[l-1].Shorthand)
			VoteList(Lists[l-1].Shorthand)
		}))
	}

}

func DrawHeading() {
	HeadingContainer.Objects = HeadingContainer.Objects[:0]

	HeadingContainer.Add(widget.NewLabel("Sie haben noch"))
	HeadingContainer.Add(widget.NewLabel(strconv.Itoa(VotesLeft)))
	HeadingContainer.Add(widget.NewLabel("Stimmen."))
	HeadingContainer.Add(widget.NewLabel("Listenwahl:"))
	HeadingContainer.Add(widget.NewLabel(VotedParty))
}

func drawPartyContainer() {
	for i := range Lists {
		n := Lists[i].Number
		PartyContainer.Add(widget.NewButton(Lists[i].Shorthand, func() {
			DrawCandidates(n)
		}))

	}
}

func help(a fyne.App) {
	win := a.NewWindow("Hilfe")

	h := widget.NewLabelWithStyle("orem ipsum dolor sit amet, consectetuer adipiscing elit. Aenean commodo ligula eget dolor. Aenean massa. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Donec quam felis, ultricies nec, pellentesque eu, pretium quis, sem. Nulla consequat massa quis enim. Donec pede justo, fringilla vel, aliquet nec, vulputate eget, arcu. In enim justo, rhoncus ut, imperdiet a, venenatis vitae, justo. Nullam dictum felis eu pede mollis pretium. Integer tincidunt. Cras dapibus. Vivamus elementum semper nisi. Aenean vulputate eleifend tellus. Aenean leo ligula, porttitor eu, consequat vitae, eleifend ac, enim. Aliquam lorem ante, dapibus in, viverra quis, feugiat a, tellus. Phasellus viverra nulla ut metus varius laoreet. Quisque rutrum. Aenean imperdiet. Etiam ultricies nisi vel augue. Curabitur ullamcorper ultricies nisi. Nam eget dui.", fyne.TextAlignLeading, fyne.TextStyle{})
	h.Wrapping = 3
	q := widget.NewButton("Schließen", func() {
		win.Close()
	})

	win.SetContent(container.NewVBox(h, q))
	win.Show()

}

func main() {
	App = app.New()
	w := App.NewWindow("Kommunalwahl 2025")

	Connecttoserver()

	CandidateContainer = container.NewGridWithColumns(6)

	PartyContainer = container.NewGridWithColumns(1)

	headingContainer1 := widget.NewLabelWithStyle("Herzlich Willkommen zur Kommunalwahl 2025 in Lummerland.", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	headingContainer2 := widget.NewLabelWithStyle("Dieses Programm erstellt Ihren Stimmzettel und druckt ihn aus. Sie können beliebig viele Stimmzettel ausdrucken, jedoch nur einen in die Urne werfen. Es werden keine Daten gespeichert und es ist kein Rückschluss auf Ihre Person möglich!", fyne.TextAlignLeading, fyne.TextStyle{})
	headingContainer2.Wrapping = 3
	HeadingContainer = container.NewGridWithColumns(3)
	HeadingContainer := container.NewVBox(headingContainer1, headingContainer2, HeadingContainer)

	bottomContainer1 := widget.NewButton("Wahl Abschließen", func() {
		PrintOut()
	})
	bottomContainer2 := widget.NewButton("Von vorne Beginnen", func() {
		Reset()
	})
	bottomContainer3 := widget.NewButton("Hilfe", func() {
		help(App)
	})
	bottomContainer := container.NewHBox(bottomContainer1, bottomContainer2, bottomContainer3)

	drawPartyContainer()
	DrawHeading()

	Scroll := container.NewScroll(CandidateContainer) //adds a container that scrolls everything
	all := container.New(layout.NewBorderLayout(HeadingContainer, bottomContainer, PartyContainer, nil), HeadingContainer, PartyContainer, bottomContainer, Scroll)

	w.SetContent(all)

	w.Resize(fyne.NewSize(1366, 768))

	w.ShowAndRun()
}
