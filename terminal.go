package main

import (
	"fmt"
	"os"

	"golang.org/x/term"
)

// InteractiveAppend shows an alternate terminal buffer, reads single bytes,
// appends them to the provided slice, and prints raw terminal output.
// Returns final slice (or an error).
func InteractiveAppend(initial []string) ([]string, error) {
	text := append([]string{}, initial...)

	fmt.Print("\033[?1049h")   // switch to alternate buffer
	fmt.Print("\033[H\033[2J") // clear

	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		// try to restore main buffer on failure
		fmt.Print("\033[?1049l")
		return text, err
	}
	// ensure restore on exit
	defer func() {
		term.Restore(int(os.Stdin.Fd()), oldState)
		fmt.Print("\033[?1049l")
	}()

	for _, part := range text {
		// print raw to terminal (do not escape) — only escape when saving images
		fmt.Print(part)
	}

	buf := make([]byte, 1)

	for {
		n, err := os.Stdin.Read(buf)
		if err != nil || n == 0 {
			break
		}
		if buf[0] == 19 { // ctrl-s to finish
			break
		}

		var s string
		switch buf[0] {
		case 13: // enter
			s = "\r\n"
		case 127: // backspace
			// represent backspace as the same escape sequence used previously
			s = "\x1b[D\x1b[K"
		default:
			s = string(buf[0])
		}

		text = append(text, s)

		fmt.Print("\033[H\033[2J") // clear term
		for _, part := range text {
			// print raw to terminal (do not escape) — only escape when saving images
			fmt.Print(part)
		}
	}

	// final clear of alt buffer before returning to main buffer
	fmt.Print("\033[H\033[2J")
	return text, nil
}
