package main

import (
	"fmt"
	"os"
	"strings"
)
type Arguments struct {
	f int
	d string
	s bool
}

func parseArgs() Arguments {
	var args Arguments
	for i := 1; i < len(os.Args); i++ {
		if len(os.Args[i] > 0) && strings.HasPrefix(os.Agrs[i], "-") {
			if strings.Contains(os.Args[i], "s") {
				args.s = true
			}
			if strings.Contains(os.Args[i], "d") {
				args.s = true
			}
		}
	}
}

func main() {
	fmt.Println(1)
	args := parseAgrs()
	
}