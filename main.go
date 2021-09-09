package main

import (
	"flag"
	"iamargus95/fetchGithubData/githubapi"
	"iamargus95/fetchGithubData/io"
	"strings"
	"sync"
)

func main() {

	var con bool
	flag.BoolVar(&con, "con", false, "Runs the application concurrently.")

	flag.Parse()
	usernames := flag.Args()

	if con {

		var wg sync.WaitGroup
		dataToFile := make(chan []string, 1)

		for _, username := range usernames {
			wg.Add(1)
			go worker(username, dataToFile, &wg)
			go writeFile(username, dataToFile, &wg)
		}
		wg.Wait()

	} else {

		for _, username := range usernames {
			userDetails := githubapi.GetUserData(username)
			reposDetails := githubapi.GetReposData(username, userDetails.Public_repos)
			userdata := strings.Split(userDetails.UserData(), ",")
			repodata := reposDetails.RepoData()
			accountData := append(userdata, repodata...)
			io.WriteToFile(username, accountData)
		}
	}

}

func worker(username string, dataToFile chan []string, wg *sync.WaitGroup) {

	defer wg.Wait()
	userDetails := githubapi.GetUserData(username)
	reposDetails := githubapi.GetReposData(username, userDetails.Public_repos)
	userdata := strings.Split(userDetails.UserData(), ",")
	repodata := reposDetails.RepoData()
	accountData := append(userdata, repodata...)
	dataToFile <- accountData
}

func writeFile(username string, dataToFile chan []string, wg *sync.WaitGroup) {
	defer wg.Done()
	result := <-dataToFile
	io.WriteToFile(username, result)
}
