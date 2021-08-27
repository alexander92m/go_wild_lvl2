package main

import (
	"fmt"
	"os"
	"strings"
)

type  argFlag struct {
	A int
	B int
	C int
	c int
	i bool
	v bool
	F bool
	n bool
}

func parseFlags() argFlag {
	var flags argFlag

	for i := 1; i < len(os.Args); i++ {
		if len(os.Args[i]) > 1 && strings.HasPrefix(os.Args[i], "-") {
			if strings.Contains(os.Args[i], "i") {
				flags.i = true
			}
			if strings.Contains(os.Args[i], "v") {
				flags.v = true
			}
			if strings.Contains(os.Args[i], "F") {
				flags.F = true
			}
			if strings.Contains(os.Args[i], "n") {
				flags.n = true
			} else {
				fmt.Errorf("аргумент не аргумент")
			}
			// if strings.Contains(os.Args[i], "A") {
			// 	if 
			// 	flags.i = true
			// }
		} else if len(os.Args) == 1 {

		}
		fmt.Println("i =", i, os.Args[i])
	}
	return flags
}


func main() {
	flags := parseFlags()
	fmt.Println(flags)
}