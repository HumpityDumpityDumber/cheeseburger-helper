package main

import (
	"bytes"
	"fmt"
)

// escapeNonPrintable converts control bytes into visible escape sequences.
func escapeNonPrintable(s string) string {
	var b bytes.Buffer
	for i := 0; i < len(s); i++ {
		c := s[i]
		// printable ASCII (space..~)
		if c >= 0x20 && c <= 0x7e {
			b.WriteByte(c)
			continue
		}
		switch c {
		case '\n':
			b.WriteString("\\n")
		case '\r':
			b.WriteString("\\r")
		case '\t':
			b.WriteString("\\t")
		case 0x1b: // ESC
			// make ESC visible as "\x1b" so sequences like ESC + '[' don't get interpreted
			b.WriteString("\\x1b")
		default:
			b.WriteString(fmt.Sprintf("\\x%02x", c))
		}
	}
	return b.String()
}
