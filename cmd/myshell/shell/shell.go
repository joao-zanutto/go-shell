package shell

import (
	"fmt"
	"strconv"
	"os"
	"strings"
)

type shell struct {
	c map[string]func([]string)
	pc map[string]string
}

func (s shell) exit(input []string) {
	exit_code, _ := strconv.Atoi(input[0])
	os.Exit(exit_code)
}

func (s shell) echo(input []string) {
	fmt.Println(strings.Join(input, " "))
}

func (s shell) isBuiltin(command string) bool{
	if s.c[command] != nil {
		return true
	}
	return false
}

func (s shell) getType(input []string) {
	if s.isBuiltin(input[0]) {
		fmt.Println(input[0] + " is a shell builtin")
	} else if s.pc[input[0]] != "" {
		fmt.Println(input[0] + " is " + s.pc[input[0]] + "/" + input[0])
	} else{
		fmt.Println(input[0] + ": not found")
	}
}

func (s *shell) initBuiltin(){
	s.c = map[string]func([]string){
		"echo": s.echo,
		"exit": s.exit,
		"type": s.getType,
	}
}

func (s *shell) initPrograms(){
	dirs := strings.Split(os.Getenv("PATH"), ":")
	for _, dir := range dirs {
		files, _ := os.ReadDir(dir)
		for _, file := range files {
			s.pc[file.Name()] = dir
		}
	}
}

func (s shell) Execute(command string, input []string) {
	if !s.isBuiltin(command) {
		fmt.Println(command + ": command not found")
		return
	}
	s.c[command](input)
}

func New() shell {
	s:= shell { c: make(map[string]func([]string)), pc: make(map[string]string) }
	s.initBuiltin()
	s.initPrograms()
	return s
}

