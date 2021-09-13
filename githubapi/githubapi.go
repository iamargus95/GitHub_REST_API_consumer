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
	Login        string `json:"login"`
	Html_url     string `json:"html_url"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Bio          string `json:"bio"`
	Public_repos int    `json:"public_repos"`
	Followers    int    `json:"followers"`
	Following    int    `json:"following"`
}

type ReposInfoJson struct {
	Name             string `json:"name"`
	Html_url         string `json:"html_url"`
	Stargazers_count int    `json:"stargazers_count"`
}

type userCollection struct {
	AccountInfo struct {
		Login        string `json:"login"`
		Html_url     string `json:"html_url"`
		Name         string `json:"name"`
		Email        string `json:"email"`
		Bio          string `json:"bio"`
		Public_repos int    `json:"public_repos"`
		Followers    int    `json:"followers"`
		Following    int    `json:"following"`
	}
	Repositories []ReposInfoJson
}

func responseToUserData(data []byte) Userinfo {
	var userData Userinfo
	_ = json.Unmarshal(data, &userData)
	return userData
}

func responseToRepoData(data []byte) []ReposInfoJson {
	var reposData []ReposInfoJson
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

func GetReposData(username string, noOfRepos int) []ReposInfoJson {

	var bodyJson []byte
	var result []ReposInfoJson

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

func Fetch(username string) userCollection {
	accountData := GetUserData(username)
	repoData := GetReposData(username, accountData.Public_repos)

	return (userCollection{accountData, repoData})
}

func MarshalFetchData(data userCollection) []byte {
	accountByte, _ := json.MarshalIndent(data, " ", "  ")
	return accountByte
}
