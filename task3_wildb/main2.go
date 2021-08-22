// чтение всех строк всех файлов в массив [][]rune

// вывод на экран в порядке возрастания или с учетом флагов
package main

import (
	"fmt"
	"os"
	"bufio"
	"io"
	"strings"
	// "strconv"
)

//struct of agrs(Flags, Files)
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



//parse flags from args
func ParseFlags() Arguments {
	var arg Arguments
	
	for i := 1; i < len(os.Args); i++ {
		
		if strings.HasPrefix(os.Args[i], "-") {
			
			if strings.Contains(os.Args[i], "k") {
				arg.k = int(os.Args[i][strings.Index(os.Args[i], "k") + 1])
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

//Create array of strings
func CreateArray(arg Arguments) [][]string{
	strs := make([][]string, 0, 0)
	for i := range arg.Files {
		file,  errOpen := os.Open(arg.Files[i])
		if errOpen != nil {
			fmt.Println("Error of open", arg.Files[i])
			os.Exit(0)
		}
		rd := bufio.NewReader(file)
		for j := 0; j < 2; {
			s, errRead := rd.ReadString('\n')
			if errRead != nil && errRead != io.EOF {
				fmt.Println("Error of read string")
			} else if errRead == io.EOF {
				j = 2
			} else {
				s = s[0:len(s) - 1]
			}
			ss := make([]string, 0, 0)
			ss = append(ss, arg.Files[i], s)
			strs = append(strs, ss)
		}
		file.Close()
	}
	return strs
}



//convert []string to [][]rune
func stringToRunes(strs []string) [][]rune {
	// strsR := make([][]rune, 0, 0)
	var strsR [][]rune
	for i := range strs {
		strsR = append(strsR, []rune(strs[i]))
	}
	return strsR
}

//convert [][]rune to []string
func runesToString(strsR [][]rune) []string {
	strs := make([]string, 0, 0)
	for i := range strsR {
		strs = append(strs, string(strsR[i]))
	}
	return strs
}

//reverse sorting


//main
func main(){
	arg := ParseFlags()
	strs := CreateArray(arg)
	fmt.Printf("%v, %T \n %v, %T\n", arg, arg, strs,  strs)
	outputStrings(strs, arg)
	
}