package io

import "testing"

func TestWriteToFile(t *testing.T) { //Fix Test
	const filename = "test"
	channel := make(chan []string)
	var content = []string{"This is a test file.", "Delete this after t.Run."}
	channel <- content

	t.Run("test", func(t *testing.T) {
		if err := WriteToFile(filename, channel); (err != nil) != false {
			t.Errorf("Test failed : %v", err)
		}
	})

}
