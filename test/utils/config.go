package utils

import (
	"os"

	"gopkg.in/yaml.v2"
)

type TestConfig struct {
	Node               string `yaml:"node"`
	CliPath            string `yaml:"cli_path"`
	Chain              string `yaml:"chain"`
	InitialBlockHeight int64  `yaml:"initial_block_height"`
	StoriesDir         string `yaml:"stories_dir"`
}

var Config TestConfig

func init() {
	ReadConfig()
}
func ReadConfig() (TestConfig, error) {
	configFilePath := "/config.yml"
	file, err := os.Open(configFilePath)
	if err == nil {
		defer file.Close()

		decoder := yaml.NewDecoder(file)
		err = decoder.Decode(&Config)
	}
	return Config, err
}
