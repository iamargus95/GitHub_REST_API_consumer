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
	Login        string `json:"Username"`
	Html_url     string `json:"URL"`
	Name         string `json:"Name"`
	Email        string `json:"Email"`
	Bio          string `json:"Bio"`
	Public_repos int    `json:"Public_Repos"`
	Followers    int    `json:"Followers"`
	Following    int    `json:"Following"`
}

type ReposInfoJson struct {
	Name             string `json:"Repo_Name"`
	Html_url         string `json:"URL"`
	Stargazers_count int    `json:"Stars"`
}

type ReposInfoJsonArray []ReposInfoJson

func responseToUserData(data []byte) Userinfo {
	var userData Userinfo
	_ = json.Unmarshal(data, &userData)
	return userData
}

func responseToRepoData(data []byte) ReposInfoJsonArray {
	var reposData ReposInfoJsonArray
	_ = json.Unmarshal(data, &reposData)
	return reposData
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

func GetReposData(username string, noOfRepos int) ReposInfoJsonArray {

	var bodyJson []byte
	var result ReposInfoJsonArray

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
		searchResult := responseToRepoData(bodyJson) //Unmarshall ReposJson to ReposInfoJsonArray.
		result = append(result, searchResult...)     //Append ReposJson after changing each Page query.
	}
	return result
}
