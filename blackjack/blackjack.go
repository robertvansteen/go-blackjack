package blackjack

import (
	"fmt"
	"os"
	"strconv"

	"github.com/rovansteen/go/terminal"

	"github.com/manifoldco/promptui"
	"github.com/rovansteen/go/deck"
)

type hand struct {
	cards []deck.Card
}

// Game is a blackjack game.
type Game struct {
	Deck           deck.Deck
	DealerHand     hand
	PlayerHand     hand
	PlayerFinished bool
	Finished       bool
}

func blackjackValueOfCard(card deck.Card) int {
	values := map[deck.Value]int{
		deck.One:   1,
		deck.Two:   2,
		deck.Three: 3,
		deck.Four:  4,
		deck.Five:  5,
		deck.Six:   6,
		deck.Seven: 7,
		deck.Eight: 8,
		deck.Nine:  9,
		deck.Ten:   10,
		deck.Jack:  10,
		deck.Queen: 10,
		deck.King:  10,
		deck.Ace:   11,
	}

	return values[card.Value]
}

// Total returns the total of a hand.
func (hand hand) Total() int {
	total := 0
	numberOfAces := 0

	for _, card := range hand.cards {
		value := blackjackValueOfCard(card)

		if card.Value == deck.Ace {
			numberOfAces++
		} else {
			total += value
		}
	}

	if numberOfAces > 0 {
		for numberOfAces > 1 {
			total++
			numberOfAces--
		}
		if total+11 > 21 {
			total++
		} else {
			total += 11
		}
	}

	return total
}

func (hand hand) IsBust() bool {
	return hand.Total() > 21
}

// Add adds a card to the hand.
func (hand *hand) Add(card deck.Card) {
	hand.cards = append(hand.cards, card)
}

// NewGame creates a new blackjack game.
func NewGame() Game {
	deck := deck.NewDeck()
	dealerHand := hand{cards: nil}
	playerHand := hand{cards: nil}

	return Game{
		Deck:           deck,
		Finished:       false,
		PlayerFinished: false,
		DealerHand:     dealerHand,
		PlayerHand:     playerHand,
	}
}

// Start starts a game of Blackjack.
func (game *Game) Start() {
	game.Deck.Shuffle()
	cards, error := game.Deck.Draw(3)
	if error != nil {
		panic(error)
	}
	game.PlayerHand.Add(cards[0])
	game.PlayerHand.Add(cards[1])
	game.DealerHand.Add(cards[2])
}

func (hand hand) String() string {
	var string string

	for index, card := range hand.cards {
		if index == 0 {
			string += "["
		}

		string += card.String()

		if index == len(hand.cards)-1 {
			string += "] "
		} else {
			string += " "
		}
	}

	string += "Total: " + strconv.Itoa(hand.Total())

	return string
}

func (game *Game) String() [2]string {
	return [2]string{
		"Dealer: " + game.DealerHand.String(),
		"Player: " + game.PlayerHand.String(),
	}
}

// PrintSummary prints a summary of the game.
func (game *Game) PrintSummary() {
	for _, line := range game.String() {
		fmt.Println(line)
	}

	if game.Finished {
		if game.PlayerHand.IsBust() {
			fmt.Println("Player busts, dealer wins!")
		} else if game.DealerHand.IsBust() {
			fmt.Println("Dealer busts, player wins!")
		} else if game.PlayerHand.Total() > game.DealerHand.Total() {
			fmt.Println("Player's hand beats the dealer. Player wins.")
		} else if game.PlayerHand.Total() < game.DealerHand.Total() {
			fmt.Println("Dealer's hand beats the player. Dealer wins.")
		} else {
			fmt.Println("Push.")
		}
	}
}

// Hit a card for the player.
func (game *Game) Hit() {
	cards, error := game.Deck.Draw(1)
	if error != nil {
		panic(error)
	}
	game.PlayerHand.Add(cards[0])

	if game.PlayerHand.Total() > 21 {
		game.PlayerFinished = true
	}
}

// Stand a card for the player.
func (game *Game) Stand() {
	game.PlayerFinished = true

}

// DealersTurn deals the cards for the dealer.
func (game *Game) DealersTurn() {
	for !game.Finished {
		if game.DealerHand.Total() >= 17 {
			game.Finished = true
			break
		}

		cards, error := game.Deck.Draw(1)
		if error != nil {
			panic(error)
		}

		game.DealerHand.Add(cards[0])
	}
}

func prompt() string {
	prompt := promptui.Select{
		Label: "Select action",
		Items: []string{"Hit", "Stand"},
	}

	_, result, err := prompt.Run()

	if err != nil {
		os.Exit(0)
	}

	return result
}

func (game Game) getWinner() (winner string) {
	if game.Finished {
		if game.PlayerHand.IsBust() {
			return "dealer"
		} else if game.DealerHand.IsBust() {
			return "player"
		} else if game.PlayerHand.Total() > game.DealerHand.Total() {
			return "player"
		} else if game.PlayerHand.Total() < game.DealerHand.Total() {
			return "dealer"
		}
	}

	return ""
}

// Play a game of Blackjack.
func Play() string {
	game := NewGame()
	game.Start()

	for !game.PlayerFinished {
		terminal.Flush()
		game.PrintSummary()

		action := prompt()

		switch action {
		case "Hit":
			game.Hit()
		case "Stand":
			game.Stand()
		}
	}

	game.DealersTurn()

	terminal.Flush()
	game.PrintSummary()

	return game.getWinner()
}
