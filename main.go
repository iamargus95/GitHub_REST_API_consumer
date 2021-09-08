package main

import (
	"bufio"
	"fmt"
	"iamargus95/fetchGithubData/githubapi"
	"iamargus95/fetchGithubData/io"
	"os"
	"strings"
	"sync"
)

func main() {

	usernames := getUsername()

	var wg sync.WaitGroup
	dataToFile := make(chan []string)

	for _, username := range usernames {
		wg.Add(1)
		go worker(username, dataToFile, &wg)
		io.WriteToFile(username, dataToFile)
	}

	wg.Wait()
}

func getUsername() []string {
	fmt.Println("Enter the desired GitHub usernames separated by a space and press enter : ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSuffix(input, "\n")
	argString := strings.Split(input, " ")
	return argString
}

func worker(username string, dataToFile chan []string, wg *sync.WaitGroup) {

	defer wg.Done()
	userDetails := githubapi.GetUserData(username)
	reposDetails := githubapi.GetReposData(username, userDetails.Public_repos)
	userdata := strings.Split(userDetails.UserData(), ",")
	repodata := reposDetails.RepoData()
	accountData := append(userdata, repodata...)
	dataToFile <- accountData
}
