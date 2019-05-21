package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func run(command string, input string) string {
	cmd := exec.Command("bash", "-c", command)
	stdin, err := cmd.StdinPipe()
	check(err)
	stdout, err := cmd.StdoutPipe()
	check(err)
	check(cmd.Start())
	_, err = io.WriteString(stdin, input)
	stdin.Close()

	check(err)
	output, err := ioutil.ReadAll(stdout)
	stdout.Close()
	check(err)
	return string(output)
}

func main() {
	// Usage: scrawl filename.tex
	//Currently output just goes to standard output
	//No real support for multiple files, just concatenates output to stdout.
	if len(os.Args) == 0 {
		io.WriteString(os.Stderr, "no args")
		doFile(os.Stdin)
	} else {
		io.WriteString(os.Stderr, "Some args")
		for _, path := range os.Args[1:] {
			io.WriteString(os.Stderr, path)
			file, err := os.Open(path)
			check(err)
			doFile(file)
		}
	}
}
func doFile(file *os.File) {

	scanner := bufio.NewScanner(file)

	inCmd := false //ghetto parsing.
	command := ""
	input := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 2 && line[:3] == "%-#" { //we are reading scrawl input. Continue.
			if inCmd { //Continue parsing command input
				if len(line) > 5 && line[3:6] == "---" { //Command is over, execute it.
					output := run(command, strings.Join(input[:], "\n"))
					inCmd = false
					command = ""
					input = make([]string, 0)
					fmt.Println(output)
				} else { //Continue accumulating command input
					input = append(input, line[3:])
				}
			} else { //We must be starting a new command
				inCmd = true
				command = line[3:]
			}
		} else { //Ordinary line, just print it.
			fmt.Println(line)
		}
	}
}

// Example usage:
// \section{Foo}
// Here is some math $\sum f(n)$. Consider this diagram:
// %-#string-diag-builder
// %-#here we put some stuff
// %-#and some more stuff
// %-#-----

// Three or more dashes ends the command.
