package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/vmihailenco/msgpack/v5"
)

// GetTextFromFiles returns text data (in order) for the provided file path.
// If fileArg is empty it returns an empty slice.
func GetTextFromFiles(fileArg string) ([]string, error) {
	if fileArg == "" {
		return []string{}, nil
	}

	fi, err := os.Stat(fileArg)
	if err != nil {
		return nil, fmt.Errorf("stat %s: %w", fileArg, err)
	}

	var allTexts []string

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
			if ext != ".msgpack" && ext != ".mp" {
				continue
			}
			path := filepath.Join(fileArg, name)

			texts, err := loadMsgPackFile(path)
			if err != nil {
				return nil, err
			}
			allTexts = append(allTexts, texts...)
		}
	} else {
		// single file
		texts, err := loadMsgPackFile(fileArg)
		if err != nil {
			return nil, err
		}
		allTexts = append(allTexts, texts...)
	}

	return allTexts, nil
}

// loadMsgPackFile loads a MessagePack file and returns the string slice
func loadMsgPackFile(filePath string) ([]string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("read file %s: %w", filePath, err)
	}

	var texts []string
	err = msgpack.Unmarshal(data, &texts)
	if err != nil {
		return nil, fmt.Errorf("unmarshal msgpack %s: %w", filePath, err)
	}

	return texts, nil
}
