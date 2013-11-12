package execwrapper

import (
	"os/exec"
)

func Command(cmd string) (output *exec.Cmd) {
	output = exec.Command("bash", "-c", cmd)
	return output
}
