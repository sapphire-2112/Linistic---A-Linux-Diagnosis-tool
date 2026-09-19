package utils

import "os/exec"

func Run(command string, args ...string) string {
    cmd := exec.Command(command, args...)

    output, err := cmd.CombinedOutput()
    if err != nil {
        return ""
    }

    return string(output)
}