package githubapi

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type Userinfo struct {
	Login        string
	Url          string
	Html_url     string
	Name         string
	Email        string
	Bio          string
	Public_repos int
	Followers    int
	Following    int
}

type ReposInfoJson struct {
	Name     string
	Html_url string
}

type ReposInfoArray []ReposInfoJson

func responseToUserData(data []byte) Userinfo {
	var searchResult Userinfo
	_ = json.Unmarshal(data, &searchResult)
	return searchResult
}

func responseToRepoData(data []byte) ReposInfoArray {
	var reposArray ReposInfoArray
	_ = json.Unmarshal(data, &reposArray)
	return reposArray
}

func UserData(username string) Userinfo {

	url := "https://api.github.com/users/" + username

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	//Send request
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	//Read the response body.
	bodyJson, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	searchResult := responseToUserData(bodyJson)
	return searchResult
}

func ReposData(username string) ReposInfoArray {

	url := "http://api.github.com/users/" + username + "/repos?per_page=100"

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Fatal(err)
	}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	bodyJson, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	searchResult := responseToRepoData(bodyJson)

	return searchResult
}

func FileUserData(data Userinfo) []string {
	var stringToPrint []string
	if data.Name != "" {
		stringToPrint = []string{"Name: " + data.Name + "\n\nUsername: " + data.Login + "\n\nE-mail: " + data.Email + "\n\nBio: " + data.Bio +
			"\nPublic Repositories: " + strconv.Itoa(data.Public_repos) + "\n\nFollowers: " + strconv.Itoa(data.Followers) +
			"\n\nFollowing: " + strconv.Itoa(data.Following)}
	}

	return stringToPrint
}

func FileRepoData(user Userinfo, data ReposInfoArray) []string {

	var stringToPrint []string
	var loopCondition int

	if user.Public_repos >= 100 {
		loopCondition = 100
	} else {
		loopCondition = user.Public_repos
	}

	for i := 0; i < loopCondition; i++ {

		stringToPrint = append(stringToPrint, "\nRepository No["+strconv.Itoa(i+1)+"]:"+data[i].Name+". \nAvailable at:"+data[i].Html_url)
	}

	return stringToPrint
}
