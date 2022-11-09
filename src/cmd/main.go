package main

import (
	"errors"
	"fmt"
	"io"
	"os"

	Auth "GoGitIt/internal/auth"
	Utils "GoGitIt/pkg/utils"
)

type config struct {
	printHelpText bool
	isAuthorized  bool
}

func (conf *config) updateAuthState(state bool) {
	conf.isAuthorized = state
}

var helpText = "Help text should print"

func printHelpText(w io.Writer) {
	fmt.Fprint(w, helpText)
	fmt.Println("")
}

func parseArgs(args []string) (config, error) {
	c := config{}

	if len(args) < 1 {
		return c, errors.New("invalid number of arguments")
	}

	//Will print help text with tag options and how to use them
	if args[0] == "-h" || args[0] == "--help" {
		c.printHelpText = true
		return c, nil
	}

	//Authorize user
	if args[0] == "-a" {
		fmt.Println("Checking for token...")

		//A is variable from auth.go to give access to the Auth struct
		if Auth.A.AuthToken == "" {
			authStateCheker := Auth.SetAuthState()
			if authStateCheker == "" {
				fmt.Println("Unable to find token in your .env file. Please confirm that the GH_TOKEN variable is not empty.")
			} else {
				fmt.Println("Authorization successful!")
				c.updateAuthState(true)
			}
		} else {
			fmt.Println("Currently authorized. You can use 'ggi -s' or 'ggi -status' to check yur authorization status")
		}
		return c, nil
	}

	//Checks that current authorization status of the current user
	if args[0] == "--status" || args[0] == "-s" {

		c.isAuthorized = Auth.CheckAuthState()

		if c.isAuthorized {
			Auth.PrintAuthorizedText()
		} else {
			Auth.PrintUnauthorizedText()
		}

		return c, nil
	}

	return c, nil
}

func runCmd(r io.Reader, w io.Writer, c config) error {
	if c.printHelpText {
		printHelpText(w)
		return nil
	}

	//TODO: To remove on final build
	fmt.Println("Command ran successfully") //Here to know that runCmd ran as expected
	return nil
}

func main() {
	Utils.LoadEnv()

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
