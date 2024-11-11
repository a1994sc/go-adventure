package cmd

// setBaseDirectory sets the base directory.
// Borrowed from https://github.com/zarf-dev/zarf/blob/main/src/cmd/cmd.go
func setBaseDirectory(args []string) string {
	if len(args) > 0 {
		return args[0]
	}
	return "."
}
