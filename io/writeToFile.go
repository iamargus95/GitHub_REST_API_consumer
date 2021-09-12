package io

import (
	"bufio"
	"log"
	"os"
)

func WriteToFile(username string, data []byte) error {

	file, err := os.OpenFile(username+".json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	defer file.Close()

	datawriter := bufio.NewWriter(file)

	_, _ = datawriter.Write(data)

	datawriter.Flush()

	return err
}
