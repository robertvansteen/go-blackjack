package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/rovansteen/go/terminal"

	"github.com/manifoldco/promptui"
	"github.com/rovansteen/go/blackjack"
)

func prompt() string {
	prompt := promptui.Select{
		Label: "Select action",
		Items: []string{"Play new game", "Quit"},
	}

	_, result, err := prompt.Run()

	if err != nil {
		os.Exit(0)
	}

	return result
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	terminal.Flush()

	score := map[string]int{"player": 0, "dealer": 0}

	for {
		fmt.Println("[Score]", "Player:", score["player"], "Dealer:", score["dealer"])
		action := prompt()
		switch action {
		case "Play new game":
			winner := blackjack.Play()
			score[winner]++
		case "Quit":
			fmt.Println("Thanks for playing, bye!")
			os.Exit(0)
		}
	}
}
