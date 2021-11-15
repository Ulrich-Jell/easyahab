package main

import (
	"sort"
)

func PersonVote(needle int) {
	sort.Slice(Field, func(i, j int) bool { // I don't really know how this works.
		return Field[i].Number <= Field[j].Number // It just does.
	}) //
	idx := sort.Search(len(Field), func(i int) bool { //
		return int(Field[i].Number) >= needle //
	}) //

	if Field[idx].Number == needle && Field[idx].CrossedOut {
		Field[idx].CrossedOut = false
		Field[idx].Votes = 1
		VotesLeft -= 1

	} else if Field[idx].Number == needle && Field[idx].Votes < 3 {
		Field[idx].Votes += 1
		VotesLeft -= 1

	} else if Field[idx].Number == needle && Field[idx].Votes == 3 {
		Field[idx].Votes = 0
		VotesLeft += 3

	}
	Rot(Field[idx].List)
	Braun()

}

func ReduceVote(needle int) {
	sort.Slice(Field, func(i, j int) bool {
		return Field[i].Number <= Field[j].Number
	})
	idx := sort.Search(len(Field), func(i int) bool {
		return int(Field[i].Number) >= needle
	})

	if Field[idx].Number == needle && Field[idx].Votes > 0 {
		Field[idx].Votes -= 1
		VotesLeft += 1
	}
	Rot(Field[idx].List)
	Braun()
}

func CrossCandidateOut(needle int) {
	sort.Slice(Field, func(i, j int) bool {
		return Field[i].Number <= Field[j].Number
	})
	idx := sort.Search(len(Field), func(i int) bool {
		return int(Field[i].Number) >= needle
	})

	if Field[idx].Number == needle && Field[idx].Votes > 0 {

		Field[idx].CrossedOut = true
		VotesLeft += Field[idx].Votes
		Field[idx].Votes = 0

	} else if Field[idx].Number == needle {
		Field[idx].CrossedOut = !Field[idx].CrossedOut
	}
	Rot(Field[idx].List)
	Braun()
}

func VoteList(s string) {
	sort.Slice(Lists, func(i, j int) bool {
		return Lists[i].Shorthand <= Lists[j].Shorthand
	})

	needle := s
	idx := sort.Search(len(Lists), func(i int) bool {
		return string(Lists[i].Shorthand) >= needle
	})

	if s == VotedParty {
		VotedParty = ""
		VotedList = false
		Lists[idx].ListVote = false

	} else if Lists[idx].Shorthand == needle && !VotedList {
		Lists[idx].ListVote = true
		VotedList = true
		VotedParty = Lists[idx].Shorthand

	} else if Lists[idx].Shorthand == needle && VotedList {
		Lists[idx].ListVote = true
		t := Lists[idx].Shorthand

		idx2 := sort.Search(len(Lists), func(i int) bool { //
			return string(Lists[i].Shorthand) >= VotedParty //to be honest, I don't recall why I put this in
		}) //
		Lists[idx2].ListVote = false
		VotedParty = t

	}
	Rot(Lists[idx].Number)
	Braun()
}

func Reset() {
	for i := range Field {
		Field[i].Votes = 0
		Field[i].CrossedOut = false
	}
	VotesLeft = AmountOfVotes
	Braun()
	Red.Objects = Red.Objects[:0]
}
