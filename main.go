package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	var fileArg string
	flag.StringVar(&fileArg, "file", "", "optional msgpack file or directory path")
	flag.Parse()

	fmt.Println(" ::::::::      :::     :::::::::  :::   ::: ")
	time.Sleep(100 * time.Millisecond)
	fmt.Println(":+:    :+:   :+: :+:   :+:    :+: :+:   :+: ")
	time.Sleep(100 * time.Millisecond)
	fmt.Println("+:+         +:+   +:+  +:+    +:+  +:+ +:+  ")
	time.Sleep(100 * time.Millisecond)
	fmt.Println(":#:        +#++:++#++: +#++:++#:    +#++:   \n+#+   +#+# +#+     +#+ +#+    +#+    +#+    ")
	time.Sleep(100 * time.Millisecond)
	fmt.Println("#+#    #+# #+#     #+# #+#    #+#    #+#    ")
	time.Sleep(100 * time.Millisecond)
	fmt.Println(" ########  ###     ### ###    ###    ###    ")
	time.Sleep(100 * time.Millisecond)
	time.Sleep(2 * time.Second)

	// get text from provided msgpack file(s) (may be empty)
	texts, err := GetTextFromFiles(fileArg)
	if err != nil {
		fmt.Println(err)
		return
	}

	// run interactive append loop
	texts, err = InteractiveAppend(texts)
	if err != nil {
		fmt.Println(err)
		return
	}

	// save everything as gary-output.msgpack
	if err := SaveAllTextsToMsgPack(texts); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("saved to gary-output.msgpack\n")
}
