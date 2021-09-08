package io

import "testing"

const filename = "test"

var content = []string{"This is a test file.", "Delete this after t.Run."}

func TestWriteToFile(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		if err := WriteToFile(filename, content); (err != nil) != false {
			t.Errorf("Test failed : %v", err)
		}
	})

}
