package main
import (
	// "fmt"
	"strings"
)

//CreateRunesArr создать масиив 2мерный рун вместо массива строк
func CreateRunesArr(s0 []string) [][]rune {
	ss := make([][]rune, 0, 0)
	for i := range s0 {
		s := []rune(s0[i])
		ss = append(ss, s)
	}
	return ss
}


func IsAnagram(s10 []rune, s20 []rune) bool {
	s10 = []rune(strings.ToLower(string(s10)))
	s1 := make([]rune, 0, 0)
	s1 = append(s1, s10...)
	s20 = []rune(strings.ToLower(string(s20)))
	s2 := make([]rune, 0, 0)

	s2 = append(s2, s20...)
	ok := 0
	for i := 0; i < len(s1); i++ {
		bin := false
		for j := 0; j < len(s2); j++ {
			if s1[i] == s2[j] {
				bin = true
				s2 = append(s2[:j], s2[j + 1:]...)
				break
			}
		}
		if bin {
			ok++
		}
	}
	if ok == len(s1) {
		return true
	} else {
		return false
	}
}

func Anagrams(s0 []string) map[string][]string {
	ss0 := CreateRunesArr(s0)
	ss := ss0
	m := map[string][]string{}
	for i := 0; i < len(ss); i++{
		cnt := 0
		for j := i; j < len(ss); j++ {
			if i != j && len(ss[i]) == len(ss[j]) {
				s1, s2 := ss[i], ss[j]
				if IsAnagram(s1, s2) {
					ss[i] = []rune(strings.ToLower(string(ss[i])))
					ss[j] = []rune(strings.ToLower(string(ss[j])))
					cnt++
					if cnt == 1 {
						m[string(ss[i])] = []string{string(ss[i]), string(ss[j])} 
					} else {
						m[string(ss[i])] = append(m[string(ss[i])], string(ss[j]))
					}
					ss = append(ss[:j], ss[j+1:]...)
					j--			
				}
			}
			

		}
	}
	return m
}