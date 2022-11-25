package auth

import (
	Utils "GoGitIt/pkg/utils"
	"fmt"
)

var AuthToken string
var Username string

func PrintAuthorizedText() {
	fmt.Println("You are currently authorized as " + Utils.GetEnvWithKey("GH_USER"))
	fmt.Println("Note: Your auth status is based on if you have a GitHub personal access token saved in your .env file. It cannot track if it has expired.")
	fmt.Println("In the case that commands no longer work, please check that your personal access token has not expired.")
	fmt.Println("If so, a new token will need to be created and you will need to update your .env file with the new token")
}

func PrintUnauthorizedText() {
	fmt.Println("You are currently unauthorized. Use 'ggi -a' to authorize yourself with a GitHub personal access token from your .env file")
	fmt.Println("Once authorization has been completed, a yaml file, '.ggiconfig.yaml', will be created holding your authorization status.")
	fmt.Println("This will be checked against a personal access token in your .env file to confirm that you are authorized.")
	fmt.Println("For more information on how to build the .env file for authorization, please check out the README.md file.")
}

func GetAuthToken() string {
	getToken := Utils.GetEnvWithKey("GH_TOKEN")
	AuthToken = getToken

	return AuthToken
}

func GetUsername() string {
	getUsername := Utils.GetEnvWithKey("GH_USER")
	Username = getUsername

	return Username
}

func CheckAuthToken() bool {
	if Utils.GetEnvWithKey("GH_TOKEN") == "" {
		return false
	} else {
		return true
	}
}
