package open

import (
	"fmt"
	"os/exec"
	"runtime"
)

func OpenRepo(repoPath string, option string) {
	fmt.Println("Opening broswer..")
	var err error

	url := "http://www.github.com/" + repoPath
	var urlWithOption string

	if option == "issues" {
		urlWithOption = "http://www.github.com/" + repoPath + "/issues"
	} else if option == "pulls" {
		urlWithOption = "http://www.github.com/" + repoPath + "/pulls"
	}

	if option == "none" {
		switch runtime.GOOS {
		case "linux":
			err = exec.Command("xdg-open", url).Start()
		case "windows":
			err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
		case "darwin":
			err = exec.Command("open", url).Start()
		default:
			err = fmt.Errorf("unsupported platform")
		}
		if err != nil {
			fmt.Println(err)
		}
	} else {
		switch runtime.GOOS {
		case "linux":
			err = exec.Command("xdg-open", urlWithOption).Start()
		case "windows":
			err = exec.Command("rundll32", "url.dll,FileProtocolHandler", urlWithOption).Start()
		case "darwin":
			err = exec.Command("open", urlWithOption).Start()
		default:
			err = fmt.Errorf("unsupported platform")
		}
		if err != nil {
			fmt.Println(err)
		}
	}

}
