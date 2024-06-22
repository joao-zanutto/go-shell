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
		fmt.Fprint(os.Stdout, "$ ")

		// Read input,
		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		cmd := strings.Split(input, " ")
		cmd[len(cmd)-1] = strings.TrimRight(cmd[len(cmd)-1], "\n")

		s.Execute(cmd[0], cmd[1:])
	}
}
