package repos

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	Auth "GoGitIt/internal/auth"
)

// Create a struct to hold the repo data
type Repos struct {
	Name string `json:"full_name"`
	Message string `json:"message"`
}

func GetRepos(user string) {
	var username string
	if user == "" {
		username = Auth.GetUsername()
	} else {
		username = user
	}

	token := Auth.GetAuthToken()
	url := "https://api.github.com/users/" + username + "/repos"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Accept", "application/vnd.github+json")
	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Cookie", "_octo=GH1.1.1748435795.1665149718; logged_in=no")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	//Parse the body into a struct
	var repos []Repos

	err = json.Unmarshal(body, &repos)
	if err != nil {
		//user is not found
		fmt.Println("User not found or user has no repos")
		fmt.Println("Please check your spelling and try again")
		return
	}

	//Loop through the repos and print out the names
	for _, repo := range repos {
		fmt.Println(repo.Name)
	}
}
