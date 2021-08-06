package main

import (
	"fmt"
	"errors"
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
	s2 := []rune(s)
	for i := range s2 {
		if i == 0 && isDigit(s2[i]) {
			error := errors.New("некорректная строка")
			return error
		}
	}
	return nil
}

//подсчет числа символов и число символов для отображения числа символов
func countTimes(s2 []rune, i int) (int, int) {
	cnt := 0
	j := 0
	for j = i; j < len(s2); j++ {
		if !isDigit(s2[j]) {
			break
		}
		cnt = cnt * 10 + int(s2[i] - rune(48))
	}
	return cnt, j - i
}

// "a4bc2d5e" => "aaaabccddddde"
func Unpack(s string) (string, error) {
	fmt.Println("")
	err := checker(s)
	if err != nil {
		return s, err
	}

	s2 := []rune(s)
	cnt := 0
	s3 := make([]rune, 0)
	for i := 0; i < len(s2); i++ {
		cnt = 0
		if isSlash(s2[i]) {	
				s3 = append(s3, s2[i + 1])
				if i + 2 < len(s2) {
					if isDigit(s2[i + 2]) {
						tmp := 0
						cnt, tmp = countTimes(s2, i + 2)
						fmt.Println(tmp)
						for j := 1; j < cnt; j++ {
							s3 = append(s3, s2[i + 1])					
						}
						i = i + tmp
					}
				}
				i++
		} else if isDigit(s2[i]) {
			tmp := 0
			cnt, tmp = countTimes(s2, i)
			fmt.Println(tmp)
			for j := 1; j < cnt; j++ {
				s3 = append(s3, s2[i - 1])
			}
			i = i + tmp - 1
		} else {
			s3 = append(s3, s2[i])
		}
		
	}
	return string(s3), nil
}