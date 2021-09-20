package main

import (
	"fmt"
	"os"
	"bufio"
	"io"
	"strings"
)

//Arguments struct of agrs(Flags, Files)
type Arguments struct {
	k			int
	n			bool
	u			bool
	r			bool
	M			bool
	b			bool
	c			bool
	h			bool

	Files	[]string
}


func isDigit(c rune) bool {
	if c >= '0' && c <= '9' {
		return true
	}
	return false
}
//ParseFlags - parsing flags from args
func ParseFlags() Arguments {
	var arg Arguments
	
	arg.k = 1
	for i := 1; i < len(os.Args); i++ {
		if strings.HasPrefix(os.Args[i], "-") {
			if strings.Contains(os.Args[i], "k") &&  isDigit(rune(os.Args[i][strings.Index(os.Args[i], "k") + 1])) {
				arg.k = 0
				
				for j := strings.Index(os.Args[i], "k") + 1; j < len(os.Args[i]) && isDigit(rune(os.Args[i][j])); j++ {
					arg.k = arg.k * 10 + int(os.Args[i][j]) - '0'
				}
			}
			if strings.Contains(os.Args[i], "n"){
				arg.n = true
			}
			if strings.Contains(os.Args[i], "r"){
				arg.r = true
			}
			if strings.Contains(os.Args[i], "u"){
				arg.u = true
			}
			if strings.Contains(os.Args[i], "M"){
				arg.M = true
			}
			if strings.Contains(os.Args[i], "b"){
				arg.b = true
			}
			if strings.Contains(os.Args[i], "c"){
				arg.c = true
			}
			if strings.Contains(os.Args[i], "h"){
				arg.h = true
			}
		} else {
			arg.Files = append(arg.Files, os.Args[i])
		}
		
	}
	return arg
}

//CreateArray of strings from files
func CreateArray(arg Arguments) [][]string{
	strs := make([][]string, 0, 0)
	if len(arg.Files) == 0 {
		sc := bufio.NewScanner(os.Stdin)
		for sc.Scan() {
			s := sc.Text()
			var s2 []string
			// s = s[0:len(s) - 1]
			s2 = strings.Split(s, " ")
			ss := make([]string, 0, 0)
			ss = append(ss, "basic input")
			for k := range s2 {
				ss = append(ss, s2[k])
			}
			strs = append(strs, ss)
		}
	} else {
		for i := range arg.Files {
			file,  errOpen := os.Open(arg.Files[i])
			if errOpen != nil {
				fmt.Println("Error of open", arg.Files[i])
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
				ss = append(ss, arg.Files[i])
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

func normalSort(strs [][]string, arg Arguments) [][]string {
	for i := len(strs) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			ind1 := 1
			ind2 := 1
			if arg.k < len(strs[j]) {
				ind1 = arg.k
			}
			if arg.k < len(strs[j + 1]) {
				ind2 = arg.k
			}
			if strs[j][ind1] > strs[j + 1][ind2] {
				strs[j], strs[j + 1] = strs[j + 1], strs[j]
			}
		}
	}
	return strs
}

func	ftOk(c rune) int {
	if c == 9 || c == 10 || c == 11 {
		return 2
	} else if c == 12 || c == 13 || c == 32 {
		return 2
	} else if isDigit(c) {
		return 1
	} else if c == '-' || c == '+' {
		return 3
	} else {
		return 5
	}
}

func	ftMult(str []rune, i int, num int64, sign int64) int64 {
	for ;i < len(str) && str[i] >= '0' && str[i] <= '9'; i++ {
		if (num * sign) < -2147483648 {
			return 0
		} else if (num * sign) > 2147483647 {
			return -1
		} else {
			num = num * 10 + int64(str[i]) - '0'
		}
		if (i + 1 < len(str)) && !(str[i + 1] >= '0' && str[i + 1] <= '9') {
			break
		}
	}
	return num * sign
}

func	ftAtoi(str0 string) (int, string) {
	var i int
	var num int64
	var sign int64
	str := []rune(str0)
	sign = 1

	if !isDigit(str[0]) && str[0] != '-' {
		return 0, "invalid string"
	}

	for str[i] != 0 && (ftOk(str[i]) < 4) {
		if (ftOk(str[i]) == 3) && (ftOk(str[i + 1]) == 1) {
			if str[i] == '-'{
				sign = -1
			}
		}
		if ftOk(str[i]) == 3 && ftOk(str[i + 1]) != 1 {
			break
		}
		if isDigit(str[i]) {
			num = ftMult(str, i, num, sign)
			break
		}
		if ftOk(str[i]) == 2 && ftOk(str[i + 1]) == 3 {
			if ftOk(str[i + 2]) != 1 {
				break
			}
		}
		i++
	}
	return int(num), "";
}

func nSort(strs [][]string, arg Arguments) [][]string {
	for i := len(strs) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			ind1 := 1
			ind2 := 1
			if arg.k < len(strs[j]) {
				ind1 = arg.k
			}
			if arg.k < len(strs[j + 1]) {
				ind2 = arg.k
			}
			num, err := ftAtoi(strs[j][ind1])
			num2, err2 := ftAtoi(strs[j + 1][ind2])
			if err == "" && err2 == "" {
				if num > num2 {
					strs[j], strs[j + 1] = strs[j + 1], strs[j]
				}
			} else if err == "" && err2 != "" {
				strs[j], strs[j + 1] = strs[j + 1], strs[j]
			} else if err != "" && err2 != "" {
				if strs[j][ind1] > strs[j + 1][ind2] {
					strs[j], strs[j + 1] = strs[j + 1], strs[j]
				}
			}
		}
	}
	return strs
}

func rSort(strs [][]string) [][]string {
	for i := 0; i < len(strs) / 2; i++ {
		strs[i], strs[len(strs) - 1 - i] = strs[len(strs) - 1 - i], strs[i]
	}
	return strs
}

func equalSlicesStr(s1 []string, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := 0; i < len(s1) && i < len(s2); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

func uSort(s [][]string) [][]string {
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			s1 := s[i][1: ]
			s2 := s[j][1:]
			if equalSlicesStr(s1, s2) {
				s = append(s[ :j], s[j + 1: ]...)
			}
		}
	}
	return s
}

func mounthToDigit(s string) int {
	if len(s) == 3 {
		s = strings.ToLower(s)
	} else if len(s) > 3 {
		s = strings.ToLower(s[0:3])
	}
	fmt.Println(s)
	switch s {
		case "jan": return 1
		case "feb": return 2
		case "mar": return 3
		case "apr": return 4
		case "may": return 5
		case "jun": return 6
		case "jul": return 7
		case "aug": return 8
		case "sep": return 9
		case "oct": return 10
		case "nov": return 11
		case "dec": return 12
		default: return 0
	}
}

//MSort is sorting on mounths
func MSort(strs [][]string, arg Arguments) [][]string {
	for i := len(strs) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			ind1 := 1
			ind2 := 1
			if arg.k < len(strs[j]) {
				ind1 = arg.k
			}
			if arg.k < len(strs[j + 1]) {
				ind2 = arg.k
			}
			if mounthToDigit(strs[j][ind1]) > mounthToDigit(strs[j + 1][ind2]) {
				strs[j], strs[j + 1] = strs[j + 1], strs[j]
			}
		}
		
	}
	return strs
}

func outputStrings(strs0 [][]string, arg Arguments) {
	strs := strs0
	if arg.n {
		strs = nSort(strs, arg)
	} else {
		strs = normalSort(strs, arg)
	}
	if arg.r {
		strs = rSort(strs)
	}
	if arg.u {
		strs = uSort(strs)
	}
	if arg.M {
		strs = MSort(strs, arg)
	}
	for i := range strs {
		for j := range strs[i] {
			if j != 0 {
				fmt.Print(strs[i][j], " ")
			}
		}
		fmt.Println()
	}
}
	
//main
func main() {
	arg := ParseFlags()
	strs := CreateArray(arg)
	outputStrings(strs, arg)
	
}