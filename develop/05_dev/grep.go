//subj.txt
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type argFlag struct {
	A     int
	B     int
	C     int
	c     bool
	i     bool
	v     bool
	F     bool
	n     bool
	word  string
	files []string
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
				if len(os.Args[i]) < 3 || !isDigit(rune(os.Args[i][ind+1])) {
					err := fmt.Errorf("UNKNOW FLAG")
					fmt.Println(err)

				} else {
					flags.A, _ = strconv.Atoi(os.Args[i][2:])
				}
				flags.i = true
			}
			if strings.Contains(os.Args[i], "B") {
				ind := strings.Index(os.Args[i], "B")
				if len(os.Args[i]) < 3 || !isDigit(rune(os.Args[i][ind+1])) {
					err := fmt.Errorf("UNKNOW FLAG")
					fmt.Println(err)

				} else {
					flags.B, _ = strconv.Atoi(os.Args[i][2:])
				}
			}
			if strings.Contains(os.Args[i], "C") {
				ind := strings.Index(os.Args[i], "C")
				if len(os.Args[i]) < 3 || !isDigit(rune(os.Args[i][ind+1])) {
					err := fmt.Errorf("UNKNOW FLAG")
					fmt.Println(err)

				} else {
					flags.C, _ = strconv.Atoi(os.Args[i][2:])
					flags.A = flags.C
					flags.B = flags.C
				}
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
func createArr(arg argFlag) [][]string {
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
			file, errOpen := os.Open(arg.files[i])
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
					s = s[0 : len(s)-1]
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
	ss := strs
	for i := range ss {
		tmp := make([]string, 0, 0)
		tmp = append(tmp, ss[i][1:]...)
		ss[i] = append(ss[i][:1], strconv.Itoa(i + 1))
		ss[i] = append(ss[i], tmp...)
	}
	return ss
}

func arrContain(ss []string, arg argFlag) bool {
	match := false
	if arg.F {
		s := ""
		if len(ss) > 2 {
			s = strings.Join(ss[2:], "")
		}
		match = (s == arg.word)
	} else {
		for i := 2; i < len(ss); i++ {
			if strings.Contains(ss[i], arg.word) {
				match = true
			}
		}
	}
	
	return match
}

func reArray(ss0 [][]string, arg argFlag) [][]string {
	ss := make([][]string, 0, 0)
	ss2 := make([][]string, 0, 0)
	var status string = "clean"
	ss = append(ss, ss0...)

	
	for i := range ss {
		match := arrContain(ss[i], arg)
		if match {
			if status == "minor+" {
				ss2 = append(ss2, []string{"..", "--"})
				status = ".."
			}
			ss2 = append(ss2, ss[i])
			status = "major"

			match2 := false

			for k := 1; !match2 && k <= arg.B && i-k >= 0; k++ {
				for l := 2; l < len(ss[i-k]); l++ {
					if strings.Contains(ss[i-k][l], arg.word) {

						match2 = true
						break
					}
				}
				if !match2 {
					s := [][]string{}
					s = append(s, ss2[len(ss2)-k:]...)
					ss2 = append(ss2[0:len(ss2)-k], ss[i-k])
					ss2 = append(ss2, s...)
					if k == arg.A {
						status = "minor+"
					} else {
						status = "minor-"
					}
				}
			}
			match2 = false
			for k := 1; !match2 && k <= arg.A && k+i < len(ss); k++ {

				for l := 2; l < len(ss[i+k]); l++ {
					if strings.Contains(ss[i+k][l], arg.word) {

						match2 = true
						break
					}
				}
				if !match2 {
					ss2 = append(ss2, ss[i+k])

					if k == arg.A {
						status = "minor+"
					} else {
						status = "minor-"
					}
				}
			}

		}

	}
	return ss2
}

func reArrayV(ss0 [][]string, arg argFlag) [][]string {
	ss := make([][]string, 0, 0)

	ss = append(ss, ss0...)
	
	for i := 0; i < len(ss); i++ {
		match := arrContain(ss[i], arg)
		if match {
			if i == len(ss)-1 {
				ss = ss[:i]
			} else {
				ss = append(ss[:i], ss[i+1:]...)
			}

		}
	}
	return ss
}

func outputStrings(strs [][]string, arg argFlag) {
	for i := range strs {
		if arg.n {
			fmt.Printf("%v:%v\n", strs[i][1], strings.Join(strs[i][2:], ""))
		} else {
			for j := range strs[i] {
				if j > 1 {
					fmt.Print(strs[i][j], " ")
				}
				
			}
			fmt.Println()
		}
	}
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("некорретные входные данные\n КАК?: grep [-civFn] [-A num] [-B num] [-C[num] [filter key] [file ...]\n-A - `after` печатать +N строк после совпадения\n-B - `before` печатать +N строк до совпадения\n-C - `context` (A+B) печатать ±N строк вокруг совпадения\n-c - `count` (количество строк)\n-i - `ignore-case` (игнорировать регистр)\n-v - `invert` (вместо совпадения, исключать)\n-F - `fixed`, точное совпадение со строкой, не паттерн\n-n - `line num`, печатать номер строки")
		return
	}
	arg := parseFlags()
	strs := createArr(arg)
	strs2 := [][]string{}
	if arg.i {
		arg.word = strings.ToLower(arg.word)
		for i := range strs {
			for j := 1; j < len(strs[i]); j++ {
				strs[i][j] = strings.ToLower(strs[i][j])
			}
		}
	}
	if !arg.v {
		strs2 = reArray(strs, arg)
	} else {
		strs2 = reArrayV(strs, arg)
	}
	if arg.c {
		fmt.Println(len(strs2))
	} else {
		outputStrings(strs2, arg)
	}
}
