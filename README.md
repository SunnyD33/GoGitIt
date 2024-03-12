# GoGitIt (Go Get It)

GoGitIt is a CLI tool that uses the GitHub API for users to be able to look up repos by name in the command line

## Installation Steps
- Download the current release from this <a href="https://github.com/SunnyD33/GoGitIt/releases/tag/v1.00">link</a></p>
- You will want to install/move this file into your 'usr/local/bin' directory. Once that is done, you should be able to use the ggi command anywhere in your terminal

## Requirements
### Configuration Yaml File
- The program will require that a yml file is created and that file will hold your auth status and the location of the env file.
- You will need to run ```touch .ggiconfig.yml``` in your home directory: 

### Example Of The Yaml File
<img width="566" alt="Screenshot 2023-11-26 at 8 13 11 PM" src="https://github.com/SunnyD33/GoGitIt/assets/44623894/196f6aea-f0df-44e8-bef3-5b22e0644ea5">

If the config file is not created first, the program will alert you that the yaml file will need to be created in the home direcroty
The values that are present can be set manually but the program can also take care of updating the file on the initial run as it will check for the yaml file, so adding values manually is not fully necessary.

### .Env file
- Create a .env file in a direcrory of your choosing. You will update the location of the env file for the program to read from in the .ggiconfig file (see previous section).
- There are two values that will need to be added into the file:
  ```GH_USER``` and ```GH_TOKEN```
- ```GH_USER``` is your profile user name from github. Since this tool uses the GitHub API, be be sure to double check your casing in your name and make sure it matches your username as you would see it in the link to your GitHub profile.
- ```GH_TOKEN``` is a generated personal token that you will need to create in your GitHub profile. Refer to GitHubs documentation on how to do this <a href="https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/managing-your-personal-access-tokens">here</a>. KEEP THIS TOKEN SECURE! Setting expiration dates for your tokens is recommended!

### Example Of The .Env File
<img width="566" alt="Screenshot 2023-11-26 at 10 40 54 PM" src="https://github.com/SunnyD33/GoGitIt/assets/44623894/f235dd87-f3ac-4c86-9702-1b40e978fd8f">

Once these steps have been taken, you should be ready to use the tool!

## Usage
- Here are the list fo commands that are available for this tool:
```go
func HelpCommand() {
    fmt.Println("Commands:")
    fmt.Println("-h / --help \n  Shows help for available commands. Can be used for each command as well")
    fmt.Println("setup  \n  Allows user to check current auth status, authorize themselves and change .env file location")
    fmt.Println("list  \n  Allows user to list repos for an entered username")
    fmt.Println("open \n  Opens an enetered repo on your browser")
    fmt.Println("search \n  Searches for repos based on entered criteria")
    fmt.Println("rate \n  Shows your current rate limit")
}
```
Subcommands are also available for some of the commands. These will be listed below in detail for usage of each command

- ```ggi --help``` or ```ggi -h``` will print the above help text to the available commands that can be used with the tool
- ```ggi setup``` will check your auth state and let you know if you are authorized and allow you to update your .env file location via subcommands (see below). Please note, if you get an error or you do not get a return, please check make sure your token has not expired.
  - ```ggi setup -a``` is used to set your auth state to true if it is currently not set. This will be based off if an env file is created with the proper parameters set.
  - ```ggi setup -setenv``` is used to be able to set the location of your env file. You will need to enter the full path to the file, including the file name of the env file.
  - ```ggi setup -s``` is used to check your current auth status.
- ```ggi list``` will allow you to see repos that a specific user has on their GitHub if you are trying to find a specific repo. Please note, this command does NOT show private repos unless you are searching against your own username.
  - ```ggi list -r 'username'``` is used to list repos for a specific user. They will be listed in the command line.
- ```ggi open``` will open the entered repo on your browser. Please note, there may be some browsers that are potentially not supported.
  - ```ggi open 'username'``` or ```ggi open 'username/repo'``` will open the entered GitHub profile or Repo for the user that was entered.
  - ```ggi open 'username/repo' -i``` or ```ggi open 'username/repo -p``` will open the repo on either the issues or pull requests page, respectively
- ```ggi search``` will allow you to search for a list of repos based on different crtieria to help narrow down finding repos for your general or specific needs. This command has multiple options to help refine searchs.
  - ```ggi search -q 'querystring'``` will search for repos in GitHub that match the query that was entered. -q is required for searches to occur. The optional strings are as follows:
  ```
  -c int
    	sets how many results are displayed (max = 100) (default 30)
  -l string
    	refine repo search by programming language
  -o string
    	ordering option - acceptable values: desc, asc (default "desc")
  -q string
    	query value used to search for repos that contain the given value (required)
  -s string
    	sorting option - acceptable values: stars, forks, help-wanted-issues, updated (default "stars")
  ```
  ## Notes
  - In case the ggi command does not work, and you have it in your usr/local/bin directory, you made need to run the chmod command on the file to give it permissions to run as an executable.
  - The program does not keep track of your token if it is expired. In case commands return an unmarshal error please check that the token that you are using is not expired.
