package main

//In a real case Scenenario, all Information regarding Candidates and Parties would be stored in a database.
//Here would be the code to transform the code into the structs
//At the moment, everything has to be put in manually

var VotesLeft int
var AmountOfVotes int
var VotedList bool
var VotedParty string

type Party struct {
	Number    int
	Name      string
	Shorthand string
	ListVote  bool
}

type Candidate struct {
	List       int
	Number     int
	Name       string
	Party      string
	Votes      int
	CrossedOut bool
}

func AddParty(no int, n string, s string) (result Party) {
	result = Party{no, n, s, false}
	return
}

func AddCandidate(l, int, no int, n string, p string) (result Candidate) {
	result = Candidate{l, no, n, p, 0, false}
	return
}

var Lists []Party
var Field []Candidate

func Connecttoserver() {
	AmountOfVotes = 16
	VotesLeft = 16

	Lists = []Party{
		{1, "A - Partei", "A", false},
		{2, "B - Partei", "B", false},
		{3, "Wählergruppe C", "C", false},
	}

	Field = []Candidate{
		{1, 101, "Kunze, Dieter", "A", 0, false},
		{1, 102, "Loisse, Claude", "A", 0, false},
		{1, 103, "Wölfel, Brigitte", "A", 0, false},
		{1, 104, "Hotel, Carla", "A", 0, false},
		{1, 105, "Müller, Ellen", "A", 0, false},
		{1, 106, "Aigari, Fabio", "A", 0, false},
		{1, 107, "Neu, Rolf", "A", 0, false},
		{1, 108, "Marold, Lorenz", "A", 0, false},
		{1, 109, "Zylka, Jennifer", "A", 0, false},
		{1, 110, "Baggio, Roberto", "A", 0, false},
		{1, 111, "Bender, Stephan", "A", 0, false},
		{1, 112, "Kappes, Günter", "A", 0, false},
		{1, 113, "Bongen, Hanno", "A", 0, false},
		{1, 114, "Seipelt, Helga", "A", 0, false},
		{1, 115, "Engel, Heidi", "A", 0, false},
		{1, 116, "Blome, Nikolaus", "A", 0, false},

		{2, 201, "Fischer, Kathrin", "B", 0, false},
		{2, 202, "Kunze, Karl", "B", 0, false},
		{2, 203, "Kalt, Angelika", "B", 0, false},
		{2, 204, "Schmitz, Paula", "B", 0, false},
		{2, 205, "Schulze, Martin", "B", 0, false},
		{2, 206, "Knops, Anton", "B", 0, false},
		{2, 207, "Lauer, Sofia", "B", 0, false},
		{2, 208, "Kautz, Paul", "B", 0, false},
		{2, 209, "Sprujit, Coby", "B", 0, false},
		{2, 210, "Pfeiffer, Gert", "B", 0, false},
		{2, 211, "Taschenbier, Bruno", "B", 0, false},
		{2, 212, "Webers, Ulrich", "B", 0, false},
		{2, 213, "Rose, Marlene", "B", 0, false},

		{3, 301, "Bremes, Peter", "C", 0, false},
		{3, 302, "Schachtner, Margarete", "C", 0, false},
		{3, 303, "Flach, Hubert", "C", 0, false},
		{3, 304, "Poensgen, Gerd", "C", 0, false},
	}
}
