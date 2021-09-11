# GitHub_REST_API_consumer

Build simple GitHub REST API consumer to download information 
- followers
- repos
- orgs, etc

for a multiple users and store in the filesystem.

# Instructions to run the program :

In the root of the repository open the terminal and run the `make build` command.

- `make all` or `make build` : Builds binary.

- `make clean`: Removes binary if any.
	
- `make run`: Executes the binary.

# USAGE :

## Sequential program execution:

`make build`

& then

`./fetchGithubData [username1] [username2] [username3] .... `

## Concurrent program execution:

`make build` 

& then

`./fetchGithubData -con [username1] [username2] [username3] .....` 