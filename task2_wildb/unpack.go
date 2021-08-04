package main

import (
	"fmt"
)

func isDigit(c rune) bool {
	if c >= '0' && c <= '9' {
		return true
	}
	return false
}
// "a4bc2d5e" => "aaaabccddddde"
func Unpack(s string) string {
	s2 := []rune(s)
	temp := 0
	s3 := make([]rune, 0)
	for i := range s2 {
		if isDigit(s2[i]) {
			temp = int(s2[i] - rune(48))
			for j := 0; j < temp; j++ {
				s3 = append(s3, s2[i - 1])
			}
		} else {
			s3 = append(s3, s2[i])
		}
		
	}
	fmt.Println(string(s3))
	return string(s2)
}