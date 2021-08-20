// чтение всех строк всех файлов в массив [][]rune

// вывод на экран в порядке возрастания или с учетом флагов
package main

import (
	"fmt"
	"os"
	"bufio"
	"io"
	"sort"
	"strings"
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
func CreateArray(arg Arguments) []string{
	strs := make([]string, 0, 0)
	for j := range arg.Files {
		file, errOpen := os.Open(arg.Files[j])
		if errOpen != nil {
			fmt.Println("Error of open")
		}
		rd := bufio.NewReader(file)
		for i := 0; i < 2; {
			s, errReadString := rd.ReadString('\n')
			if errReadString != nil && errReadString != io.EOF {
				fmt.Println("Error of readString")
			} else if errReadString == io.EOF {
				i = 2
			} else {
				s = s[0:len(s) - 1]
			}
			strs = append(strs, s)
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
func rSort(strs []string) []string {
	strsR := stringToRunes(strs)
	for i := 0; i < len(strsR) / 2; i++ {
		strsR[i], strsR[len(strsR) - 1 - i] = strsR[len(strsR) - 1 - i], strsR[i]
	}
	strs = runesToString(strsR)
	return strs
}

//normal sort
func NormalSort(strs []string) {
	sort.Strings(strs)
}

//undouble sorting
func uSort(strs []string) []string {
	for i := 0; i < len(strs); i++ {
		for j := i + 1; j < len(strs); j++ {
			if i < j && strs[i] == strs[j] {
				strs = append(strs[ :j], strs[j + 1: ]...)
			}
		}
	}
	return strs
}
//output strings
func Output_strings(strs []string, arg Arguments) {

	NormalSort(strs)
	if arg.r {
		strs = rSort(strs)
	}
	if arg.u {
		strs = uSort(strs)
	}
	for i := range strs {
		fmt.Printf("i=%d, |%v|\n", i, strs[i])
	}
}

//main
func main(){
	arg := ParseFlags()

	fmt.Printf("%v, %T\n", arg, arg)
	strs := CreateArray(arg)
	Output_strings(strs, arg)
	// fmt.Println(strs)
	
}