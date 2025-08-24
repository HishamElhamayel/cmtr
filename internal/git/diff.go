package git

import "os/exec"

func GetDiff() (string, error) {
	cmd := exec.Command("git", "diff", "--cached")

	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return string(output), nil

}