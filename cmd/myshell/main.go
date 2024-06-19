package main

import (
	"bufio"
	"strconv"

	"fmt"
	"os"

	"strings"
)

func exit(input []string) {
	exit_code, _ := strconv.Atoi(input[0])
	os.Exit(exit_code)
}

func echo(input []string) {
	fmt.Println(strings.Join(input, " "))
}

func getType(input []string, c map[string]func([]string)) {
	if c[input[0]] != nil || input[0] == "type" {
		fmt.Println(input[0] + " is a shell builtin")
	} else {
		fmt.Println(input[0] + ": not found")
	}
}

func main() {
	c := map[string]func([]string){
		"echo": echo,
		"exit": exit,
	}

	for {
		fmt.Fprint(os.Stdout, "$ ")

		// Read input,
		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		cmd := strings.Split(input, " ")
		cmd[len(cmd)-1] = strings.TrimRight(cmd[len(cmd)-1], "\n")

		if c[cmd[0]] == nil {
			if(cmd[0]) != "type" {
				fmt.Println(cmd[0] + ": command not found")
				continue
			}
		}

		switch cmd[0]{
		case "type": getType(cmd[1:], c)
		default: c[cmd[0]](cmd[1:])
		}

	}
}
