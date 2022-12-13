package search

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	Auth "GoGitIt/internal/auth"
)

type Items struct {
	Items []Searches `json:"items"`
}

type Searches struct {
	Name string `json:"full_name"`
}

func Search(query string, language string, sort string, order string, per_page string) {
	token := Auth.GetAuthToken()
	url := "https://api.github.com/search/repositories?q=" + query + "+language:" + language + "&sort=" + sort + "&order=" + order + 
	"&per_page=" + per_page
	method := "GET"

	client := http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Add("Accept", "application/vnd.github+json")
	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("X-GitHub-Api-Version", "2022-11-28")

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

	//fmt.Println(string(body))

	//Parse the body of the request into a Searches struct
	var items Items

	err = json.Unmarshal(body, &items)
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(items.Items) == 0 {
		fmt.Println("No results found. Please try searching again.")
	}

	for _, item := range items.Items {
		fmt.Println(item.Name)
	}
}
