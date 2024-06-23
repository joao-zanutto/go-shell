package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/shell"
)

func main() {
	s := shell.New()
	for {
		current, _ := os.Getwd()
		fmt.Fprint(os.Stdout, current + "$ ")

		// Read input,
		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		cmd := strings.Split(input, " ")
		cmd[len(cmd)-1] = strings.TrimRight(cmd[len(cmd)-1], "\n")

		s.Execute(cmd[0], cmd[1:])
	}
}
