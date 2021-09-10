package main

import (
	"iamargus95/fetchGithubData/githubapi"
	"iamargus95/fetchGithubData/io"
	"os"
	"strings"
	"sync"
)

func sequence(usernames []string) {

	for _, username := range usernames {
		userDetails := githubapi.GetUserData(username)
		reposDetails := githubapi.GetReposData(username, userDetails.Public_repos)
		userdata := strings.Split(userDetails.UserData(), ",")
		repodata := reposDetails.RepoData()
		accountData := append(userdata, repodata...)
		io.WriteToFile(username, accountData)
	}

}

func concurrently(usernames []string) {

	var wg sync.WaitGroup

	dataToFile := make(chan map[string][]string)

	for _, username := range usernames {

		wg.Add(1)
		go fetch(username, dataToFile, &wg)
		go writeFile(username, dataToFile, &wg)
	}

	wg.Wait()

}

func fetch(username string, dataToFile chan map[string][]string, wg *sync.WaitGroup) {

	defer wg.Wait()
	userDetails := githubapi.GetUserData(username)
	reposDetails := githubapi.GetReposData(username, userDetails.Public_repos)
	userdata := strings.Split(userDetails.UserData(), ",")
	repodata := reposDetails.RepoData()
	accountData := append(userdata, repodata...)

	dataToChannel := make(map[string][]string)
	dataToChannel[username] = accountData
	dataToFile <- dataToChannel

}

func writeFile(username string, dataToFile chan map[string][]string, wg *sync.WaitGroup) {

	result := <-dataToFile
	value := result[username]
	io.WriteToFile(username, value)
	i, _ := os.Stat(username + ".txt")
	if i.Size() != 0 {
		defer wg.Done()
	} else {
		wg.Wait()
	}
}
