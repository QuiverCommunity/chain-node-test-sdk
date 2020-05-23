package test

import (
	"os"
	"path/filepath"
	"testing"
)

func TestStories(t *testing.T) {
	var paths []string

	stories_directory := "stories"
	err := filepath.Walk(stories_directory, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			paths = append(paths, path)
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	for _, path := range paths {
		if filepath.Ext(path) != ".yml" {
			continue
		}
		t.Log("Running story on for", path)
		// TODO should parse story.yml file
		// TODO run parsed commands by command name and param based on offset block height
	}
}
