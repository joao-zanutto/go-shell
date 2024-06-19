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

		// Get user input, split it based on whitespaces and trim newline from last element
		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		cmd := strings.Split(input, " ")
		cmd[len(cmd)-1] = strings.TrimRight(cmd[len(cmd)-1], "\n")

		// Run bult-in commands
		if cmd[0] == "exit" {
			exit_code, _ := strconv.Atoi(cmd[1])
			os.Exit(exit_code)
		}
		if cmd[0] == "echo" {
			fmt.Fprint(os.Stdout, strings.Join(cmd[1:], " "))
			continue
		}

		fmt.Fprint(os.Stdout, input[:len(input)-1]+": command not found\n")
	}
}
