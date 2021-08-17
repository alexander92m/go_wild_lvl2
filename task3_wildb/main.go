// чтение всех строк всех файлов в массив [][]rune

// вывод на экран в порядке возрастания или с учетом флагов
package main

import (
	"fmt"
	"os"
	"bufio"
	"io"
	"sort"
)

//struct of agrs(Flags, FileNames)
type Arguments struct {
	Flags		string
	FileNames	[]string
}

//normal sort
func NormalSort(strs []string) {
	fmt.Println("==========================\n")
	sort.Strings(strs)
	for i := range strs {
		fmt.Printf("|%v|\n", strs[i])
	}
}

//parse flags from args
func ParseFlags() {
	for i := 0; i < len(os.Args); i++ {
		
	}

}

//Create array of strings
func CreateArray() []string{
	strs := make([]string, 0, 0)
	filename1 := "1"
	file, errOpen := os.Open(filename1)
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
	for i := range strs {
		fmt.Printf("|%v|\n", strs[i])
	}
	return strs

}

//output strings
func Output_strings(strs []string) {
	NormalSort(strs)
}

//main
func main(){
	for i := 0; i < len(os.Args); i++ {
		fmt.Printf("i=%d, args i= %v, type= %T\n", i, os.Args[i], os.Args[i])
	}
	strs := CreateArray()

	Output_strings(strs)
	
}
