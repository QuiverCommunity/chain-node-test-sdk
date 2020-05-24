package test

import (
	"os"
	"path/filepath"
	"testing"

	story "github.com/QuiverCommunity/chain-node-test-sdk/test/stories"
	"github.com/QuiverCommunity/chain-node-test-sdk/test/utils"
	"gopkg.in/yaml.v2"
)

func RegisterActions() {
	// Register functions that are used in stories
	// You can also register custom functions for customized blockchain test
	utils.RegisterAction("send_balance", story.SendBalance)
	utils.RegisterAction("check_balance", story.CheckBalance)
}

func TestStories(t *testing.T) {
	var paths []string

	// read test configuration
	_, configErr := utils.ReadConfig()
	if configErr != nil {
		t.Fatal("reading configuration file failure", configErr)
	}
	RegisterActions()

	// read and run stories
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

		// run parsed command on test story
		log, storyErr := utils.FollowStory(testStory)

		// log the result of the story
		t.Log("Please check story log for", testStory.Name)
		t.Log(log)
		t.Log("-----------------------------------------")
		if storyErr != nil {
			t.Fatal("FAIL:\terror running story for", path, "\n", storyErr)
		} else {
			t.Log("PASS:\tsuccess running story for", path)
		}
		t.Log("That's all of the log collected for", testStory.Name)
	}
}
