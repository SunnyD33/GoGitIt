package main

import (
	"errors"
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v2"

	Auth "GoGitIt/internal/auth"
	Utils "GoGitIt/pkg/utils"
)

type config struct {
	IsAuthorized bool `yaml:"authorized"`
}

func (conf *config) updateAuthState(state bool) {
	conf.IsAuthorized = state
}

// Function to save the config struct to a yaml file that will be housed
// in the $HOME directory
func saveConfig(c config, filename string) error {
	bytes, err := yaml.Marshal(c)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, bytes, 0644)
}

// Function to load the config struct from a yaml file and check if
// authorized is set to true and that there is a token in the .env file
func loadConfig(filename string) (config, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return config{}, err
	}

	var c config
	err = yaml.Unmarshal(bytes, &c)
	if err != nil {
		return config{}, err
	}

	return c, nil
}

//Function for checking the users auth state when certain commands are
//passed in to be parsed
func checkAuthStatus() config {
	tokenCheck := Auth.CheckAuthToken()
	homeDir, _ := os.UserHomeDir()
	c, err := loadConfig(homeDir + "/.ggiconfig.yaml")

	if err != nil {
		c.IsAuthorized = false
	} else if err == nil && tokenCheck {
		c.IsAuthorized = true
	} else {
		c.IsAuthorized = false
	}

	return c
}

//To be re-worked and potentially moved into it's own go file
func printHelpText() {
	helpText := "Help text should print"
	fmt.Println(helpText)
	fmt.Println("")
}

func parseArgs(args []string) (config, error) {
	c := config{}

	if len(args) < 1 {
		return c, errors.New("invalid number of arguments")
	}

	//Will print help text with tag options and how to use them
	if args[0] == "-h" || args[0] == "--help" {
		printHelpText()
		return c, nil
	}

	//Authorize user
	if args[0] == "-a" {
		authStatus := checkAuthStatus()
		fmt.Println("Checking for token...")

		if authStatus.IsAuthorized {
			fmt.Println("Currently authorized. You can use 'ggi -s' or 'ggi -status' to check your authorization status")
		} else if Auth.AuthToken == "" {
			authTokenCheker := Auth.GetAuthToken()
			if authTokenCheker == "" {
				fmt.Println("Authorization failed")
				fmt.Println("Unable to find token in your .env file. Please confirm that the GH_TOKEN variable is not empty.")
			} else {
				homeDir, err := os.UserHomeDir()
				if err != nil {
					return c, errors.New("could not find user home directory")
				}

				c.updateAuthState(true)
				saveConfig(c, homeDir+"/.ggiconfig.yaml")
				fmt.Println("Authorization successful!")
			}
		} else {
			fmt.Println("yaml file was not created and/or .env file does not exist or a token is not set.")
		}

		return c, nil
	}

	//Checks the current authorization status of the current user
	if args[0] == "--status" || args[0] == "-s" {
		authStatus := checkAuthStatus()

		if !authStatus.IsAuthorized {
			Auth.PrintUnauthorizedText()
		} else {
			Auth.PrintAuthorizedText()
		}
	}

	return c, nil
}

func runCmd(r io.Reader, w io.Writer, c config) error {
	//TODO: To use for some commands that will require reading of inputs
	fmt.Println("Command ran successfully") //Here to know that runCmd ran as expected
	return nil
}

func main() {
	loadErr := Utils.LoadEnv()
	if loadErr != nil {
		fmt.Println(loadErr)
		return
	}

	c, err := parseArgs(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		printHelpText()
		os.Exit(1)
	}

	err = runCmd(os.Stdin, os.Stdout, c)
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}
}