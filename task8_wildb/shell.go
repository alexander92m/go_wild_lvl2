package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Command struct {
	Name string

}

type Status struct {
	Dir string

}

//parseStatus выясняет переменные для текущего каталога
func parseStatus(stat Status) {
	
}

//parseCmd выясняет какая команда введена
func parseCmd(input string, cmd *Command) {
	if input == "cd" {
		cmd.Name = "cd"
	} else if input == "pwd" {
		cmd.Name = "pwd"
	} else if input == "echo" {
		cmd.Name = "echo"
	} else if input == "kill" {
		cmd.Name = "kill"
	} else if input == "ps" {
		cmd.Name = "ps"
	} else {
		cmd.Name = "error, unknown command"
	}
}

func main() {
	var stat Status
	var cmd Command
	reader := bufio.NewReader(os.Stdin)
	for {
		cmd.Name = ""
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			break
		}
		input = input[:strings.Index(input, "\n")]

		parseCmd(input, &cmd)
		parseStatus(stat)
		fmt.Printf("input=%v, &cmd=%p, cmd=%v, stat=%v\n", input, &cmd, cmd, stat)
	}
}