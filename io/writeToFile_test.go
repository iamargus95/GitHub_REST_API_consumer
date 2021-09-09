package io

import "testing"

func TestWriteToFile(t *testing.T) { //Fix Test
	const filename = "test"
	var content = []string{"This is a test file.", "Delete this after t.Run."}

	t.Run("test", func(t *testing.T) {
		if err := WriteToFile(filename, content); (err != nil) != false {
			t.Errorf("Test failed : %v", err)
		}
	})

}
