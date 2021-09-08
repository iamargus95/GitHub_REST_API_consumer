package main

import (
	"flag"
	"iamargus95/fetchGithubData/githubapi"
	"iamargus95/fetchGithubData/io"
	"strings"
	"sync"
)

func main() {
	flag.Parse()
	usernames := flag.Args()

	var wg sync.WaitGroup
	dataToFile := make(chan []string)

	for _, username := range usernames {
		wg.Add(1)
		go worker(username, dataToFile, &wg)
		io.WriteToFile(username, dataToFile)
	}

	wg.Wait()
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
