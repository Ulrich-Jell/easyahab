package main

//This is a dummy-file until I understand how to connect to MongoDB properly
//The goal is to import the MongoDB files als structs and let the voter manipulate these
//Until then the structs are coded here

type Party struct {
	Name      string
	Shorthand string
	List_vote bool
}

type Candidate struct {
	Number      int
	Name        string
	Party       string
	Votes       int
	Crossed_out bool
}

func AddParty(n string, s string) (result Party) {
	result = Party{n, s, false}
	return
}

func AddCandidate(no int, n string, p string) (result Candidate) {
	result = Candidate{no, n, p, 0, false}
	return
}

var Lists []Party
var Field []Candidate

func Connecttoserver() {
	Lists = []Party{
		{"English Songs", "ES", false},
		{"Deutsche Lieder", "DL", false},
		{"Deutscher Schlager", "S", false},
	}

	Field = []Candidate{
		{101, "Bohemian Rhaosody", "ES", 0, false},
		{102, "Lucy in the Sky with Diamonds", "ES", 0, false},
		{103, "November Rain", "ES", 0, false},
		{104, "Hotel California", "ES", 0, false},
		{105, "Thunderstruck", "ES", 0, false},
		{106, "La Isla Bonita", "ES", 0, false},
		{107, "Smells Like Teen Spirit", "ES", 0, false},
		{108, "You can go your own Way", "ES", 0, false},
		{109, "Sympathy for the Devil", "ES", 0, false},
		{110, "Light my Fire", "ES", 0, false},
		{111, "Nothing else Matters", "ES", 0, false},
		{112, "Another Brick in the Wall (Part 2)", "ES", 0, false},
		{113, "Thriller", "ES", 0, false},
		{114, "Me and Bobby McGee", "ES", 0, false},
		{115, "Heart if Glas", "ES", 0, false},
		{116, "Don't Speak", "ES", 0, false},

		{201, "Junimond von Rio Reiser", "DL", 0, false},
		{202, "Einer mehr (Es geht voran) von Fehlfarben", "DL", 0, false},
		{203, "Das Model von Kraftwerk", "DL", 0, false},
		{204, "Weißes Papier von Element of Crime", "DL", 0, false},
		{205, "Der Mussolini von DAF", "DL", 0, false},
		{206, "Blaue Augen von Ideal", "DL", 0, false},
		{207, "Der Traum ist aus von Ton Steine Scherben", "DL", 0, false},
		{208, "Paul ist tot von Fehlfarben", "DL", 0, false},
		{209, "Keine Macht für niemand von Ton Steine Scherben", "DL", 0, false},
		{210, "Macht kaputt was euch kaputt macht", "DL", 0, false},
		{211, "Gottseidank nicht England von Fehlfarben", "DL", 0, false},
		{212, "Schwarz zu blau von Peter Fox", "DL", 0, false},
		{213, "Wenn ich mir was wünschen dürfte von Marlene Dietrich", "DL", 0, false},

		{301, "Atemlos durch die Nacht", "S", 0, false},
		{302, "Herzilein, du musst nicht traurig sein", "S", 0, false},
		{303, "Verlieben, Verloren, Vergessen, Verzeihn", "S", 0, false},
		{304, "Lebt denn der alte Holzmichl noch", "S", 0, false},
	}
}
