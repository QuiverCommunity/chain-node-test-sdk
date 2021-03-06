package utils

import (
	"fmt"
	"os/exec"
	"strings"
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
	if Config.Node == "" {
		return args
	}
	if !shouldConfigureNode(args) {
		return args
	}
	return append(args, "--node", Config.Node)
}

func configureKeyring(args []string) []string {
	if len(args) == 0 {
		return args
	}
	switch args[0] {
	case "status":
		return args
	case "query":
		if len(args) > 1 && args[1] == "account" {
			return args
		}
	}
	return append(args, "--keyring-backend", "test")
}

func RunCli(args []string) ([]byte, string, error) {
	return RunCliStdin(args, "")
}

func RunCliStdin(args []string, stdin string) ([]byte, string, error) {
	args = configureNode(args)
	args = configureKeyring(args)

	cmd := exec.Command(Config.CliPath, args...)
	cmd.Stdin = strings.NewReader(stdin)
	cmdLog := fmt.Sprintf("%s %s", Config.CliPath, strings.Join(args, " "))
	if stdin != "" {
		cmdLog += fmt.Sprintf("<<< %s", stdin)
	}
	res, err := cmd.CombinedOutput()
	return res, cmdLog, err
}

func RunShellScript(args []string, stdin string) ([]byte, string, error) {
	cmd := exec.Command("bash", args...)
	cmd.Stdin = strings.NewReader(stdin)
	cmdLog := fmt.Sprintf("bash %s <<< %s", strings.Join(args, " "), stdin)
	res, err := cmd.CombinedOutput()
	return res, cmdLog, err
}
