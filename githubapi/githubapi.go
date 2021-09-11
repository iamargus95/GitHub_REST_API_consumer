package githubapi

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

const (
	SCHEME = "https"
	HOST   = "api.github.com"
	PATH1  = "users/"
	PATH2  = "/repos"
	QUERY  = "type=public&per_page=100&page="
)

type Userinfo struct {
	Login        string
	Html_url     string
	Name         string
	Email        string
	Bio          string
	Public_repos int
	Followers    int
	Following    int
}

type ReposInfoJson struct {
	Name             string
	Html_url         string
	Stargazers_count int
}

type ReposInfoArray []ReposInfoJson

func responseToUserData(data []byte) Userinfo {
	var userData Userinfo
	_ = json.Unmarshal(data, &userData)
	return userData
}

func responseToRepoData(data []byte) ReposInfoArray {
	var reposDataArray ReposInfoArray
	_ = json.Unmarshal(data, &reposDataArray)
	return reposDataArray
}

func GetUserData(username string) Userinfo {

	url1 := url.URL{
		Scheme: SCHEME,
		Host:   HOST,
		Path:   PATH1 + username,
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", url1.String(), nil)
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

func GetReposData(username string, noOfRepos int) ReposInfoArray {

	var bodyJson []byte
	var result ReposInfoArray

	for i := 1; i <= ((noOfRepos / 100) + 1); i++ {

		url1 := url.URL{
			Scheme:   SCHEME,
			Host:     HOST,
			Path:     PATH1 + username + PATH2,
			RawQuery: QUERY + strconv.Itoa(i),
		}

		client := &http.Client{}
		req, err := http.NewRequest("GET", url1.String(), nil)

		if err != nil {
			log.Fatal(err)
		}
		res, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()

		bodyJson, err = ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}
		searchResult := responseToRepoData(bodyJson)
		result = append(result, searchResult...)
	}

	return result
}

func (data *Userinfo) UserData() string {

	var stringToPrint string

	if data.Name != "" {
		stringToPrint = "Name: " + data.Name + ",\nUsername: " + data.Login + ",\nE-mail: " + data.Email + ",\nBio: " + data.Bio +
			",\nPublic Repositories: " + strconv.Itoa(data.Public_repos) + ",\nFollowers: " + strconv.Itoa(data.Followers) +
			",\nFollowing: " + strconv.Itoa(data.Following)
	}
	return stringToPrint
}

func (data ReposInfoArray) RepoData() []string {

	var stringToPrint []string

	for i := 0; i < len(data); i++ {

		stars := data[i].Stargazers_count
		stringToPrint = append(stringToPrint, "\nRepository No["+strconv.Itoa(i+1)+"]:"+data[i].Name+
			".\nAvailable at :"+data[i].Html_url+".\nStars Count :"+strconv.Itoa(stars))
	}

	return stringToPrint
}
