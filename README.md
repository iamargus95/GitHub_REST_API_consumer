# GitHub_REST_API_consumer

Build simple GitHub REST API consumer to download information 
- followers
- repos
- orgs, etc

for a multiple users and store in the filesystem.

# Instructions to run the program :

In the root of the repository open the terminal and run the `make run` command.

- `make all`: Builds binary.

- `make clean`: Removes binary if any.
	
- `make run`: Executes the binary.

USAGE :

`go run main.go [username1] [username2] [username3] .... ` OR 

`./fetchGithubData [username1] [username2] [username3] .... `