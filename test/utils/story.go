package utils

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
