package main

import (
	"errors"
	"fmt"
	"io"
	"os"

	"../pkg/utils"
)

type config struct {
	printHelpText bool
}

var helpText = "Help text should print"

func printHelpText(w io.Writer) {
	fmt.Fprintf(w, helpText)
	fmt.Println("")
}

func parseArgs(args []string) (config, error) {
	c := config{}

	if len(args) != 1 {
		return c, errors.New("Invalid number of arguments\n")
	}

	if args[0] == "-h" || args[0] == "--help" {
		c.printHelpText = true
		return c, nil
	}

	return c, nil
}

func runCmd(r io.Reader, w io.Writer, c config) error {
	if c.printHelpText {
		printHelpText(w)
		return nil
	}

	fmt.Println("Command ran successfully")
	return nil
}

func main() {
	c, err := parseArgs(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		printHelpText(os.Stdout)
		os.Exit(1)
	}

	err = runCmd(os.Stdin, os.Stdout, c)
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}
}
