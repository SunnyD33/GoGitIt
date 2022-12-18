package open

import (
	"fmt"
	"os/exec"
	"runtime"
)

func OpenRepo(repoPath string) {
	fmt.Println("Opening broswer..")

	url := "http://www.github.com/" + repoPath
	var err error
	
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
}
