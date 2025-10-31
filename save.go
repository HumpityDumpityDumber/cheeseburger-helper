package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// SaveAll wipes out outDir and writes each text item as numbered PNGs.
func SaveAll(outDir string, texts []string) error {
	// remove existing dir if present
	if _, err := os.Stat(outDir); err == nil {
		if err := os.RemoveAll(outDir); err != nil {
			return fmt.Errorf("remove output dir: %w", err)
		}
	}
	if err := os.MkdirAll(outDir, 0o755); err != nil {
		return fmt.Errorf("mkdir output: %w", err)
	}

	for i, item := range texts {
		escaped := escapeNonPrintable(item)
		fname := filepath.Join(outDir, fmt.Sprintf("%d.png", i+1))
		if err := RenderTextToPNG(escaped, fname); err != nil {
			// continue saving others but report error
			return fmt.Errorf("save png %s: %w", fname, err)
		}
	}
	return nil
}
