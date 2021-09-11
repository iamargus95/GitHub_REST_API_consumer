package io

import (
	"bufio"
	"log"
	"os"
)

func WriteToFile(username string, data []string) error {

	file, err := os.OpenFile(username+".txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	defer file.Close()

	datawriter := bufio.NewWriter(file)

	for _, data := range data {
		_, _ = datawriter.WriteString(data + "\n")
	}

	datawriter.Flush()

	return err
}
