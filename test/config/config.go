package config

type TestConfig struct {
	Node string `yaml:"node"`
	CliPath string `yaml:"cli_path"`
	Chain string `yaml:"chain"`
	InitialBlockHeight int64 `yaml:"initial_block_height"`
	StoriesDir string `yaml:"stories_dir"`
}

var Config TestConfig

func init() {
}
