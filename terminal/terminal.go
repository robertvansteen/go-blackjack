package terminal

import (
	"fmt"

	"github.com/buger/goterm"
)

// Flush the screen
func Flush() {
	goterm.Clear()
	goterm.MoveCursor(1, 1)
	goterm.Flush()
	fmt.Println("=============================")
	fmt.Println("Blackjack")
	fmt.Println("=============================")
}
