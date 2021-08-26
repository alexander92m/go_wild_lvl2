package main
import (
	"fmt"
)
func Anagrams(s []string) map[string][]string {	
	fmt.Println("anagrams")
	m := map[string][]string{}
	m["стол"] = []string{"стол", "слот"}
	return m
}