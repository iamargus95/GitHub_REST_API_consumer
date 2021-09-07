package main

import (
	"fmt"
	"iamargus95/githubassignment/githubapi"
	"iamargus95/githubassignment/io"
	"strings"
)

func main() {

	fmt.Println("\n Enter desired GitHub username : \n ")
	username := getUsername()
	userDetails := githubapi.GetUserData(username)
	reposDetails := githubapi.GetReposData(username, userDetails.Public_repos)
	userdata := githubapi.UserData(userDetails)
	repodata := githubapi.RepoData(reposDetails)
	io.WriteToFile(username, strings.Split(userdata, ","))
	io.WriteToFile(username, repodata)
}

func getUsername() string {
	var username string
	fmt.Scanln(&username)
	return username
}
