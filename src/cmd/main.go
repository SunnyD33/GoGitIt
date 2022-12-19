package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"gopkg.in/yaml.v2"

	Auth "GoGitIt/internal/auth"
	Open "GoGitIt/internal/open"
	Rate "GoGitIt/internal/rate"
	Repos "GoGitIt/internal/repos"
	Search "GoGitIt/internal/search"
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
}

func main() {
	var c Config

	loadErr := Utils.LoadEnv(checkEnvLocation())

	if loadErr != nil {
		fmt.Println(loadErr)
		homeDir, err := os.UserHomeDir()

		if err != nil {
			fmt.Println(err)
			return
		}
		envFileLocation, _ := Utils.SetEnv(os.Stdin, os.Stdout)

		result, _ := loadConfig(homeDir + "/.ggiconfig.yml")
		if err != nil {
			fmt.Println(err)
		}

		result.updateEnvLocation(envFileLocation)
		result.updateAuthState(false)

		saveConfig(result, homeDir+"/.ggiconfig.yml")

		return
	}

	homeDir, _ := os.UserHomeDir()
	_, err := loadConfig(homeDir + "/.ggiconfig.yml")

	//This should not be reached.. Something is really wrong if this fires
	if err != nil {
		fmt.Println("Error opening yml config file. Run 'touch .ggiconfig.yml' in your home directory")
		return
	}

	if len(os.Args) < 2 {
		fmt.Println("Invalid number of arguments")
		printHelpText()
		return
	}

	//Create commands for the user to use
	helpCommand := flag.NewFlagSet("-h", flag.ExitOnError)
	repoCommand := flag.NewFlagSet("-r", flag.ExitOnError)
	statusCommand := flag.NewFlagSet("-s", flag.ExitOnError)
	setEnvCommand := flag.NewFlagSet("--setenv", flag.ExitOnError)
	authCommand := flag.NewFlagSet("-a", flag.ExitOnError)
	searchCommand := flag.NewFlagSet("search", flag.ExitOnError)
	rateCommand := flag.NewFlagSet("rate", flag.ExitOnError)
	openCommand := flag.NewFlagSet("-o", flag.ExitOnError)

	//Create subcommands for the user to use on specific commands
	searchQuery := searchCommand.String("q", "", "Used to search for repos that contain the given value (Required)")
	searchLanguage := searchCommand.String("l", "", "Refine repo search by language")
	searchSort := searchCommand.String("s", "stars", "Acceptable values: stars, forks, help-wanted-issues, updated")
	searchOrder := searchCommand.String("o", "desc", "Acceptable values: desc, asc")
	searchCount := searchCommand.Int("c", 30, "Sets how many results are displayed (Max = 100)")

	// openPulls := openCommand.String("p","-p","Opens pull request for entered repo via your browser")
	// openIssues := openCommand.String("i","-i", "Opens issue for entered repo via your browser")

	switch os.Args[1] {
	case "-h":
		helpCommand.Parse(os.Args[2:])
	case "-r":
		repoCommand.Parse(os.Args[2:])
	case "-s":
		statusCommand.Parse(os.Args[2:])
	case "-a":
		authCommand.Parse(os.Args[2:])
	case "-o":
		openCommand.Parse(os.Args[2:])
	case "--setenv":
		setEnvCommand.Parse(os.Args[2:])
	case "search":
		searchCommand.Parse(os.Args[2:])
	case "rate":
		rateCommand.Parse(os.Args[2:])
	}

	//Check which commands are parsed
	if helpCommand.Parsed() {
		if len(os.Args) > 2 {
			fmt.Println("Too many arguments!")
			printHelpText()
		} else {
			printHelpText()
		}
	}

	if repoCommand.Parsed() {
		if len(os.Args) < 3 {
			Repos.GetRepos("")
		} else {
			Repos.GetRepos(os.Args[2])
		}
	}

	if statusCommand.Parsed() {
		if len(os.Args) > 2 {
			fmt.Println("Too many arguments! Please enter only -s flag to get the current authorization status")
		} else {
			authStatus := checkAuthStatus()
			if authStatus.IsAuthorized {
				Auth.PrintAuthorizedText()
			} else {
				Auth.PrintUnauthorizedText()
			}
		}
	}

	if authCommand.Parsed() {
		fmt.Println("Checking for token...")
		homeDir, _ := os.UserHomeDir()
		authToken := Auth.CheckAuthToken()

		//TODO: Add logic to not authorize user if already authorized
		result, _ := loadConfig(homeDir + "/.ggiconfig.yml")

		if result.IsAuthorized {
			fmt.Println("You are already authorized! Cancelling operation!")
			return
		}

		if !authToken {
			fmt.Println("Authorization failed")
			fmt.Println("Unable to find token in your .env file. Please confirm that the GH_TOKEN variable is not empty.")
		} else {
			homeDir, err := os.UserHomeDir()
			if err != nil {
				fmt.Println(err)
				return
			}

			result, _ := loadConfig(homeDir + "/.ggiconfig.yml")

			c.updateAuthState(true)
			c.updateEnvLocation(result.EnvLocation)
			saveConfig(c, homeDir+"/.ggiconfig.yml")
			fmt.Println("Authorization successful!")
		}
	}

	if openCommand.Parsed() {
		if len(os.Args) < 3 {
			fmt.Println("Please enter a user or repo to open on your broswer")
			fmt.Println("Path can either be just a username or a username/repo")
			openCommand.PrintDefaults()
			return
		} else if len(os.Args) == 3 {
			Open.OpenRepo(os.Args[2], "none")
		} else if len(os.Args) > 4 {
			fmt.Println("Too many arguments! Can either use only -i or -p as subcommands for -o")
		} else if os.Args[2] != "" && os.Args[3] == "-i" {
			Open.OpenRepo(os.Args[2], "issues")
		} else if os.Args[2] != "" && os.Args[3] == "-p" {
			Open.OpenRepo(os.Args[2], "pulls")
		} else {
			fmt.Println("Please enter a user or repo to open on your broswer")
		}
	}

	if setEnvCommand.Parsed() {
		if len(os.Args) > 2 {
			fmt.Println("Too many arguments! Please enter only --setenv flag to set a custom .env file location")
		} else {
			envFileLocation, err := Utils.SetEnv(os.Stdin, os.Stdout)

			if err != nil {
				fmt.Println(err)
				return
			}

			result, err := loadConfig(homeDir + "/.ggiconfig.yml")
			if err != nil {
				fmt.Println(err)
				return
			}

			c.updateEnvLocation(envFileLocation)
			c.updateAuthState(result.IsAuthorized)
			saveConfig(c, homeDir+"/.ggiconfig.yml")
		}
	}

	if searchCommand.Parsed() {
		sortChoices := [4]string{"stars", "forks", "help-wanted-issues", "updated"}
		orderByChoice := [2]string{"desc", "asc"}

		//Query is required
		if *searchQuery == "" {
			searchCommand.PrintDefaults()
			return
		}

		//Check to make sure that the sort option is valid
		i := 0 //Counting variabls in array
		for _, choice := range sortChoices {
			if *searchSort != choice && i < len(sortChoices)-1 {
				i++
				continue
			} else if *searchSort == choice {
				break
			} else {
				searchCommand.PrintDefaults()
				return
			}
		}

		//Check to make sure that the order option is valid
		i = 0 //Counting variabls in array
		for _, choice := range orderByChoice {
			if *searchOrder != choice && i < len(orderByChoice)-1 {
				i++
				continue
			} else if *searchOrder == choice {
				break
			} else {
				searchCommand.PrintDefaults()
				return
			}
		}

		//Check value of count. If not set, use default value of 30
		if *searchCount != 30 {
			if *searchCount <= 0 || *searchCount > 100 {
				searchCommand.PrintDefaults()
				return
			}
		}

		countString := strconv.Itoa(*searchCount)

		Search.Search(*searchQuery, *searchLanguage, *searchSort, *searchOrder, countString)
	}

	if rateCommand.Parsed() {
		if len(os.Args) > 3 {
			fmt.Println("Too many arguments")
			rateCommand.PrintDefaults()
		} else {
			Rate.GetRate()
		}
	}
}
