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

func (s shell) exit(args []string) {
	os.Exit(0)
}

func (s shell) echo(args []string) {
	fmt.Println(strings.Join(args, " "))
}

func (s shell) getType(args []string) {
	if s.c[args[0]] != nil {
		fmt.Println(args[0] + " is a shell builtin")
	} else if s.pc[args[0]] != "" {
		fmt.Println(args[0] + " is " + s.pc[args[0]] + "/" + args[0])
	} else {
		fmt.Println(args[0] + ": not found")
	}
}

func (s shell) Execute(command string, args []string) {
	if s.c[command] != nil {
		s.c[command](args)
	} else {
		cmd := exec.Command(command, args...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Println(command + ": command not found")
		}
	}
}

func New() shell {
	s:= shell { c: make(map[string]func([]string)), pc: make(map[string]string) }
	s.c["echo"] = s.echo
	s.c["exit"] = s.exit
	s.c["type"] = s.getType
	dirs := strings.Split(os.Getenv("PATH"), ":")
	for _, dir := range dirs {
		files, _ := os.ReadDir(dir)
		for _, file := range files {
			s.pc[file.Name()] = dir
		}
	}
	return s
}
