package githubapi

import (
	"os"
	"reflect"
	"testing"
)

func TestResponseToUserData(t *testing.T) {

	data, _ := os.ReadFile("./tests/userdata.json")
	jsonResponse := responseToUserData(data)
	want := Userinfo{"Ocktokit", "https://github.com/Ocktokit", "name", "one2n.in", "Consulting", 0, 0, 0}

	if !reflect.DeepEqual(jsonResponse, want) {
		t.Fatal("JSON Unmarshal failed.")
	}
}

func TestResponseToRepoData(t *testing.T) {

	data, _ := os.ReadFile("./tests/repodata.json")
	jsonResponse := responseToRepoData(data)

	want := []ReposInfoJson{
		{"GitHub_REST_API_consumer", "https://github.com/iamargus95/GitHub_REST_API_consumer", 0},
		{"go-grep", "https://github.com/iamargus95/go-grep", 0},
		{"iamargus95", "https://github.com/iamargus95/iamargus95", 0},
	}

	if !reflect.DeepEqual(jsonResponse, want) {
		t.Fatal("JSON Unmarshal failed.")
	}
}

func TestGetUserData(t *testing.T) {

	var nil string

	want := Userinfo{"Ocktokit", "https://github.com/Ocktokit", nil, nil, nil, 0, 0, 0}

	var username string = "ocktokit"
	data := GetUserData(username)

	if !reflect.DeepEqual(data, want) {
		t.Fatal("JSON Unmarshal failed.")
	}
}
