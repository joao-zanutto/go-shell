package shell

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type shell struct {
	c map[string]func([]string)
	pc map[string]string
}

func (s shell) exit(input []string) {
	os.Exit(0)
}

func (s shell) echo(input []string) {
	fmt.Println(strings.Join(input, " "))
}

func (s shell) getType(input []string) {
	if s.c[input[0]] != nil {
		fmt.Println(input[0] + " is a shell builtin")
	} else if s.pc[input[0]] != "" {
		fmt.Println(input[0] + " is " + s.pc[input[0]] + "/" + input[0])
	}
	fmt.Println(input[0] + ": not found")
}

func (s shell) Execute(command string, input []string) {
	if s.c[command] != nil {
		s.c[command](input)
	}
	cmd := exec.Command(s.pc[command] + "/" + command, input...)
	err := cmd.Run()
	if err != nil {
		fmt.Println(command + ": command not found")
	}
}

func New() shell {
	s:= shell { c: make(map[string]func([]string)), pc: make(map[string]string) }

	// Load builtins
	s.c["echo"] = s.echo
	s.c["exit"] = s.exit
	s.c["type"] = s.getType

	// Load programs from PATH
	dirs := strings.Split(os.Getenv("PATH"), ":")
	for _, dir := range dirs {
		files, _ := os.ReadDir(dir)
		for _, file := range files {
			s.pc[file.Name()] = dir
		}
	}
	return s
}
