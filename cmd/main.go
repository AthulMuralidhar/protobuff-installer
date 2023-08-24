package main

import (
	"bufio"
	"fmt"
	"os"
)

// GOALS
// - make a cmd line interface to install protobuf and protobuf compiler
// - so the program should have an interactive interface which asks for the protobuf version and the language and the arch etc, all the important stuff
// - same with the protoc compiler, version arch full or partial etc
// - if left empty, just used the latest and install
// it should also install everything at the usr level without sudo perms unless explicitely mentioned

const PB_REL = "https://github.com/protocolbuffers/protobuf/releases"

//$ curl -LO $PB_REL/download/v3.15.8/protoc-3.15.8-linux-x86_64.zip

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("would you like to install the proto compiler? (y/n): ")
		// Scans a line from Stdin(Console)
		scanner.Scan()
		// Holds the string that scanned
		text := scanner.Text()
		if len(text) != 0 {
			if text == "y" {

				// FIXME: this kinda prints with the same promt above
				// the googled answer is here: https://stackoverflow.com/questions/20895552/how-can-i-read-from-standard-input-in-the-console
				fmt.Print("what version of protoc would you like to install? (v24.1): ")
				text = scanner.Text()
				if len(text) != 0 {
					fmt.Print(text)
				}
			}
		}
	}

}
