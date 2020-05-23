package utils

import (
	"fmt"
	"os/exec"
	"strings"

	cf "github.com/QuiverCommunity/chain-node-test-sdk/config"
)

func shouldConfigureNode(args []string) bool {
	if len(args) == 0 {
		return false
	}
	switch args[0] {
	case "query", "tx", "status":
		return true
	default:
		return false
	}
}

func configureNode(args []string) []string {
	if cf.Config.Node == "" {
		return args
	}
	if !shouldConfigureNode(args) {
		return args
	}
	return append(args, "--node", cf.Config.Node)
}

func RunCli(args []string, stdin string) ([]byte, string, error) {
	args = configureNode(args)

	cmd := exec.Command(cf.Config.CliPath, args...)
	cmd.Stdin = strings.NewReader(stdin)
	cmdLog := fmt.Sprintf("%s %s <<< %s", cf.Config.CliPath, strings.Join(args, " "), stdin)
	res, err := cmd.CombinedOutput()
	return res, cmdLog, err
}
