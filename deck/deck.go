package deck

import (
	"errors"
	"math/rand"
)

// Suit of the card.
type Suit int

const (
	club Suit = iota
	diamond
	heart
	spade
)

// Suits is the possible suit values of a card.
var Suits = []Suit{
	club,
	diamond,
	heart,
	spade,
}

// Value of the card.
type Value int

// Enum for the card values.
const (
	One Value = iota
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	Ace
)

// Values are the possible values of a card.
var Values = []Value{
	One,
	Two,
	Three,
	Four,
	Five,
	Six,
	Seven,
	Eight,
	Nine,
	Ten,
	Jack,
	Queen,
	King,
	Ace,
}

// Card is a playing card
type Card struct {
	Suit  Suit
	Value Value
}

// Deck holds a deck of cards
type Deck struct {
	cards []Card
}

// NewDeck creates a new deck.
func NewDeck() Deck {
	deck := Deck{cards: nil}
	for _, suit := range Suits {
		for _, value := range Values {
			deck.cards = append(deck.cards, Card{Suit: suit, Value: value})
		}
	}
	return deck
}

// Shuffle shuffles the deck.
func (deck *Deck) Shuffle() {
	rand.Shuffle(len(deck.cards), func(i, j int) {
		deck.cards[i], deck.cards[j] = deck.cards[j], deck.cards[i]
	})
}

// Draw draws a card from the deck.
func (deck *Deck) Draw(count int) (cards []Card, error error) {
	if count > len(deck.cards) {
		return nil, errors.New("deck does not have enough cards left to draw the specified amount")
	}

	hand := deck.cards[0:count]
	deck.cards = deck.cards[count:]

	return hand, nil
}

// String returns a string representation of the card.
func (card Card) String() string {
	suits := [...]string{"♣", "♦", "♥", "♠"}
	values := [...]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}
	return suits[card.Suit] + values[card.Value]
}
