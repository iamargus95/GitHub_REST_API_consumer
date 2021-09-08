.PHONY = all clean run test
all: build

build: main.go
	@echo "Building executable binary..."
	go build -o fetchGithubData.o

clean:
	@echo "Cleaning up..."
	rm fetchGithubData.o
	go clean

run:
	go run .

test:
	go test