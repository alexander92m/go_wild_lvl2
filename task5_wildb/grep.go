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
				if len(os.Args[i]) < 3 || !isDigit(rune(os.Args[i][ind + 1])) {
					err := fmt.Errorf("UNKNOW FLAG")
					fmt.Println(err)
					
				} else {
					flags.A, _ = strconv.Atoi(os.Args[i][2:])
				}
				flags.i = true
			}
			if strings.Contains(os.Args[i], "B") {
				ind := strings.Index(os.Args[i], "B")
				if len(os.Args[i]) < 3 || !isDigit(rune(os.Args[i][ind + 1])) {
					err := fmt.Errorf("UNKNOW FLAG")
					fmt.Println(err)
					
				} else {
					flags.B, _ = strconv.Atoi(os.Args[i][2:])
				}
				flags.i = true
			}
			if strings.Contains(os.Args[i], "C") {
				ind := strings.Index(os.Args[i], "C")
				if len(os.Args[i]) < 3 || !isDigit(rune(os.Args[i][ind + 1])) {
					err := fmt.Errorf("UNKNOW FLAG")
					fmt.Println(err)
					
				} else {
					flags.C, _ = strconv.Atoi(os.Args[i][2:])
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

func arrarrContain(ss []string, arg argFlag) bool {
	for i := range ss {
		if strings.Contains(ss[i], arg.word) {
			return true
		}
	}
	return false
}

func reArray(ss0 [][]string, arg argFlag) [][]string {
	ss := make([][]string, 0, 0)
	ss2 := make([][]string, 0, 0)
	
	ss = append(ss, ss0...)
	for i := range ss {
		match := arrarrContain(ss[i], arg)
		if match {
			// if len(ss2) != 0 && ss2[len(ss2) - 1] {

			// }
			ss2 = append(ss2, ss[i])
			match2 := false
			for k := 1; !match2 && k <= arg.A && k + i < len(ss); k++ {
				match2 = false
				for l := 1; l < len(ss[i + k]); l++ {
					if strings.Contains(ss[i + k][l], arg.word) {
						match2 = true
						break
					}
				}
				if !match2 {
					ss2 = append(ss2, ss[i + k])
				}
			}
		}
		
	}
	return ss2
}

func outputStrings(strs [][]string, arg argFlag) {
		
	fmt.Println("len(strs)", len(strs))
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
	if len(os.Args) < 3 {
		fmt.Println("некорретные входные данные\n КАК?: grep [-civFn] [-A num] [-B num] [-C[num] [filter key] [file ...]\n-A - `after` печатать +N строк после совпадения\n-B - `before` печатать +N строк до совпадения\n-C - `context` (A+B) печатать ±N строк вокруг совпадения\n-c - `count` (количество строк)\n-i - `ignore-case` (игнорировать регистр)\n-v - `invert` (вместо совпадения, исключать)\n-F - `fixed`, точное совпадение со строкой, не паттерн\n-n - `line num`, печатать номер строки")
		return
	}
	arg := parseFlags()
	fmt.Println(arg)
	strs := createArr(arg)
	strs2 := reArray(strs, arg)
	
	outputStrings(strs, arg)
	outputStrings(strs2, arg)
}