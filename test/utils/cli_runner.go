package utils

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

func RunCli(args []string, stdin string) ([]byte, string, error) {
	args = configureNode(args)

	cmd := exec.Command(Config.CliPath, args...)
	cmd.Stdin = strings.NewReader(stdin)
	cmdLog := fmt.Sprintf("%s %s <<< %s", Config.CliPath, strings.Join(args, " "), stdin)
	res, err := cmd.CombinedOutput()
	return res, cmdLog, err
}