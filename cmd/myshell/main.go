package main

import (
	"bufio"
	// Uncomment this block to pass the first stage
	"fmt"
	"os"
)

func main() {
	for {
		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input
		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		fmt.Fprint(os.Stdout, input[:len(input)-1]+": command not found\n")
	}
}
