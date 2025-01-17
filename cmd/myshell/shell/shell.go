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

func (s shell) cd(args []string) {
	if strings.HasPrefix(args[0], "/") {
		if err := os.Chdir(args[0]); err != nil {
			fmt.Println(args[0] + ": No such file or directory")
		}
		return
	}
	if !strings.HasSuffix(args[0], "/") {
		args[0] = args[0] + "/"
	}
	target_array := strings.Split(args[0], "/")
	current, _ := os.Getwd()
	current_array := strings.Split(current, "/")
	for idx, path := range target_array {
		switch path {
		case "..":
			current_array = current_array[:len(current_array)-1]
			continue
		case ".":
			continue
		case "~":
			current_array = strings.Split(os.Getenv("HOME"), "/")
			continue
		default:
			resolved := strings.Join(current_array, "/") + "/" + strings.Join(target_array[idx:], "/")
			if err := os.Chdir(resolved); err != nil {
				fmt.Println(resolved + ": No such file or directory")
			}
			return
		}
	}
}

func (s shell) pwd(args []string) {
	dir, _ := os.Getwd()
	fmt.Println(dir)
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
	s.c["pwd"] = s.pwd
	s.c["cd"] = s.cd
	dirs := strings.Split(os.Getenv("PATH"), ":")
	for _, dir := range dirs {
		files, _ := os.ReadDir(dir)
		for _, file := range files {
			s.pc[file.Name()] = dir
		}
	}
	return s
}
