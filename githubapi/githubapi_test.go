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

func TestGetUserData(t *testing.T) {

	var nil string

	want := Userinfo{"Ocktokit", "https://github.com/Ocktokit", nil, nil, nil, 0, 0, 0}

	var username string = "ocktokit"
	data := GetUserData(username)

	if !reflect.DeepEqual(data, want) {
		t.Fatal("JSON Unmarshal failed.")
	}
}

func TestResponseToRepoData(t *testing.T) {

	data := []byte(
		`[
			{
				"id": 403538617,
				"node_id": "MDEwOlJlcG9zaXRvcnk0MDM1Mzg2MTc=",
				"name": "GitHub_REST_API_consumer",
				"full_name": "iamargus95/GitHub_REST_API_consumer",
				"private": false,
				"owner": {
					"login": "iamargus95",
					"id": 77744293,
					"node_id": "MDQ6VXNlcjc3NzQ0Mjkz",
					"avatar_url": "https://avatars.githubusercontent.com/u/77744293?v=4",
					"gravatar_id": "",
					"url": "https://api.github.com/users/iamargus95",
					"html_url": "https://github.com/iamargus95",
					"followers_url": "https://api.github.com/users/iamargus95/followers",
					"following_url": "https://api.github.com/users/iamargus95/following{/other_user}",
					"gists_url": "https://api.github.com/users/iamargus95/gists{/gist_id}",
					"starred_url": "https://api.github.com/users/iamargus95/starred{/owner}{/repo}",
					"subscriptions_url": "https://api.github.com/users/iamargus95/subscriptions",
					"organizations_url": "https://api.github.com/users/iamargus95/orgs",
					"repos_url": "https://api.github.com/users/iamargus95/repos",
					"events_url": "https://api.github.com/users/iamargus95/events{/privacy}",
					"received_events_url": "https://api.github.com/users/iamargus95/received_events",
					"type": "User",
					"site_admin": false
				},
				"html_url": "https://github.com/iamargus95/GitHub_REST_API_consumer",
				"description": null,
				"fork": false,
				"url": "https://api.github.com/repos/iamargus95/GitHub_REST_API_consumer",
				"forks_url": "https://api.github.com/repos/iamargus95/GitHub_REST_API_consumer/forks",
				"keys_url": "https://api.github.com/repos/iamargus95/GitHub_REST_API_consumer/keys{/key_id}",
				"collaborators_url": "https://api.github.com/repos/iamargus95/GitHub_REST_API_consumer/collaborators{/collaborator}",
				"teams_url": "https://api.github.com/repos/iamargus95/GitHub_REST_API_consumer/teams",
				"hooks_url": "https://api.github.com/repos/iamargus95/GitHub_REST_API_consumer/hooks",
				"issue_events_url": "https://api.github.com/repos/iamargus95/GitHub_REST_API_consumer/issues/events{/number}",
				"events_url": "https://api.github.com/repos/iamargus95/GitHub_REST_API_consumer/events",
				"assignees_url": "https://api.github.com/repos/iamargus95/GitHub_REST_API_consumer/assignees{/user}",
				"branches_url": "https://api.github.com/repos/iamargus95/GitHub_REST_API_consumer/branches{/branch}",
				"tags_url": "https://api.github.com/repos/iamargus95/GitHub_REST_API_consumer/tags",
				"blobs_url": "https://api.github.com/repos/iamargus95/GitHub_REST_API_consumer/git/blobs{/sha}",
				"git_tags_url": "https://api.github.com/repos/iamargus95/GitHub_REST_API_consumer/git/tags{/sha}",
				"git_refs_url": "https://api.github.com/repos/iamargus95/GitHub_REST_API_consumer/git/refs{/sha}",
				"trees_url": "https://api.github.com/repos/iamargus95/GitHub_REST_API_consumer/git/trees{/sha}",
				"statuses_url": "https://api.github.com/repos/iamargus95/GitHub_REST_API_consumer/statuses/{sha}",
				"languages_url": "https://api.github.com/repos/iamargus95/GitHub_REST_API_consumer/languages",
				"stargazers_url": "https://api.github.com/repos/iamargus95/GitHub_REST_API_consumer/stargazers",
				"contributors_url": "https://api.github.com/repos/iamargus95/GitHub_REST_API_consumer/contributors",
				"subscribers_url": "https://api.github.com/repos/iamargus95/GitHub_REST_API_consumer/subscribers",
				"subscription_url": "https://api.github.com/repos/iamargus95/GitHub_REST_API_consumer/subscription",
				"commits_url": "https://api.github.com/repos/iamargus95/GitHub_REST_API_consumer/commits{/sha}",
				"git_commits_url": "https://api.github.com/repos/iamargus95/GitHub_REST_API_consumer/git/commits{/sha}",
				"comments_url": "https://api.github.com/repos/iamargus95/GitHub_REST_API_consumer/comments{/number}",
				"issue_comment_url": "https://api.github.com/repos/iamargus95/GitHub_REST_API_consumer/issues/comments{/number}",
				"contents_url": "https://api.github.com/repos/iamargus95/GitHub_REST_API_consumer/contents/{+path}",
				"compare_url": "https://api.github.com/repos/iamargus95/GitHub_REST_API_consumer/compare/{base}...{head}",
				"merges_url": "https://api.github.com/repos/iamargus95/GitHub_REST_API_consumer/merges",
				"archive_url": "https://api.github.com/repos/iamargus95/GitHub_REST_API_consumer/{archive_format}{/ref}",
				"downloads_url": "https://api.github.com/repos/iamargus95/GitHub_REST_API_consumer/downloads",
				"issues_url": "https://api.github.com/repos/iamargus95/GitHub_REST_API_consumer/issues{/number}",
				"pulls_url": "https://api.github.com/repos/iamargus95/GitHub_REST_API_consumer/pulls{/number}",
				"milestones_url": "https://api.github.com/repos/iamargus95/GitHub_REST_API_consumer/milestones{/number}",
				"notifications_url": "https://api.github.com/repos/iamargus95/GitHub_REST_API_consumer/notifications{?since,all,participating}",
				"labels_url": "https://api.github.com/repos/iamargus95/GitHub_REST_API_consumer/labels{/name}",
				"releases_url": "https://api.github.com/repos/iamargus95/GitHub_REST_API_consumer/releases{/id}",
				"deployments_url": "https://api.github.com/repos/iamargus95/GitHub_REST_API_consumer/deployments",
				"created_at": "2021-09-06T08:10:48Z",
				"updated_at": "2021-09-07T12:29:27Z",
				"pushed_at": "2021-09-07T12:29:27Z",
				"git_url": "git://github.com/iamargus95/GitHub_REST_API_consumer.git",
				"ssh_url": "git@github.com:iamargus95/GitHub_REST_API_consumer.git",
				"clone_url": "https://github.com/iamargus95/GitHub_REST_API_consumer.git",
				"svn_url": "https://github.com/iamargus95/GitHub_REST_API_consumer",
				"homepage": null,
				"size": 13,
				"stargazers_count": 0,
				"watchers_count": 0,
				"language": "Go",
				"has_issues": true,
				"has_projects": true,
				"has_downloads": true,
				"has_wiki": true,
				"has_pages": false,
				"forks_count": 0,
				"mirror_url": null,
				"archived": false,
				"disabled": false,
				"open_issues_count": 0,
				"license": {
					"key": "apache-2.0",
					"name": "Apache License 2.0",
					"spdx_id": "Apache-2.0",
					"url": "https://api.github.com/licenses/apache-2.0",
					"node_id": "MDc6TGljZW5zZTI="
				},
				"forks": 0,
				"open_issues": 0,
				"watchers": 0,
				"default_branch": "master"
			},
			{
				"id": 401311266,
				"node_id": "MDEwOlJlcG9zaXRvcnk0MDEzMTEyNjY=",
				"name": "go-grep",
				"full_name": "iamargus95/go-grep",
				"private": false,
				"owner": {
					"login": "iamargus95",
					"id": 77744293,
					"node_id": "MDQ6VXNlcjc3NzQ0Mjkz",
					"avatar_url": "https://avatars.githubusercontent.com/u/77744293?v=4",
					"gravatar_id": "",
					"url": "https://api.github.com/users/iamargus95",
					"html_url": "https://github.com/iamargus95",
					"followers_url": "https://api.github.com/users/iamargus95/followers",
					"following_url": "https://api.github.com/users/iamargus95/following{/other_user}",
					"gists_url": "https://api.github.com/users/iamargus95/gists{/gist_id}",
					"starred_url": "https://api.github.com/users/iamargus95/starred{/owner}{/repo}",
					"subscriptions_url": "https://api.github.com/users/iamargus95/subscriptions",
					"organizations_url": "https://api.github.com/users/iamargus95/orgs",
					"repos_url": "https://api.github.com/users/iamargus95/repos",
					"events_url": "https://api.github.com/users/iamargus95/events{/privacy}",
					"received_events_url": "https://api.github.com/users/iamargus95/received_events",
					"type": "User",
					"site_admin": false
				},
				"html_url": "https://github.com/iamargus95/go-grep",
				"description": "Grep clone using Golang",
				"fork": false,
				"url": "https://api.github.com/repos/iamargus95/go-grep",
				"forks_url": "https://api.github.com/repos/iamargus95/go-grep/forks",
				"keys_url": "https://api.github.com/repos/iamargus95/go-grep/keys{/key_id}",
				"collaborators_url": "https://api.github.com/repos/iamargus95/go-grep/collaborators{/collaborator}",
				"teams_url": "https://api.github.com/repos/iamargus95/go-grep/teams",
				"hooks_url": "https://api.github.com/repos/iamargus95/go-grep/hooks",
				"issue_events_url": "https://api.github.com/repos/iamargus95/go-grep/issues/events{/number}",
				"events_url": "https://api.github.com/repos/iamargus95/go-grep/events",
				"assignees_url": "https://api.github.com/repos/iamargus95/go-grep/assignees{/user}",
				"branches_url": "https://api.github.com/repos/iamargus95/go-grep/branches{/branch}",
				"tags_url": "https://api.github.com/repos/iamargus95/go-grep/tags",
				"blobs_url": "https://api.github.com/repos/iamargus95/go-grep/git/blobs{/sha}",
				"git_tags_url": "https://api.github.com/repos/iamargus95/go-grep/git/tags{/sha}",
				"git_refs_url": "https://api.github.com/repos/iamargus95/go-grep/git/refs{/sha}",
				"trees_url": "https://api.github.com/repos/iamargus95/go-grep/git/trees{/sha}",
				"statuses_url": "https://api.github.com/repos/iamargus95/go-grep/statuses/{sha}",
				"languages_url": "https://api.github.com/repos/iamargus95/go-grep/languages",
				"stargazers_url": "https://api.github.com/repos/iamargus95/go-grep/stargazers",
				"contributors_url": "https://api.github.com/repos/iamargus95/go-grep/contributors",
				"subscribers_url": "https://api.github.com/repos/iamargus95/go-grep/subscribers",
				"subscription_url": "https://api.github.com/repos/iamargus95/go-grep/subscription",
				"commits_url": "https://api.github.com/repos/iamargus95/go-grep/commits{/sha}",
				"git_commits_url": "https://api.github.com/repos/iamargus95/go-grep/git/commits{/sha}",
				"comments_url": "https://api.github.com/repos/iamargus95/go-grep/comments{/number}",
				"issue_comment_url": "https://api.github.com/repos/iamargus95/go-grep/issues/comments{/number}",
				"contents_url": "https://api.github.com/repos/iamargus95/go-grep/contents/{+path}",
				"compare_url": "https://api.github.com/repos/iamargus95/go-grep/compare/{base}...{head}",
				"merges_url": "https://api.github.com/repos/iamargus95/go-grep/merges",
				"archive_url": "https://api.github.com/repos/iamargus95/go-grep/{archive_format}{/ref}",
				"downloads_url": "https://api.github.com/repos/iamargus95/go-grep/downloads",
				"issues_url": "https://api.github.com/repos/iamargus95/go-grep/issues{/number}",
				"pulls_url": "https://api.github.com/repos/iamargus95/go-grep/pulls{/number}",
				"milestones_url": "https://api.github.com/repos/iamargus95/go-grep/milestones{/number}",
				"notifications_url": "https://api.github.com/repos/iamargus95/go-grep/notifications{?since,all,participating}",
				"labels_url": "https://api.github.com/repos/iamargus95/go-grep/labels{/name}",
				"releases_url": "https://api.github.com/repos/iamargus95/go-grep/releases{/id}",
				"deployments_url": "https://api.github.com/repos/iamargus95/go-grep/deployments",
				"created_at": "2021-08-30T10:59:50Z",
				"updated_at": "2021-09-02T12:18:27Z",
				"pushed_at": "2021-09-03T06:53:26Z",
				"git_url": "git://github.com/iamargus95/go-grep.git",
				"ssh_url": "git@github.com:iamargus95/go-grep.git",
				"clone_url": "https://github.com/iamargus95/go-grep.git",
				"svn_url": "https://github.com/iamargus95/go-grep",
				"homepage": null,
				"size": 46,
				"stargazers_count": 0,
				"watchers_count": 0,
				"language": "Go",
				"has_issues": true,
				"has_projects": true,
				"has_downloads": true,
				"has_wiki": true,
				"has_pages": false,
				"forks_count": 0,
				"mirror_url": null,
				"archived": false,
				"disabled": false,
				"open_issues_count": 1,
				"license": {
					"key": "apache-2.0",
					"name": "Apache License 2.0",
					"spdx_id": "Apache-2.0",
					"url": "https://api.github.com/licenses/apache-2.0",
					"node_id": "MDc6TGljZW5zZTI="
				},
				"forks": 0,
				"open_issues": 1,
				"watchers": 0,
				"default_branch": "master"
			},
			{
				"id": 374159905,
				"node_id": "MDEwOlJlcG9zaXRvcnkzNzQxNTk5MDU=",
				"name": "iamargus95",
				"full_name": "iamargus95/iamargus95",
				"private": false,
				"owner": {
					"login": "iamargus95",
					"id": 77744293,
					"node_id": "MDQ6VXNlcjc3NzQ0Mjkz",
					"avatar_url": "https://avatars.githubusercontent.com/u/77744293?v=4",
					"gravatar_id": "",
					"url": "https://api.github.com/users/iamargus95",
					"html_url": "https://github.com/iamargus95",
					"followers_url": "https://api.github.com/users/iamargus95/followers",
					"following_url": "https://api.github.com/users/iamargus95/following{/other_user}",
					"gists_url": "https://api.github.com/users/iamargus95/gists{/gist_id}",
					"starred_url": "https://api.github.com/users/iamargus95/starred{/owner}{/repo}",
					"subscriptions_url": "https://api.github.com/users/iamargus95/subscriptions",
					"organizations_url": "https://api.github.com/users/iamargus95/orgs",
					"repos_url": "https://api.github.com/users/iamargus95/repos",
					"events_url": "https://api.github.com/users/iamargus95/events{/privacy}",
					"received_events_url": "https://api.github.com/users/iamargus95/received_events",
					"type": "User",
					"site_admin": false
				},
				"html_url": "https://github.com/iamargus95/iamargus95",
				"description": null,
				"fork": false,
				"url": "https://api.github.com/repos/iamargus95/iamargus95",
				"forks_url": "https://api.github.com/repos/iamargus95/iamargus95/forks",
				"keys_url": "https://api.github.com/repos/iamargus95/iamargus95/keys{/key_id}",
				"collaborators_url": "https://api.github.com/repos/iamargus95/iamargus95/collaborators{/collaborator}",
				"teams_url": "https://api.github.com/repos/iamargus95/iamargus95/teams",
				"hooks_url": "https://api.github.com/repos/iamargus95/iamargus95/hooks",
				"issue_events_url": "https://api.github.com/repos/iamargus95/iamargus95/issues/events{/number}",
				"events_url": "https://api.github.com/repos/iamargus95/iamargus95/events",
				"assignees_url": "https://api.github.com/repos/iamargus95/iamargus95/assignees{/user}",
				"branches_url": "https://api.github.com/repos/iamargus95/iamargus95/branches{/branch}",
				"tags_url": "https://api.github.com/repos/iamargus95/iamargus95/tags",
				"blobs_url": "https://api.github.com/repos/iamargus95/iamargus95/git/blobs{/sha}",
				"git_tags_url": "https://api.github.com/repos/iamargus95/iamargus95/git/tags{/sha}",
				"git_refs_url": "https://api.github.com/repos/iamargus95/iamargus95/git/refs{/sha}",
				"trees_url": "https://api.github.com/repos/iamargus95/iamargus95/git/trees{/sha}",
				"statuses_url": "https://api.github.com/repos/iamargus95/iamargus95/statuses/{sha}",
				"languages_url": "https://api.github.com/repos/iamargus95/iamargus95/languages",
				"stargazers_url": "https://api.github.com/repos/iamargus95/iamargus95/stargazers",
				"contributors_url": "https://api.github.com/repos/iamargus95/iamargus95/contributors",
				"subscribers_url": "https://api.github.com/repos/iamargus95/iamargus95/subscribers",
				"subscription_url": "https://api.github.com/repos/iamargus95/iamargus95/subscription",
				"commits_url": "https://api.github.com/repos/iamargus95/iamargus95/commits{/sha}",
				"git_commits_url": "https://api.github.com/repos/iamargus95/iamargus95/git/commits{/sha}",
				"comments_url": "https://api.github.com/repos/iamargus95/iamargus95/comments{/number}",
				"issue_comment_url": "https://api.github.com/repos/iamargus95/iamargus95/issues/comments{/number}",
				"contents_url": "https://api.github.com/repos/iamargus95/iamargus95/contents/{+path}",
				"compare_url": "https://api.github.com/repos/iamargus95/iamargus95/compare/{base}...{head}",
				"merges_url": "https://api.github.com/repos/iamargus95/iamargus95/merges",
				"archive_url": "https://api.github.com/repos/iamargus95/iamargus95/{archive_format}{/ref}",
				"downloads_url": "https://api.github.com/repos/iamargus95/iamargus95/downloads",
				"issues_url": "https://api.github.com/repos/iamargus95/iamargus95/issues{/number}",
				"pulls_url": "https://api.github.com/repos/iamargus95/iamargus95/pulls{/number}",
				"milestones_url": "https://api.github.com/repos/iamargus95/iamargus95/milestones{/number}",
				"notifications_url": "https://api.github.com/repos/iamargus95/iamargus95/notifications{?since,all,participating}",
				"labels_url": "https://api.github.com/repos/iamargus95/iamargus95/labels{/name}",
				"releases_url": "https://api.github.com/repos/iamargus95/iamargus95/releases{/id}",
				"deployments_url": "https://api.github.com/repos/iamargus95/iamargus95/deployments",
				"created_at": "2021-06-05T16:18:45Z",
				"updated_at": "2021-08-22T17:30:00Z",
				"pushed_at": "2021-08-22T17:29:58Z",
				"git_url": "git://github.com/iamargus95/iamargus95.git",
				"ssh_url": "git@github.com:iamargus95/iamargus95.git",
				"clone_url": "https://github.com/iamargus95/iamargus95.git",
				"svn_url": "https://github.com/iamargus95/iamargus95",
				"homepage": null,
				"size": 10,
				"stargazers_count": 0,
				"watchers_count": 0,
				"language": null,
				"has_issues": true,
				"has_projects": true,
				"has_downloads": true,
				"has_wiki": true,
				"has_pages": false,
				"forks_count": 0,
				"mirror_url": null,
				"archived": false,
				"disabled": false,
				"open_issues_count": 0,
				"license": null,
				"forks": 0,
				"open_issues": 0,
				"watchers": 0,
				"default_branch": "master"
			}
	]`)

	jsonResponse := responseToRepoData(data)

	want := ReposInfoArray{
		{"GitHub_REST_API_consumer", "https://github.com/iamargus95/GitHub_REST_API_consumer", 0},
		{"go-grep", "https://github.com/iamargus95/go-grep", 0},
		{"iamargus95", "https://github.com/iamargus95/iamargus95", 0},
	}

	if !reflect.DeepEqual(jsonResponse, want) {
		t.Fatal("JSON Unmarshal failed.")
	}
}

func TestUserData(t *testing.T) {

	data := Userinfo{"Ocktokit", "https://github.com/Ocktokit", "Ocktokit", "one2n.in", "Consulting", 0, 0, 0}

	want := "Name: Ocktokit,\nUsername: Ocktokit,\nE-mail: one2n.in,\nBio: Consulting,\nPublic Repositories: 0,\nFollowers: 0,\nFollowing: 0"

	actual := data.UserData()

	if actual != want {
		t.Fatal("JSON Unmarshal failed.")
	}
}
