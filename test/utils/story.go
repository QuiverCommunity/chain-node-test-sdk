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
	StoryContent []struct {
		OffsetHeight int64  `yaml:"offset_height"`
		FromKey      string `yaml:"from_key"`
		Action       string `yaml:"action"`
		Param        string `yaml:"param"`
	} `yaml:"story_content"`
}

// run parsed commands by command name and param based on offset block height
func FollowStory(testStory TestStory) (string, error) {
	// TODO create accounts from keys received from testStory
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
	// TODO run actions based on Story
	return log, nil
}
