package main

import (
	"bufio"
	"fmt"
	// "log"
	"os"
	"os/exec"
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
func parseCmd(input string) Command {
	var cmd Command
	inputArr := strings.Split(input, " ")
	if inputArr[0] == "cd" {
		cmd.Name = "cd"
	} else if inputArr[0] == "pwd" {
		cmd.Name = "pwd"
	} else if inputArr[0] == "echo" {
		cmd.Name = "echo"
	} else if inputArr[0] == "kill" {
		cmd.Name = "kill"
	} else if inputArr[0] == "ps" {
		cmd.Name = "ps"
	} else if inputArr[0] == "cat" {
		cmd.Name = "cat"
	} else if inputArr[0] == "" {
		cmd.Name = ""
	} else {
		cmd.Name = "error, unknown command"
	}
	return cmd
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		s, _ := os.Hostname()

		s = strings.Split(s, ".")[0]
		fmt.Printf("%v%% ", s)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			break
		}
		input = input[:strings.Index(input, "\n")]
		cmd := parseCmd(input)
		
		if cmd.Name == "" {
			
		} else {
			out, err := exec.Command(cmd.Name).Output()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
			fmt.Print(string(out))
			
			// cmd1 := exec.Command(cmd.Name)
			
			// err := cmd1.Run()
		
			// if err != nil {
			// 	log.Println(err)
			// }
		}
	}
}