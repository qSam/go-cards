package main

import (
	"fmt"
)

// Create a new type of 'deck'
// which is a slice of strings

type deck []string

func newDeck() deck {

	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for i, suit := range cardSuits {
		for j, value := range cardValues {
			cards = append(cards, suit+" of "+value)
		}
	}
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
