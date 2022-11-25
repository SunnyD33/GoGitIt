package main

import (
	"errors"
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v2"

	Auth "GoGitIt/internal/auth"
	Repos "GoGitIt/internal/repos"
	Utils "GoGitIt/pkg/utils"
)

type Config struct {
	IsAuthorized bool   `yaml:"authorized"`
	EnvLocation  string `yaml:"envLocation"`
}

func (conf *Config) updateAuthState(state bool) {
	conf.IsAuthorized = state
}

func (conf *Config) updateEnvLocation(filepath string) {
	conf.EnvLocation = filepath
}

// Function to save the config struct to a yml file that will be housed
// in the $HOME directory
func saveConfig(c Config, filename string) error {
	bytes, err := yaml.Marshal(c)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, bytes, 0644)
}

// Function to load the config struct from a yml file and check if
// authorized is set to true and that there is a token in the .env file
func loadConfig(filename string) (Config, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return Config{}, err
	}

	var c Config
	err = yaml.Unmarshal(bytes, &c)
	if err != nil {
		return Config{}, err
	}

	return c, nil
}

// Function to check if the env filepath is set
func checkEnvLocation() string {
	homeDir, _ := os.UserHomeDir()
	result, err := loadConfig(homeDir + "/.ggiconfig.yml")

	if err != nil {
		fmt.Println("Error opening yml config file. Run 'touch .ggiconfig.yml' in your home directory")
		os.Exit(1)
	}

	return result.EnvLocation
}

// Function for checking the users auth state when certain commands are
// passed in to be parsed
func checkAuthStatus() Config {
	tokenCheck := Auth.CheckAuthToken()
	homeDir, _ := os.UserHomeDir()
	result, err := loadConfig(homeDir + "/.ggiconfig.yml")

	if err != nil {
		fmt.Println(err)
	} else if err == nil && tokenCheck {
		result.updateAuthState(true)
		result.updateEnvLocation(result.EnvLocation)
		saveConfig(result, homeDir+"/.ggiconfig.yml")
	} else {
		result.updateAuthState(false)
		result.updateEnvLocation(result.EnvLocation)
		saveConfig(result, homeDir+"/.ggiconfig.yml")
	}

	return result
}

// To be re-worked and potentially moved into it's own go file
func printHelpText() {
	helpText := "Help text should print"
	fmt.Println(helpText)
	fmt.Println("")
}

func parseArgs(args []string) (Config, error) {
	c := Config{}

	//Check for .env file when commands are run
	loadErr := Utils.LoadEnv(checkEnvLocation())

	if loadErr != nil {
		fmt.Println(loadErr)
		homeDir, err := os.UserHomeDir()

		if err != nil {
			return c, errors.New("could not find user home directory")
		}
		envFileLocation, _ := Utils.SetEnv(os.Stdin, os.Stdout)

		result, _ := loadConfig(homeDir + "/.ggiconfig.yml")
		if err != nil {
			fmt.Println(err)
		}

		result.updateEnvLocation(envFileLocation)
		result.updateAuthState(false)

		saveConfig(result, homeDir+"/.ggiconfig.yml")

		return c, nil
	}

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
		fmt.Println("Checking for token...")
		homeDir, _ := os.UserHomeDir()
		authToken := Auth.CheckAuthToken()

		//TODO: Add logic to not authorize user if already authorized
		result, _ := loadConfig(homeDir + "/.ggiconfig.yml")

		if result.IsAuthorized {
			fmt.Println("You are already authorized! Cancelling operation!")
			return result, nil
		}

		if !authToken {
			fmt.Println("Authorization failed")
			fmt.Println("Unable to find token in your .env file. Please confirm that the GH_TOKEN variable is not empty.")
		} else {
			homeDir, err := os.UserHomeDir()
			if err != nil {
				return c, errors.New("could not find user home directory")
			}

			result, _ := loadConfig(homeDir + "/.ggiconfig.yml")

			c.updateAuthState(true)
			c.updateEnvLocation(result.EnvLocation)
			saveConfig(c, homeDir+"/.ggiconfig.yml")
			fmt.Println("Authorization successful!")
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

	//Allow user to set a custom .env file location
	if args[0] == "--setenv" {
		homeDir, err := os.UserHomeDir()

		if err != nil {
			return c, errors.New("could not find user home directory")
		}
		envFileLocation, err := Utils.SetEnv(os.Stdin, os.Stdout)

		if err != nil {
			return c, err
		}

		result, err := loadConfig(homeDir + "/.ggiconfig.yml")
		if err != nil {
			return result, err
		}

		c.updateEnvLocation(envFileLocation)
		c.updateAuthState(result.IsAuthorized)
		saveConfig(c, homeDir+"/.ggiconfig.yml")
	}

	//Allow for users to pull a list of repos for a user
	if args[0] == "-r" {
		Repos.GetRepos()
	}

	return c, nil
}

func runCmd(r io.Reader, w io.Writer, c Config) error {
	//TODO: To use for some commands that will require reading of inputs
	fmt.Println("Command ran successfully") //Here to know that runCmd ran as expected
	return nil
}

func main() {
	homeDir, _ := os.UserHomeDir()
	/* 	setCmd := os.Args[1]
	   	if setCmd == "--setenv" {
	   		parseArgs(os.Args[1:])
	   		os.Exit(1)
	   	} */
	result, err := loadConfig(homeDir + "/.ggiconfig.yml")

	if err != nil {
		fmt.Println("Error opening yml config file. Run 'touch .ggiconfig.yml' in your home directory")
		os.Exit(1)
	}

	Utils.LoadEnv(result.EnvLocation)

	c, err := parseArgs(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}

	err = runCmd(os.Stdin, os.Stdout, c)
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}
}
