package global

import (
	"errors"
	"os/exec"
)

func InstallSelf() error {
	cmdString := "cp ./md /usr/bin"
	cmd := exec.Command("bash", "-c", cmdString)
	if _, err := cmd.Output(); err != nil {
		return errors.New("安装失败,详细原因:" + err.Error())
	}
	return nil
}