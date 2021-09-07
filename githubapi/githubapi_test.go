package githubapi

import (
	"reflect"
	"testing"
)

func TestResponseToUserData(t *testing.T) {
	data := []byte(`{
		"login": "Ocktokit",
		"id": 41244090,
		"node_id": "MDEyOk9yZ2FuaXphdGlvbjQxMjQ0MDkw",
		"avatar_url": "https://avatars.githubusercontent.com/u/41244090?v=4",
		"gravatar_id": "",
		"url": "https://api.github.com/users/Ocktokit",
		"html_url": "https://github.com/Ocktokit",
		"followers_url": "https://api.github.com/users/Ocktokit/followers",
		"following_url": "https://api.github.com/users/Ocktokit/following{/other_user}",
		"gists_url": "https://api.github.com/users/Ocktokit/gists{/gist_id}",
		"starred_url": "https://api.github.com/users/Ocktokit/starred{/owner}{/repo}",
		"subscriptions_url": "https://api.github.com/users/Ocktokit/subscriptions",
		"organizations_url": "https://api.github.com/users/Ocktokit/orgs",
		"repos_url": "https://api.github.com/users/Ocktokit/repos",
		"events_url": "https://api.github.com/users/Ocktokit/events{/privacy}",
		"received_events_url": "https://api.github.com/users/Ocktokit/received_events",
		"type": "Organization",
		"site_admin": false,
		"name": "name",
		"company": null,
		"blog": "",
		"location": null,
		"email": "one2n.in",
		"hireable": null,
		"bio": "Consulting",
		"twitter_username": null,
		"public_repos": 0,
		"public_gists": 0,
		"followers": 0,
		"following": 0,
		"created_at": "2018-07-14T23:40:57Z",
		"updated_at": "2018-07-14T23:40:57Z"
	}`)

	jsonResponse := responseToUserData(data)
	want := Userinfo{"Ocktokit", "https://github.com/Ocktokit", "name", "one2n.in", "Consulting", 0, 0, 0}

	if !reflect.DeepEqual(jsonResponse, want) {
		t.Fatal("JSON Unmarshal failed.")
	}
}

func TestResponseToRepoData(t *testing.T) {
	data := []byte(`[
	{
		"name": "GitHub_REST_API_consumer",
		"html_url": "https://github.com/iamargus95/GitHub_REST_API_consumer",
		"stargazers_count": 0
	},
	{
		"name": "go-grep",
		"html_url": "https://github.com/iamargus95/go-grep",
		"stargazers_count": 0,
	},
	{
		"name": "iamargus95",
		"html_url": "https://github.com/iamargus95/iamargus95",
		"stargazers_count": 0,
	}
	]`)

	jsonResponse := responseToRepoData(data)

	want := ReposInfoArray{
		
		[]ReposInfoJson{
			{"GitHub_REST_API_consumer", "https://github.com/iamargus95/GitHub_REST_API_consumer", 0},
			{"go-grep", "https://github.com/iamargus95/go-grep", 0},
			{"iamargus95", "https://github.com/iamargus95/iamargus95", 0},
		}
	} //Fix this test 

	if !reflect.DeepEqual(jsonResponse, want) {
		t.Fatal("JSON Unmarshal failed.")
	}
}
