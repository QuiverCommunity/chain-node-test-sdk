package utils

import (
	"fmt"
)

type TestStory struct {
	Name     string `yaml:"name"`
	Accounts []struct {
		Key     string `yaml:"key"`
		Secret  string `yaml:"secret"`
		Address string `yaml:"address"`
	} `yaml:"accounts"`
	StoryContent []Action `yaml:"story_content"`
}

// run parsed commands by command name and param based on offset block height
func FollowStory(testStory TestStory) (string, error) {
	// create accounts from keys received from testStory
	log := ""
	for index, account := range testStory.Accounts {
		log += "\n"
		addedKeyBytes, cmdLog, err := RunCliStdin([]string{"keys", "add", account.Key}, "y\n"+account.Secret)
		log += fmt.Sprintf("adding %dth account for %s\n", index, account.Key)
		log += cmdLog
		log += "\n"
		log += string(addedKeyBytes)
		if err != nil {
			log += "\n"
			log += err.Error()
		}
	}
	// run worker for the story
	RunWorker(testStory.StoryContent)
	return log, nil
}
