package main

import (
	"fmt"
)
//проверка на число
func isDigit(c rune) bool {
	if c >= '0' && c <= '9' {
		return true
	}
	return false
}
//проверка на слэш
func isSlash(c rune) bool {
	if c == '\\' {
		return true
	}
	return false
}
//проверка на ошибку
func checker(s string) error{
	return nil
}

// "a4bc2d5e" => "aaaabccddddde"
func Unpack(s string) (string, error) {
	err := checker(s)
	if err != nil {
		return s, nil
	}
	s2 := []rune(s)
	temp := 0
	s3 := make([]rune, 0)
	for i := range s2 {
		
		if isSlash(s2[i]) {
			
			if i + 2 < len(s2) {
				if isDigit(s2[i + 2]) {
					fmt.Println("ch")
					temp = int(s2[i + 2] - rune(48))
					for j := 0; j < temp; j++ {
						s3 = append(s3, s2[i + 1])
						// fmt.Println("slash")						
					}
				}
			}
			
		} else if isDigit(s2[i]) {
			// fmt.Println("digit")
			temp = int(s2[i] - rune(48))
			for j := 0; j < temp; j++ {

				s3 = append(s3, s2[i - 1])
			}
		} else {
			// fmt.Println("other")
			s3 = append(s3, s2[i])
		}
		
	}
	return string(s3), nil
}