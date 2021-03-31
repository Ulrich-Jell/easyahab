package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

func aufraeumen() {
	reset_v, err := CandidatesCollection.UpdateMany(
		context.TODO(),
		bson.M{},
		bson.D{
			{"$set", bson.D{{"votes", 0}}},
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Updated %v Documents!\n", reset_v.ModifiedCount)
	// alle Streichungen löschen
	reset_c, err := CandidatesCollection.UpdateMany(
		context.TODO(),
		bson.M{},
		bson.D{
			{"$set", bson.D{{"crossed_out", false}}},
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Updated %v Documents!\n", reset_c.ModifiedCount)
	// alle Listenkreuze löschen
	reset_l, err := PartiesCollection.UpdateMany(
		context.TODO(),
		bson.M{},
		bson.D{
			{"$set", bson.D{{"list_vote", false}}},
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Updated %v Documents!\n", reset_l.ModifiedCount)
}
