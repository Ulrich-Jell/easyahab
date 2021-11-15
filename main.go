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

var Red *fyne.Container
var Blue *fyne.Container
var Brown0 *fyne.Container

var VotesLabel binding.Int
var App fyne.App

//All Colors refer to specific Elements of the Layout
//functions, that are named after colors create or refresh these elements in the GUI
//see Layout.png for a visual representation

func Rot(l int) {
	Red.Objects = Red.Objects[:0]

	Red.Add(widget.NewLabel("Listenplatz"))
	Red.Add(widget.NewLabel("Name"))
	Red.Add(widget.NewLabel("Stimmen"))
	Red.Add(widget.NewLabel("Simme geben"))
	Red.Add(widget.NewLabel("Stimmen reduzieren"))
	Red.Add(widget.NewLabel("Von der Listenwahl ausgeschlossen"))

	for i := range Field {
		if Field[i].List == l {
			n := Field[i].Number
			ns := strconv.Itoa(n)
			v := Field[i].Votes
			vs := strconv.Itoa(v)
			t := Field[i].Name

			Red.Add(widget.NewLabel(ns))
			Red.Add(widget.NewLabel(t))
			Red.Add(widget.NewLabel(vs))
			Red.Add(widget.NewButton("+", func() {
				PersonVote(n)
			}))
			Red.Add(widget.NewButton("-", func() {
				ReduceVote(n)
			}))
			if Field[i].CrossedOut {
				Red.Add(widget.NewButton("Ja", func() {
					CrossCandidateOut(n)
				}))
			} else {
				Red.Add(widget.NewButton("Nein", func() {
					CrossCandidateOut(n)
				}))
			}
		}

	}

	if Lists[l-1].ListVote {
		Red.Add(widget.NewButton("Liste wurde gewählt", func() {
			VoteList(Lists[l-1].Shorthand)
		}))
	} else {
		Red.Add(widget.NewButton("Liste wählen", func() {
			fmt.Println(Lists[l-1].Shorthand)
			VoteList(Lists[l-1].Shorthand)
		}))
	}

}

func Braun() {
	Brown0.Objects = Brown0.Objects[:0]

	Brown0.Add(widget.NewLabel("Sie haben noch"))
	Brown0.Add(widget.NewLabel(strconv.Itoa(VotesLeft)))
	Brown0.Add(widget.NewLabel("Stimmen."))
	Brown0.Add(widget.NewLabel("Listenwahl:"))
	Brown0.Add(widget.NewLabel(VotedParty))
}

func blue() {
	for i := range Lists {
		n := Lists[i].Number
		Blue.Add(widget.NewButton(Lists[i].Shorthand, func() {
			Rot(n)
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

	Red = container.NewGridWithColumns(6)

	Blue = container.NewGridWithColumns(1)

	brown1 := widget.NewLabelWithStyle("Herzlich Willkommen zur Kommunalwahl 2025 der Stadt Fulda.", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	brown2 := widget.NewLabelWithStyle("Dieses Programm erstellt Ihren Stimmzettel und druckt ihn aus. Sie können beliebig viele Stimmzettel ausdrucken, jedoch nur einen in die Urne werfen. Es werden keine Daten gespeichert und es ist kein Rückschluss auf Ihre Person möglich!", fyne.TextAlignLeading, fyne.TextStyle{})
	brown2.Wrapping = 3
	Brown0 = container.NewGridWithColumns(3)
	brown := container.NewVBox(brown1, brown2, Brown0)

	pink1 := widget.NewButton("Wahl Abschließen", func() {
		PrintOut()
	})
	pink2 := widget.NewButton("Von vorne Beginnen", func() {
		Reset()
	})
	pink3 := widget.NewButton("Hilfe", func() {
		help(App)
	})
	pink := container.NewHBox(pink1, pink2, pink3)

	blue()
	Braun()

	all := container.New(layout.NewBorderLayout(brown, pink, Blue, Red), brown, Blue, Red, pink)
	scroll := container.NewScroll(all) //adds a container that scrolls everything

	w.SetContent(scroll)

	w.Resize(fyne.NewSize(1366, 768))

	w.ShowAndRun()
}
