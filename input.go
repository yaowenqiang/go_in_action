package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func binaryToDecimal() {
	// TODO finish the binary to decimal method
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Simple Shell")
	fmt.Println("--------------------")
	for {
		fmt.Println("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)
		if strings.Compare("hi", text) == 0 {
			fmt.Println("Hello, yourself")
		}
	}
}
