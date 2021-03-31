package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"time"
)

type Party struct {
	Name      string
	Shorthand string
	List_vote bool
}

type Candidate struct {
	Position    int
	Name        string
	Party       string
	Votes       int
	Crossed_out bool
}

func AddParty() {
	fmt.Println("Bitte Namen der Partei eingeben: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() // use `for scanner.Scan()` to keep reading
	party_name := scanner.Text()
	fmt.Println("Bitte Kurzform eingeben: ")
	var shorthand string
	fmt.Scanln(&shorthand)

	new_party := Party{party_name, shorthand, false}

	//fmt.Println(new_party)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := Client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	insertResult, err := PartiesCollection.InsertOne(ctx, new_party)
	if err != nil {
		fmt.Println("Da ist was schief gelaufen!")
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
}
