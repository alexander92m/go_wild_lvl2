package main

import (
	// "fmt"
	"testing"
	"os"
	"flag"
)

type tester struct {
	com 	string
	files 	string
}


// func TestMain(m *testing.M) {
// 	input := []tester{
// 		{
// 		"-nk2",
// 		"text",
// 		},{
// 		"-rnk2",
// 		"text",
// 		},
// 	}
	
// 	for _, testCase := range input {
// 		flag.Parse()
// 		os.Exit(m.Run())
// 		fmt.Println(testCase)
// 	}
// }
func TestFlags(T *testing.T) {
    // We manipuate the Args to set them up for the testcases
    // After this test we restore the initial args
    oldArgs := os.Args
    defer func() { os.Args = oldArgs }()
    cases := []struct {
        Args           []string
        ExpectedExit   int
        ExpectedOutput string
    }{
        {[]string{"-nk2", "text"}, 0, "22a y aug  5 6 7 a8 1 1 3\n1c 2 octobe 4 5 6 7 c8\n51 2 Jan 4 5 6 7 g8 1 1 1\n9 2 3 4 5 6 7 c8\n9 2 3 4 5 6 7 c8\na 2 June 4 5 6 7 c8\nb5 2 FEB 4 5 6 7 b8 1 1 2\n"},
    }
    for _, tc := range cases {
        // this call is required because otherwise flags panics,
        // if args are set between flag.Parse call
        flag.CommandLine = flag.NewFlagSet(flag.ExitOnError)
        // we need a value to set Args[0] to cause flag begins parsing at Args[1]
        os.Args = append([]string{tc.Name}, tc.Args...)
        actualExit := realMain()
        if tc.ExpectedExit != actualExit {
            T.Errorf("Wrong exit code for args: %v, expected: %v, got: %v",
                tc.Args, tc.ExpectedExit, actualExit)
        }
    }
}