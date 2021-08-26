package main

import (
	"testing"
	"fmt"
	"strings"
)

type tester struct {
	in []string
	out map[string][]string
}
func TestAnagramgs(t *testing.T) {
	tes := []tester{
		tester{
			in: []string{"стол", "слот", "стул"},
			out: map[string][]string{"стол": []string{"стол", "слот"}},
		},
		// tester{
		// 	[]string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"},
		// 	map[string][]string{"пятак": []string{"пятак", "пятка", "тяпка"}, "листок": []string{"листок", "слиток", "столик"}},
		// },
		
	}
	// fmt.Println(tes[0].out["стол"])
	for i := range tes {
		got := Anagrams(tes[i].in)
		for j := range got {
			// fmt.Println(j, i, " | ", got[j], " | ", tes[i].out[j])
			if strings.Join(got[j], "") != strings.Join(tes[i].out[j], "") {
				fmt.Println(i)
				t.Errorf("in %v, want %v, got %v", tes[i].in, tes[i].out, got)
			}
		}
		
	}
	fmt.Println("end test")
}