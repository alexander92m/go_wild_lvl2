package main

import (
	"fmt"
	"testing"
	"os"
    "os/exec"
)

type tester struct {
	com 	string
	files 	string
}


func TestMain(m *testing.M) {
	input := []tester{
		{
            com:"-nk1",
            files:"text",
		},{
            com:"-nk1",
		    files:"text",
		},
	}
	
	for _, testCase := range input {
        out, err := exec.Command("go", "run", "sort.go", testCase.com, testCase.files).CombinedOutput()
		if err != nil {
			fmt.Fprintf(os.Stderr, "building testgo failed: %v\n%s", err, out)
			os.Exit(2)
		}
		fmt.Println(string(out))
        out2, err2 := exec.Command("sort", testCase.com, testCase.files).CombinedOutput()
		if err2 != nil {
			fmt.Fprintf(os.Stderr, "building testgo failed: %v\n%s", err2, out2)
			os.Exit(2)
		}
		fmt.Println(string(out2))
        for i, char := range out {
            if (char != out2[i]) {
                fmt.Fprintf(os.Stderr, "РАЗЛИЧАЕТСЯ")
                os.Exit(2)
            }
        }
	}
}
