package main

import (
	"encoding/json"
	"iamargus95/fetchGithubData/githubapi"
	"iamargus95/fetchGithubData/io"
	"sync"
)

type userCollection struct {
	accountInfo githubapi.Userinfo
	repoInfo    []githubapi.ReposInfoJson
}

func fetch(username string) userCollection {
	accountData := githubapi.GetUserData(username)
	repoData := githubapi.GetReposData(username, accountData.Public_repos)

	return (userCollection{accountInfo: accountData, repoInfo: repoData})
}

func marshalFetchData(userCollection) []byte {
	var data userCollection
	marshalAccountData, _ := json.MarshalIndent(data.accountInfo, " ", "  ")
	marshalRepoData, _ := json.MarshalIndent(data.repoInfo, " ", "  ")
	result := append(marshalAccountData, marshalRepoData...)
	return result
}

func sequence(usernames []string) {

	for _, username := range usernames {
		resultJson := fetch(username)
		resultByte := marshalFetchData(resultJson)
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

	resultJson := fetch(username)
	resultByte := marshalFetchData(resultJson)

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
