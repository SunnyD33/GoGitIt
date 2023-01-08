package help

import (
    "fmt"
)

func HelpCommand() {
    fmt.Println("Commands:")
    fmt.Println("-h / --help \n  Shows help for available commands. Can be used for each command as well")
    fmt.Println("setup  \n  Allows user to check current auth status, authorize themselves and change .env file location")
    fmt.Println("list  \n  Allows user to list repos for an entered username")
    fmt.Println("open \n  Opens an enetered repo on your browser")
    fmt.Println("search \n  Searches for repos based on entered criteria")
    fmt.Println("rate \n  Shows your current rate limit")
}
