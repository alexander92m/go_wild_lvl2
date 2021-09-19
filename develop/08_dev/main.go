package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
)

func startClient(addr string) error {
	var conn net.Conn
	var err error
	conn, err = net.Dial("tcp", addr)
	if err != nil {
		conn, err = net.Dial("udp", addr)
		if err != nil {
			fmt.Printf("Can't connect to server: %s\n", err)
			return err
			} else {
				fmt.Println("connection established")
			}
	}
	_, err = io.Copy(conn, os.Stdin)
	if err != nil {
		fmt.Printf("Connection error: %s\n", err)
		return err
	} else {
		fmt.Println("connection sucsessfuly over")
	}
	return nil
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		dir, _ := os.Getwd()
		fmt.Printf("%s$ > ", dir)
		input, err := reader.ReadString('\n')
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
		}
		if err = execInput(input); err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
		}
	}
}

var ErrNoPath = errors.New("path required")

func execInput(input string) error {
	var cmd *exec.Cmd

	input = strings.TrimSuffix(input, "\n")

	args := strings.Split(input, " ")

	switch args[0] {
	case "cd":

		if len(args) < 2 {
			return ErrNoPath
		}

		return os.Chdir(args[1])
	case "netcat":
		return startClient(fmt.Sprintf("%s:%s", args[1], args[2]))

	case "fork":
		res, err := forkCommand(args)
		fmt.Println(res)
		if err != nil {
			return err
		}

	case "exit":
		os.Exit(0)
	}

	if !strings.Contains(input, "|") {
		cmd = exec.Command(args[0], args[1:]...)
	} else {
		cmd = exec.Command("bash", "-c", input)
	}
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}

func forkCommand(args []string) (string, error) {
	if len(args) < 2 {
		return "", errors.New("shell: fork command need amount of processes")
	}
	fork, err := strconv.Atoi(args[1])
	if err != nil {
		return "", err
	}
	children := []int{}
	var builder strings.Builder
	pid := os.Getpid()
	ppid := os.Getppid()
	builder.WriteString(
		fmt.Sprintf("pid: %d, ppid: %d, forks: %d\n", pid, ppid, fork),
	)
	if _, isChild := os.LookupEnv("CHILD_ID"); !isChild {
		for i := 0; i < fork; i++ {
			args := append(os.Args, fmt.Sprintf("#child_%d_of_%d", i, os.Getpid()))
			childENV := []string{
				fmt.Sprintf("CHILD_ID=%d", i),
			}
			pwd, err := os.Getwd()
			if err != nil {
				return "", err
			}
			childPID, _ := syscall.ForkExec(args[0], args, &syscall.ProcAttr{
				Dir: pwd,
				Env: append(os.Environ(), childENV...),
				Sys: &syscall.SysProcAttr{
					Setsid: true,
				},
				Files: []uintptr{0, 1, 2}, // print message to the same pty
			})
			builder.WriteString(
				fmt.Sprintf("parent %d fork %d\n", pid, childPID),
			)
			if childPID != 0 {
				children = append(children, childPID)
			}
		}
		// print children
		builder.WriteString(
			fmt.Sprintf("parent: PID=%d children=%v", pid, children),
		)
		if len(children) == 0 && fork != 0 {
			return "", errors.New("shell: fork no child avaliable, exit")
		}

		// set env
		for _, childID := range children {
			if c := os.Getenv("CHILDREN"); c != "" {
				os.Setenv("CHILDREN", fmt.Sprintf("%s,%d", c, childID))
			} else {
				os.Setenv("CHILDREN", fmt.Sprintf("%d", childID))
			}
		}
	}
	return "", nil
}
