package rate

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	Auth "GoGitIt/internal/auth"
)

type Limits struct {
	Resources struct {
		Core struct {
			Limit     int `json:"limit"`
			Used      int `json:"used"`
			Remaining int `json:"remaining"`
			Reset     int `json:"reset"`
		} `json:"core"`
		Search struct {
			Limit     int `json:"limit"`
			Used      int `json:"used"`
			Remaining int `json:"remaining"`
			Reset     int `json:"reset"`
		} `json:"search"`
	}
}

func GetRate() {
	token := Auth.GetAuthToken()
	url := "https://api.github.com/rate_limit"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Add("Accept", "application/vnd.github+json")
	req.Header.Add("Authorization", "Bearer "+token)

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

	limits := Limits{}

	err = json.Unmarshal(body, &limits)
	if err != nil {
		fmt.Println(err)
		return
	}

	//print out the core data
	fmt.Println()
	fmt.Println("Core Usage")
	fmt.Println("----------")
	fmt.Println("Core Limit: ", limits.Resources.Core.Limit)
	fmt.Println("Used: ", limits.Resources.Core.Used)
	fmt.Println("Remaining: ", limits.Resources.Core.Remaining)
	fmt.Println("Reset: ", limits.Resources.Core.Reset)
	fmt.Println()

	//Print out the search data
	fmt.Println("Search Usage")
	fmt.Println("------------")
	fmt.Println("Search Limit: ", limits.Resources.Search.Limit)
	fmt.Println("Used: ", limits.Resources.Search.Used)
	fmt.Println("Remaining: ", limits.Resources.Search.Remaining)
	fmt.Println("Reset: ", limits.Resources.Search.Reset)

}
