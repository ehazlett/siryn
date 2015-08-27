package server

import (
	"os/exec"

	log "github.com/Sirupsen/logrus"
)

func startPrometheus(binPath string, args []string) (*exec.Cmd, error) {
	log.Debugf("prometheus: path=%s args=%v", binPath, args)

	cmd := exec.Command(binPath, args...)
	if err := cmd.Start(); err != nil {
		return nil, err
	}

	return cmd, nil
}
