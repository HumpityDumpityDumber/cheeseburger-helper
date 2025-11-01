package main

import (
	"bytes"
	"os"

	"github.com/andybalholm/brotli"
	"github.com/vmihailenco/msgpack/v5"
)

// SaveAllTextsToMsgPack saves all collected texts to gary-output.msgpack with Brotli compression
func SaveAllTextsToMsgPack(texts []string) error {
	data, err := msgpack.Marshal(texts)
	if err != nil {
		return err
	}

	// Compress with Brotli
	var buf bytes.Buffer
	brotliWriter := brotli.NewWriterLevel(&buf, brotli.BestCompression)

	_, err = brotliWriter.Write(data)
	if err != nil {
		return err
	}

	err = brotliWriter.Close()
	if err != nil {
		return err
	}

	return os.WriteFile("gary-output.msgpackz", buf.Bytes(), 0644)
}
