package main

import (
	"fmt"
	"iamargus95/githubassignment/githubapi"
	"iamargus95/githubassignment/io"
)

func main() {
	fmt.Println("\n Enter desired GitHub username : \n ")
	username := getUsername()
	userDetails := githubapi.UserData(username)
	reposDetails := githubapi.ReposData(username)
	userdata := githubapi.FileUserData(userDetails)
	repodata := githubapi.FileRepoData(userDetails, reposDetails)
	io.WriteToFile(username, userdata)
	io.WriteToFile(username, repodata)
}

func getUsername() string {
	var username string
	fmt.Scanln(&username)
	return username
}
