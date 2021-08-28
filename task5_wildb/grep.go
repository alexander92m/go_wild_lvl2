//subj.txt
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"bufio"
	"io"
)

type  argFlag struct {
	A		int
	B		int
	C		int
	c		bool
	i		bool
	v		bool
	F		bool
	n 		bool
	word	string
	files	[]string
}

func isDigit(c rune) bool {
	if c >= '0' && c <= '9' {
		return true
	}
	return false

}

func parseFlags() argFlag {
	var flags argFlag
	cnt := 0
	for i := 1; i < len(os.Args); i++ {
		if len(os.Args[i]) > 1 && strings.HasPrefix(os.Args[i], "-") {
			if strings.Contains(os.Args[i], "c") {
				flags.c = true
			}
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
			}
			if strings.Contains(os.Args[i], "A") {
				ind := strings.Index(os.Args[i], "A")
				if isDigit(rune(os.Args[i][ind + 1])) {
					flags.A, _ = strconv.Atoi(os.Args[i][2:])
				} else {
					err := fmt.Errorf("UNKNOW FLAG")
					fmt.Println(err)
				}
				flags.i = true
			}
			if strings.Contains(os.Args[i], "B") {
				ind := strings.Index(os.Args[i], "B")
				if isDigit(rune(os.Args[i][ind + 1])) {
					flags.B, _ = strconv.Atoi(os.Args[i][2:])
				} else {
					err := fmt.Errorf("UNKNOW FLAG")
					fmt.Println(err)
				}
				flags.i = true
			}
			if strings.Contains(os.Args[i], "C") {
				ind := strings.Index(os.Args[i], "C")
				if isDigit(rune(os.Args[i][ind + 1])) {
					flags.C, _ = strconv.Atoi(os.Args[i][2:])
				} else {
					err := fmt.Errorf("UNKNOW FLAG")
					fmt.Println(err)
				}
				flags.i = true
			}
		} else if cnt == 0 {
			flags.word = os.Args[i]
			cnt++
		} else {
			flags.files = append(flags.files, os.Args[i])
		}
		
	}
	if len(os.Args) == 1 {
		err := fmt.Errorf("Нет аргументов")
		fmt.Println(err)
	}
	fmt.Println(len(os.Args))
	return flags
}

//CreateArray of strings from files
func createArr(arg argFlag) [][]string{
	strs := make([][]string, 0, 0)
	if len(arg.files) == 0 {
		sc := bufio.NewScanner(os.Stdin)
		for sc.Scan() {
			s := sc.Text()
			var s2 []string
			s2 = strings.Split(s, " ")
			ss := make([]string, 0, 0)
			ss = append(ss, "basic input")
			for k := range s2 {
				ss = append(ss, s2[k])
			}
			strs = append(strs, ss)
		}
	} else {
		for i := range arg.files {
			file,  errOpen := os.Open(arg.files[i])
			if errOpen != nil {
				fmt.Println("Error of open", arg.files[i])
				os.Exit(0)
			}
			rd := bufio.NewReader(file)
			for j := 0; j < 2; {
				s, errRead := rd.ReadString('\n')
				var s2 []string
				if errRead != nil && errRead != io.EOF {
					fmt.Println("Error of read string")
				} else if errRead == io.EOF {
					j = 2
					if s != "\n" {
						
						s2 = strings.Split(s, " ")
					}
				} else {
					s = s[0:len(s) - 1]
					s2 = strings.Split(s, " ")
				}
				ss := make([]string, 0, 0)
				ss = append(ss, arg.files[i])
				for k := range s2 {
					ss = append(ss, s2[k])
				}
				strs = append(strs, ss)
			}
			file.Close()
		}
	}
	
	return strs
}

func outputStrings(strs0 [][]string, arg argFlag) {
	strs := make([][]string, 0, 0)
	strs = append(strs, strs0...)
	
	fmt.Printf("%p %p\n", strs, strs0)
	
	for i := range strs {
		for j := range strs[i] {
			if j != 0 {
				fmt.Print(strs[i][j], " ")
			}
		}
		fmt.Println()
	}
}

func main() {
	arg := parseFlags()
	strs := createArr(arg)
	outputStrings(strs, arg)
	fmt.Println(arg)
	// outputStrings(strs, arg)
}