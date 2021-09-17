package main

import (
	"fmt"
	"testing"
	"os"
)

type tester struct {
	com 	string
	files 	string
}

func Testmain(m *testing.M) {
	input := []tester{
		{
		"-nk2",
		"text",
		},{
		"-rnk2",
		"text",
		},
	}
	
	for _, testCase := range input {
		flag.Parse("text")
		os.Exit(m.Run())
		fmt.Println(testCase)
	}
}