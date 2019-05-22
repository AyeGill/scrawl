package main

import (
	"bufio"
	"flag"
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

func getCommand(cmd string, cmdpath string) string {
	//Runs through each line of cmdpath.
	//If the first word of the line (ended by a space) is cmd,
	//return the rest of line, without the space.
	//If no matches, return cmd.

	if cmdpath == "" {
		return cmd
	}

	contents, err := ioutil.ReadFile(cmdpath)
	check(err)
	for _, line := range strings.Split(string(contents), "\n") {
		split := strings.SplitN(line, " ", 2)
		if len(split) == 0 {
			io.WriteString(os.Stderr, "Empty line in command file\n")
		} else if len(split) == 1 {
			io.WriteString(os.Stderr, "Line in command file without spaces\n")
		} else if split[0] == cmd { //Found a proper split, we're good to go.
			return split[1]
		}
	}
	return cmd
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

func doFile(file *os.File, cmdfile string) {

	scanner := bufio.NewScanner(file)

	inCmd := false //ghetto parsing.
	command := ""
	input := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 2 && line[:3] == "%-#" { //we are reading scrawl input. Continue.
			if inCmd { //Continue parsing command input
				if len(line) > 5 && line[3:6] == "---" { //Command is over, execute it.
					output := run(getCommand(command, cmdfile), strings.Join(input[:], "\n"))
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

func errmsg(msg string, v bool) {
	if v {
		io.WriteString(os.Stderr, msg)
	}

}

func main() {
	verbose := flag.Bool("v", false, "Print debugging information to stderr")
	cmdfile := flag.String("c", "", "Load a macro file - see README or source code for info")

	flag.Parse()

	if len(flag.Args()) == 0 {
		errmsg("No args\n", *verbose)
		doFile(os.Stdin, *cmdfile)
	} else {
		errmsg("Some args\n", *verbose)
		for _, path := range flag.Args() {
			errmsg(path, *verbose)
			file, err := os.Open(path)
			check(err)
			doFile(file, *cmdfile)
		}
	}
}
