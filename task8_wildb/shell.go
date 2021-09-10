package main

import (
	"bufio"
	"fmt"
	"os"
)

type Command struct {
    Name string
}

type Status struct {
    Dir string

}

func parseDir(com Command) {
    
}

func parseCmd(input string) {
    if input == "ls" {
        var cmd Command
        cmd.Name = "ls"
    }
}

func main() {
    var stat Status
    reader := bufio.NewReader(os.Stdin)
    for {
        
        
        input, err := reader.ReadString('\n')
        if err != nil {
            fmt.Fprintln(os.Stderr, err)
            break
        }

        com := parseCmd(input)
        parseStatus(stat)

		fmt.Printf("input=%v, err=%v T=%T\n", input, err, err)
        
    }
}