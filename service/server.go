package service

import (
	"os/exec"
)

// WebServer 做出反应的结构体
type WebServer struct {
}

// DoExec
func (svc *WebServer) DoExec(requestID string) (string, error) {
	command := "./push-copy.sh"
	//cmd := exec.Command("cmd", "-c", command)
	cmd := exec.Command("bash", "-c", command)

	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}
