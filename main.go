package main

import (
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"strings"
)

func main() {
	args := os.Args[1:]
	newArgs := parse(args)
	cmd := exec.Command("go", newArgs...)
	output, err := cmd.CombinedOutput()
	fmt.Print(string(output))
	exitCode := cmd.ProcessState.ExitCode()
	if err != nil {
		os.Exit(exitCode)
		return
	}

}

func parse(args []string) []string {
	newArgs := []string{"get"}
	for _, arg := range args {
		parsed, err := url.Parse(arg)
		if err != nil {
			newArgs = append(newArgs, arg)
			continue
		}
		newArgs = append(newArgs, fmt.Sprintf("%s%s", parsed.Host, strings.TrimSuffix(parsed.Path, ".git")))
	}
	return newArgs
}
