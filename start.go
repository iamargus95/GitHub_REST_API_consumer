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

	dataToFile := make(chan map[string][]byte)

	for _, username := range usernames {
		go conFetch(username, dataToFile, &wg)
	}

	wg.Add(len(usernames))

	for _, username := range usernames {
		go writeFile(username, dataToFile, &wg)
	}
	wg.Wait()

}

func conFetch(username string, dataToFile chan map[string][]byte, wg *sync.WaitGroup) {

	resultByte := githubapi.Fetch(username)
	dataToChannel := make(map[string][]byte)
	dataToChannel[username] = resultByte
	dataToFile <- dataToChannel
}

func writeFile(username string, dataToFile chan map[string][]byte, wg *sync.WaitGroup) {

	result := <-dataToFile
	value := result[username]
	io.WriteToFile(username, value)
	defer wg.Done()
}
