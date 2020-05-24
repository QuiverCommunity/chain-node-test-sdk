package test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/QuiverCommunity/chain-node-test-sdk/test/utils"
	"gopkg.in/yaml.v2"
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
		// parse story.yml file
		file, fopenError := os.Open(path)
		if fopenError != nil {
			t.Fatal("file open failure", fopenError)
		}
		decoder := yaml.NewDecoder(file)
		testStory := utils.TestStory{}
		storyDecodeErr := decoder.Decode(&testStory)
		if storyDecodeErr != nil {
			t.Fatal("story decode failure", storyDecodeErr)
		}
		t.Log("Running story on for", testStory.Name, "for", path)
		t.Log(testStory)
		// TODO run parsed commands by command name and param based on offset block height
	}
}
