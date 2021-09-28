package main

import (
	"iamargus95/fetchGithubData/githubapi"
	"iamargus95/fetchGithubData/io"
	"sync"
)

func sequence(usernames []string) {

	for _, username := range usernames {
		resultByte := githubapi.Fetch(username)
		io.WriteToFile(username, resultByte)
	}
}

func concurrently(usernames []string) {

	var wg sync.WaitGroup

	for _, username := range usernames {
		wg.Add(1)
		go worker(username, &wg)
	}
	wg.Wait()

}

func worker(username string, wg *sync.WaitGroup) {
	defer wg.Done()
	resultByte := githubapi.Fetch(username)
	io.WriteToFile(username, resultByte)
}
