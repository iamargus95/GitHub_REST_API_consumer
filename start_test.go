package main

import (
	"os"
	"testing"
)

var usernames = []string{"iamargus95", "jdk2588", "kalpeshdeo", "chinmay185"}

func deleteFiles(usernames []string) {
	for _, username := range usernames {
		os.Remove(username + ".txt")
	}
}

func benchmarkSequence(noOfExec int, b *testing.B) {

	for i := 0; i < b.N; i++ {
		sequence(usernames)
	}
}

func benchmarkConcurrent(noOfExec int, b *testing.B) {

	for i := 0; i < b.N; i++ {
		concurrently(usernames)
	}
}

func BenchmarkSequence100(b *testing.B) {
	benchmarkSequence(5, b)
	deleteFiles(usernames)
}

func BenchmarkConcurrent100(b *testing.B) {
	benchmarkConcurrent(5, b)
	deleteFiles(usernames)
}
