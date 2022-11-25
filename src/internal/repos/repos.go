package repos

import (
	"fmt"
	"io"
	"net/http"

	Auth "GoGitIt/internal/auth"
)

func GetRepos() {
	//username := Auth.GetUsername()
	token := Auth.GetAuthToken()
	url := "https://api.github.com/users/SunnyD33/repos"
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
	fmt.Println(string(body))
}
