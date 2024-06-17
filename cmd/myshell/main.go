package main

import (
	"bufio"
	"strconv"

	// Uncomment this block to pass the first stage
	"fmt"
	"os"

	"strings"
)

func main() {
	for {
		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input
		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')

		cmd := strings.Split(input, " ")
		if cmd[0] == "exit" {
			exit_code, _ := strconv.Atoi(cmd[1])
			os.Exit(exit_code)
		}

		fmt.Fprint(os.Stdout, input[:len(input)-1]+": command not found\n")

	}
}
