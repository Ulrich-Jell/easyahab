package main

import "fmt"

//This is a dummy-file until I understand how to connect to MongoDB properly
//The goal is to import the MongoDB files als structs and let the voter manipulate these
//Until then the structs are coded here

// type Party struct {
// 	Name      string
// 	Shorthand string
// 	List_vote bool
// }

// type Candidate struct {
// 	Number      int
// 	Name        string
// 	Party       string
// 	Votes       int
// 	Crossed_out bool
// }

func AddParty(n string, s string) (result Party) {
	result = Party{n, s, false}
	return
}

func AddCandidate(no int, n string, p string) (result Candidate) {
	result = Candidate{no, n, p, 0, false}
	return
}

func Connecttoserver() {
	P1 := AddParty("English Songs", "ES")
	P2 := AddParty("Deutsche Lieder", "DL")
	P3 := AddParty("Deutscher Schlager", "S")

	C101 := AddCandidate(101, "Bohemian Rhaosody", "ES")
	C102 := AddCandidate(102, "Lucy in the Sky with Diamonds", "ES")
	C103 := AddCandidate(103, "November Rain", "ES")
	C104 := AddCandidate(104, "Hotel California", "ES")
	C105 := AddCandidate(105, "Thunderstruck", "ES")
	C106 := AddCandidate(106, "La Isla Bonita", "ES")
	C107 := AddCandidate(107, "Smells Like Teen Spirit", "ES")
	C108 := AddCandidate(108, "You can go your own Way", "ES")
	C109 := AddCandidate(109, "Sympathy for the Devil", "ES")
	C110 := AddCandidate(110, "Light my Fire", "ES")
	C111 := AddCandidate(111, "Nothing else Matters", "ES")
	C112 := AddCandidate(112, "Another Brick in the Wall (Part 2)", "ES")
	C113 := AddCandidate(113, "Thriller", "ES")
	C114 := AddCandidate(114, "Me and Bobby McGee", "ES")
	C115 := AddCandidate(115, "Heart if Glas", "ES")
	C116 := AddCandidate(116, "Don't Speak", "ES")

	C201 := AddCandidate(201, "An Tagen wie Diesen", "DL")
	C202 := AddCandidate(202, "Schrei nach Liebe", "DL")
	C203 := AddCandidate(203, "Mensch", "DL")
	C204 := AddCandidate(204, "Der Typ der bei der GEMA die Titel eintippt ist ein ganz blöder Penner", "DL")
	C205 := AddCandidate(205, "Abenteuerland", "DL")
	C206 := AddCandidate(206, "Küssen verboten", "DL")
	C207 := AddCandidate(207, "Nur ein Wort", "DL")
	C208 := AddCandidate(208, "Junimond", "DL")
	C209 := AddCandidate(209, "Ein Jahr (Es geht voran)", "DL")
	C210 := AddCandidate(210, "Das Model", "DL")
	C211 := AddCandidate(211, "Keine Macht für Niemnden", "DL")
	C212 := AddCandidate(212, "Haus am See", "DL")

	C301 := AddCandidate(301, "Atemlos durch die Nacht", "S")
	C302 := AddCandidate(302, "Herzilein, du musst nicht traurig sein", "S")
	C303 := AddCandidate(303, "Verlieben, Verloren, Vergessen, Verzeihn", "S")
	C304 := AddCandidate(304, "Lebt denn der alte Holzmichl noch", "S")

	fmt.Println("Diese Parteien treten zur Wahl an:")
	fmt.Print("1.", P1.Name, " (", P1.Shorthand, ")\n")
	fmt.Print("2.", P2.Name, " (", P2.Shorthand, ")\n")
	fmt.Print("3.", P3.Name, " (", P3.Shorthand, ")\n")
	fmt.Println("")

	fmt.Println("Sie können diese Personen wählen:")
	fmt.Println(C101.Number, C101.Name, C101.Party)
	fmt.Println(C102.Number, C102.Name, C102.Party)
	fmt.Println(C103.Number, C103.Name, C103.Party)
	fmt.Println(C104.Number, C104.Name, C104.Party)
	fmt.Println(C105.Number, C105.Name, C105.Party)
	fmt.Println(C106.Number, C106.Name, C106.Party)
	fmt.Println(C107.Number, C107.Name, C107.Party)
	fmt.Println(C108.Number, C108.Name, C108.Party)
	fmt.Println(C109.Number, C109.Name, C109.Party)
	fmt.Println(C110.Number, C110.Name, C110.Party)
	fmt.Println(C111.Number, C111.Name, C111.Party)
	fmt.Println(C112.Number, C112.Name, C112.Party)
	fmt.Println(C113.Number, C113.Name, C113.Party)
	fmt.Println(C114.Number, C114.Name, C114.Party)
	fmt.Println(C115.Number, C115.Name, C115.Party)
	fmt.Println(C116.Number, C116.Name, C116.Party)
	fmt.Println("")

	fmt.Println(C201.Number, C201.Name, C201.Party)
	fmt.Println(C202.Number, C202.Name, C202.Party)
	fmt.Println(C203.Number, C203.Name, C203.Party)
	fmt.Println(C204.Number, C204.Name, C204.Party)
	fmt.Println(C205.Number, C205.Name, C205.Party)
	fmt.Println(C206.Number, C206.Name, C206.Party)
	fmt.Println(C207.Number, C207.Name, C207.Party)
	fmt.Println(C208.Number, C208.Name, C208.Party)
	fmt.Println(C209.Number, C209.Name, C209.Party)
	fmt.Println(C210.Number, C210.Name, C210.Party)
	fmt.Println(C211.Number, C211.Name, C211.Party)
	fmt.Println(C212.Number, C212.Name, C212.Party)
	fmt.Println("")

	fmt.Println(C301.Number, C301.Name, C301.Party)
	fmt.Println(C302.Number, C302.Name, C302.Party)
	fmt.Println(C303.Number, C304.Name, C303.Party)
	fmt.Println(C304.Number, C304.Name, C303.Party)

}
