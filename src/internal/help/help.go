package help

import (
    "fmt"
)

func HelpCommand() {
    fmt.Println("Commands:")
    fmt.Println("-h / --help \n  Shows help for available commands. Can be used for each command as well")
    fmt.Println("-a  \n  Authorize yourself to increase rate limit of use and access some commands that requires auth")
    fmt.Println("-r  \n  Searches for the repo that was entered by username")
    fmt.Println("-s \n  Checks your auth status")
    fmt.Println("-o \n  Opens an enetered repo on your browser")
    fmt.Println("--setenv \n allows you to set where the tool will read data from your .env file")
    fmt.Println("search \n  Searches for repos based on entered criteria")
    fmt.Println("rate \n  Shows your current rate limit")
}
