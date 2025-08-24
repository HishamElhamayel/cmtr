package git

import (
	"os/exec"
)

func Commit(message string) (string, error) {
	cmd := exec.Command("git", "commit", "-m", message)

	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return string(output), nil

}
