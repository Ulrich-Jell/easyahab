package main

// import (
// 	"bufio"
// 	"context"
// 	"fmt"
// 	"log"
// 	"os"

// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// 	for {
// 		fmt.Println("Drücken Sie [P], um eine Partei anzulegen")
// 		fmt.Println("Drücken Sie [K], um eine*n Kandidat*in anzulegen.")
// 		//########################
// 		fmt.Println("Drücken Sie [W], um zu wählen.")
// 		fmt.Println("Drücken Sie [V], um das Programm zu verlassen.")
// 		var input string
// 		fmt.Scanln(&input)
// 		if input == "P" || input == "p" {
// 			//Add Party
// 			AddParty()

// 		} else if input == "K" || input == "k" {
// 			//add candidate

// 			fmt.Println("Bitte Namen des*der Kandidat*in eingeben: ")
// 			scannerc := bufio.NewScanner(os.Stdin)
// 			scannerc.Scan() // use `for scanner.Scan()` to keep reading
// 			candidate_name := scannerc.Text()
// 			fmt.Println("Für welche Partei (Kurzform) tritt er*sie an?")
// 			var for_party string
// 			fmt.Scanln(&for_party)

// 			fmt.Println("Bitte geben Sie den Listenplatz ein.")
// 			var position int
// 			fmt.Scanln(&position)

// 			new_candidate := Candidate{position, candidate_name, for_party, 0, false}

// 		} else if input == "W" || input == "w" {

// 			// election process

// 			for votes_left := 16; votes_left > 0; {
// 				println("Sie haben noch", votes_left, "Stimmen.")
// 				fmt.Println("Drücken Sie [L], um ein Listenkreuz zu setzen.")
// 				fmt.Println("Drücken Sie [S], um ein*e Kandidat*in von der Listenwahl auszuschließen")
// 				fmt.Println("Drücken Sie [P], um Personenstimmen zu vergeben.")
// 				fmt.Println("Drücken Sie [V], um die Abstimmung zu abzuschließen.")
// 				var i2 string
// 				fmt.Scanln(&i2)

// 				//vote a person
// 				if i2 == "P" || i2 == "p" {
// 					fmt.Println("Wen möchten Sie Wählen? Bitte geben Sie den Listenplatz an.")
// 					var cast int
// 					fmt.Scanln(&cast)
// 					var cast_check bson.M
// 					if err := CandidatesCollection.FindOne(context.TODO(), bson.M{"position": cast}).Decode(&cast_check); err != nil {
// 						fmt.Println("Da ist was schief gelaufen!")
// 						continue
// 					}
// 					//if votes < 3 {
// 					filter := bson.D{{"position", cast}}
// 					update := bson.D{
// 						{"$inc", bson.D{
// 							{"votes", 1},
// 						}},
// 					}
// 					// } else if votes == 3 {
// 					// 	fmt.println("Der*die Kandidat*in hat bereits drei Stimmen.")
// 					// 	fmt.println("Drücken Sie [A] um abzubrechen oder [0] um die Stimmen auf Null zu setzen.")
// 					// 	var three string
// 					// 	fmt.Scanln(&three)
// 					// 	if three == "a" || three == "A" {
// 					// 		continue
// 					// 	} else if three == "0" {
// 					// filter := bson.D{{"position", three}}
// 					// update := bson.D{
// 					// 	{"$set", bson.D{
// 					// 		{"votes", 0},
// 					// 	}},
// 					// }
// 					// 		votes_left += 3
// 					// 	}
// 					// }

// 					updateResult, err := CandidatesCollection.UpdateOne(context.TODO(), filter, update)
// 					if err != nil {
// 						log.Fatal(err)
// 					}
// 					fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)

// 					votes_left -= 1

// 					//Kandidat streichen
// 				} else if i2 == "S" || i2 == "s" {
// 					fmt.Println("Welche*n Kandidat*in möchten Sie von der Listenwahl ausschließen. Bitte geben Sie den Listenplatz ein")
// 					var cross int
// 					fmt.Scanln(&cross)
// 					var cross_check bson.M
// 					if err := CandidatesCollection.FindOne(context.TODO(), bson.M{"position": cross}).Decode(&cross_check); err != nil {
// 						fmt.Println("Da ist was schief gelaufen!")
// 						continue
// 					}
// 					result, err := CandidatesCollection.UpdateOne(
// 						context.TODO(),
// 						bson.M{"position": cross},
// 						bson.D{
// 							{"$set", bson.D{{"crossed_out", !true}}},
// 						},
// 					)
// 					if err != nil {
// 						fmt.Println("Da ist was schief gelaufen!")
// 						continue
// 					}
// 					fmt.Printf("Updated %v Documents!\n", result.ModifiedCount)

// 				} else if i2 == "L" || i2 == "l" {

// 					//Listenwahl
// 					fmt.Println("Bei welcher Partei oder Wählergemeinschaft möchten Sie ein Listenkreuz setzen?")
// 					fmt.Println("Bitte geben Sie die Kurzform an.")
// 					var list string
// 					fmt.Scanln(&list)
// 					var list_check Party
// 					if err := PartiesCollection.FindOne(context.TODO(), bson.M{"shorthand": list}).Decode(&list_check); err != nil {
// 						fmt.Println("Da ist was schief gelaufen!")
// 						continue
// 					}
// 					fmt.Println(list_check.List_vote)
// 					result, err := PartiesCollection.UpdateOne(
// 						context.TODO(),
// 						bson.M{"shorthand": list},
// 						bson.M{
// 							"$set": bson.M{"list_vote": !list_check.List_vote},
// 						},
// 					)
// 					if err != nil {
// 						fmt.Println("Da ist was schief gelaufen!")
// 						continue
// 					}
// 					fmt.Printf("Updated %v Documents!\n", result.ModifiedCount)

// 					//Wahl beenden
// 				} else if i2 == "V" || i2 == "v" {
// 					fmt.Println("Sie haben bei folgender Partei ein Listenkreuz gesetzt:")
// 					var party_vote bson.M
// 					if err := PartiesCollection.FindOne(context.TODO(), bson.M{"list_vote": true}).Decode(&party_vote); err != nil {
// 						log.Fatal(err)
// 					}
// 					fmt.Println(party_vote)
// 					fmt.Println("##############################")

// 					fmt.Println("Folgende Kandidat*innen sind von der Listenwahl ausgeschlossen")
// 					opts_c := options.Find()
// 					opts_c.SetSort(bson.D{{"position", 1}})
// 					sortCursor, err := CandidatesCollection.Find(context.TODO(), bson.D{
// 						{"votes", bson.D{
// 							{"$gt", 0},
// 						}},
// 					}, opts_c)
// 					var crossedCandidatesSorted []bson.M
// 					if err = sortCursor.All(context.TODO(), &crossedCandidatesSorted); err != nil {
// 						log.Fatal(err)
// 					}
// 					fmt.Println(crossedCandidatesSorted)
// 					fmt.Println("##############################")

// 					fmt.Println("Sie haben für folgende Kandidat*innen gestimmt:")
// 					opts := options.Find()
// 					opts.SetSort(bson.D{{"position", 1}})
// 					sortCursorC, err := CandidatesCollection.Find(context.TODO(), bson.D{
// 						{"votes", bson.D{
// 							{"$gt", 0},
// 						}},
// 					}, opts)
// 					var votedCandidatesSorted []bson.M
// 					if err = sortCursorC.All(context.TODO(), &votedCandidatesSorted); err != nil {
// 						log.Fatal(err)
// 					}
// 					fmt.Println(votedCandidatesSorted)
// 					fmt.Println("##############################")
// 					fmt.Println("Möchten Sie Ihre Wahl ändern oder den Stimmzettel ausdrucken?")
// 					fmt.Println("Drücken Sie [ä] zum Ändern oder [d] zum Drucken.")
// 					var i3 string
// 					fmt.Scanln(&i3)
// 					//drucken
// 					if i3 == "d" || i3 == "D" {
// 						//druckbefehl
// 						break
// 					} else if i3 == "ä" || i3 == "Ä" {
// 						continue
// 					}

// 				}
// 			}

// 		} else if input == "V" || input == "v" {
// 			// alle Personenstimmen auf Null Setzen

// 			break
// 		}
// 	}
// }

// //party ObjectIDs

// //id, _ := primitive.ObjectIDFromHex(
