package main

import (
	"fmt"
	"os"

	"golang.org/x/term"
)

func main() {
	fmt.Print("\033[?1049h")   // switch to alternate buffer
	fmt.Print("\033[H\033[2J") // clear
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}

	buf := make([]byte, 1)
	text := []string{}

	for {
		n, err := os.Stdin.Read(buf)
		if err != nil || n == 0 {
			break
		}
		if buf[0] == 19 { // ctrl s
			break
		}

		var s string
		switch buf[0] {
		case 13: // enter
			s = "\n"
		case 127: // backspace
			s = "\x1b[D\x1b[K"
		default:
			s = string(buf[0])
		}

		text = append(text, s)
		fmt.Print(s)
	}

	// restore terminal mode and switch back to the main screen before printing
	term.Restore(int(os.Stdin.Fd()), oldState) // make sure we're in normal (cooked) mode
	fmt.Print("\033[?1049l")                   // switch back to main buffer

	fmt.Println("\nBack in main screen buffer. You typed:")
	for _, line := range text {
		fmt.Print(line)
	}
}
