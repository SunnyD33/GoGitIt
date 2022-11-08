package auth

import (
	"fmt"
	"GoGitIt/pkg/utils"
)

func PrintAuthorizedText() {
	fmt.Println("You are currently authorized as " + utils.GetEnvWithKey("GH_USER"))
	fmt.Println("Note: Your auth status is based on if you have a GitHub personal access token saved! It cannot track if it has expired.")
	fmt.Println("In the case that commands no longer work, please check that your personal access token has not expired.")
	fmt.Println("If so, a new token will need to be created and you will need to reauthorize with the new token using the 'ggi -a' command")
}

func PrintUnauthorizedText() {
	fmt.Println("You are currently unauthorized. Use 'ggi -a' to authorize yourself with a GitHub personal access token.")
}