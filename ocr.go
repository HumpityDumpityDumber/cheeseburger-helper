package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

// GetTextFromFiles returns OCR results (in order) for the provided file path.
// If fileArg is empty it returns an empty slice.
func GetTextFromFiles(fileArg string) ([]string, error) {
	if fileArg == "" {
		return []string{}, nil
	}

	fi, err := os.Stat(fileArg)
	if err != nil {
		return nil, fmt.Errorf("stat %s: %w", fileArg, err)
	}

	var texts []string

	if fi.IsDir() {
		entries, err := os.ReadDir(fileArg)
		if err != nil {
			return nil, fmt.Errorf("read dir: %w", err)
		}

		var names []string
		for _, e := range entries {
			if e.IsDir() {
				continue
			}
			names = append(names, e.Name())
		}
		sort.Strings(names)

		for _, name := range names {
			ext := strings.ToLower(filepath.Ext(name))
			switch ext {
			case ".png", ".jpg", ".jpeg", ".bmp", ".tif", ".tiff", ".gif":
			default:
				continue
			}
			path := filepath.Join(fileArg, name)

			// run tesseract CLI, output to stdout, use PSM 10 for single-character mode
			// command: tesseract <image> stdout --psm 10
			cmd := exec.Command("tesseract", path, "stdout", "--psm", "10")
			out, err := cmd.Output()
			if err != nil {
				// capture stderr if available for better error message
				if ee, ok := err.(*exec.ExitError); ok {
					return nil, fmt.Errorf("ocr %s: %s", path, strings.TrimSpace(string(ee.Stderr)))
				}
				return nil, fmt.Errorf("ocr %s: %w", path, err)
			}
			txt := string(out)
			// convert any literal escape sequences returned by OCR into real bytes
			txt = unescapeEscapes(txt)
			txt = strings.TrimSpace(txt)
			if txt == "" {
				txt = " "
			}
			texts = append(texts, txt)
		}
	} else {
		// single file
		cmd := exec.Command("tesseract", fileArg, "stdout", "--psm", "10")
		out, err := cmd.Output()
		if err != nil {
			if ee, ok := err.(*exec.ExitError); ok {
				return nil, fmt.Errorf("ocr %s: %s", fileArg, strings.TrimSpace(string(ee.Stderr)))
			}
			return nil, fmt.Errorf("ocr %s: %w", fileArg, err)
		}
		txt := string(out)
		// convert any literal escape sequences returned by OCR into real bytes
		txt = unescapeEscapes(txt)
		txt = strings.TrimSpace(txt)
		if txt == "" {
			txt = " "
		}
		texts = append(texts, txt)
	}

	return texts, nil
}

func unescapeEscapes(s string) string {
	var b strings.Builder
	for i := 0; i < len(s); i++ {
		if s[i] != '\\' || i+1 >= len(s) {
			b.WriteByte(s[i])
			continue
		}
		i++ // skip backslash
		switch s[i] {
		case 'n':
			b.WriteByte('\n')
		case 'r':
			b.WriteByte('\r')
		case 't':
			b.WriteByte('\t')
		case '\\':
			b.WriteByte('\\')
		case 'x':
			// expect two hex digits after 'x'
			if i+2 < len(s) {
				hex := s[i+1 : i+3]
				if v, err := strconv.ParseUint(hex, 16, 8); err == nil {
					b.WriteByte(byte(v))
					i += 2 // consume hex digits
				} else {
					// malformed \x -- write literally
					b.WriteString("\\x")
				}
			} else {
				// not enough chars for \xHH, write literally
				b.WriteString("\\x")
			}
		default:
			// unknown escape, write character as-is
			b.WriteByte(s[i])
		}
	}
	return b.String()
}
