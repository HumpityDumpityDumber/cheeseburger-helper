package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	var fileArg string
	flag.StringVar(&fileArg, "file", "", "optional file or directory path")
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

	// get text from provided image(s) (may be empty)
	texts, err := GetTextFromFiles(fileArg)
	if err != nil {
		fmt.Println(err)
		return
	}

	// run interactive append loop (renders escaped output while typing)
	texts, err = InteractiveAppend(texts)
	if err != nil {
		fmt.Println(err)
		return
	}

	// save everything into output/
	if err := SaveAll("gary-output", texts); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("files saved to %s\n", "gary-output/")
}
