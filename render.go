package main

import (
	"os"

	"github.com/vmihailenco/msgpack/v5"
)

// SaveAllTextsToMsgPack saves all collected texts to gary-output.msgpack
func SaveAllTextsToMsgPack(texts []string) error {
	data, err := msgpack.Marshal(texts)
	if err != nil {
		return err
	}

	return os.WriteFile("gary-output.msgpack", data, 0644)
}
