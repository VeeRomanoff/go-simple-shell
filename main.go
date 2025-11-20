package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"shell/inp"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {

		raw, err := os.Getwd()
		if err != nil {
			log.Fatalf(err.Error())
		}

		dirS := strings.Split(raw, " ")

		fmt.Println(dirS[0])

		fmt.Print("> ")

		// Reading the keyboard input
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
		}

		if err := inp.ExecInput(input); err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
		}
	}
}
