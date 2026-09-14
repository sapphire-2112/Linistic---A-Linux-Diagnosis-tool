
package networking

import "os/exec"

func run(cmd string, args ...string) string {
	out, _ := exec.Command(cmd, args...).CombinedOutput()
	return string(out)
}