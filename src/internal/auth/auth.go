package auth

import (
	Utils "GoGitIt/pkg/utils"
	"fmt"
)

type Auth struct {
	AuthToken string
}

// Exported auth variable for access in main.go
var A Auth

func (auth *Auth) updateAuthState(token string) {
	auth.AuthToken = token
}

func PrintAuthorizedText() {
	fmt.Println("You are currently authorized as " + Utils.GetEnvWithKey("GH_USER"))
	fmt.Println("Note: Your auth status is based on if you have a GitHub personal access token saved in your .env file. It cannot track if it has expired.")
	fmt.Println("In the case that commands no longer work, please check that your personal access token has not expired.")
	fmt.Println("If so, a new token will need to be created and you will need to reauthorize with the new token using the 'ggi -a' command")
}

func PrintUnauthorizedText() {
	fmt.Println("You are currently unauthorized. Use 'ggi -a' to authorize yourself with a GitHub personal access token.")
	fmt.Println("For more information on how to build the .env file for authorization, please check out the README.md file.")
}

func SetAuthState() string {
	getToken := Utils.GetEnvWithKey("GH_TOKEN")
	A.AuthToken = getToken

	A.updateAuthState(A.AuthToken)

	return A.AuthToken
}

func CheckAuthState() bool {
	if A.AuthToken == "" {
		return false
	} else {
		return true
	}
}
