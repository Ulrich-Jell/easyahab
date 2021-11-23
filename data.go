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
		{1, "Demokratische Partei Deutschland", "DPD", false},
		{2, "Deutsche Humanistische Liberale", "DHL", false},
		{3, "Gemeinsam Lummerland stärken", "GLS", false},
	}

	Field = []Candidate{
		{1, 101, "Kunze, Dieter", "DPD", 0, false},
		{1, 102, "Loisse, Claude", "DPD", 0, false},
		{1, 103, "Wölfel, Brigitte", "DPD", 0, false},
		{1, 104, "Hotel, Carla", "DPD", 0, false},
		{1, 105, "Müller, Ellen", "DPD", 0, false},
		{1, 106, "Aigari, Fabio", "DPD", 0, false},
		{1, 107, "Neu, Rolf", "DPD", 0, false},
		{1, 108, "Marold, Lorenz", "DPD", 0, false},
		{1, 109, "Zylka, Jennifer", "DPD", 0, false},
		{1, 110, "Baggio, Roberto", "DPD", 0, false},
		{1, 111, "Bender, Stephan", "DPD", 0, false},
		{1, 112, "Kappes, Günter", "DPD", 0, false},
		{1, 113, "Bongen, Hanno", "DPD", 0, false},
		{1, 114, "Seipelt, Helga", "DPD", 0, false},
		{1, 115, "Engel, Heidi", "DPD", 0, false},
		{1, 116, "Blome, Nikolaus", "DPD", 0, false},

		{2, 201, "Fischer, Kathrin", "DHL", 0, false},
		{2, 202, "Kunze, Karl", "DHL", 0, false},
		{2, 203, "Kalt, Angelika", "DHL", 0, false},
		{2, 204, "Schmitz, Paula", "DHL", 0, false},
		{2, 205, "Schulze, Martin", "DHL", 0, false},
		{2, 206, "Knops, Anton", "DHL", 0, false},
		{2, 207, "Lauer, Sofia", "DHL", 0, false},
		{2, 208, "Kautz, Paul", "DHL", 0, false},
		{2, 209, "Sprujit, Coby", "DHL", 0, false},
		{2, 210, "Pfeiffer, Gert", "DHL", 0, false},
		{2, 211, "Taschenbier, Bruno", "DHL", 0, false},
		{2, 212, "Webers, Ulrich", "DHL", 0, false},
		{2, 213, "Rose, Marlene", "DHL", 0, false},

		{3, 301, "Bremes, Peter", "GLS", 0, false},
		{3, 302, "Schachtner, Margarete", "GLS", 0, false},
		{3, 303, "Flach, Hubert", "GLS", 0, false},
		{3, 304, "Poensgen, Gerd", "GLS", 0, false},
	}
}
