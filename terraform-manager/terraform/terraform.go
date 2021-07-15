package terraform

import (
	"bufio"
	"errors"
	"os"
	"os/exec"

	"github.com/juju/loggo"
)

var logger = loggo.GetLogger("terraform")

type Terraform struct {
}

func (tf *Terraform) AddScript(script string) {
	logger.Infof("Adding new script...")
	file, err := os.Create("./tmp.tf")
	defer file.Close()

	if err != nil {
		logger.Errorf("Got error while creating script file. Error:", err.Error())
		return
	}
	writer := bufio.NewWriter(file)
	defer writer.Flush()

	_, err = writer.WriteString(script)
	if err != nil {
		logger.Errorf("Got error while writing to a script file. Error: %s", err.Error())
	}
}

func (tf *Terraform) Validate() (string, error) {
	cmd := exec.Command("./terraform", "validate")
	output, err := cmd.StdoutPipe()
	if err != nil {
		logger.Errorf("can not get pipe from terraform. Reason: ", err.Error())
		return "", errors.New("can not get pipe from terraform")
	}
	err = cmd.Start()
	if err != nil {
		logger.Errorf("can not run command. Reason: ", err.Error())
		return "", errors.New("can not run command")
	}

	out := make([]byte, 1024)
	n, err := output.Read(out)
	if err != nil {
		logger.Errorf("can not read terraform output. Reason: ", err.Error())
		return "", errors.New("can not read terraform output")
	}
	if err := cmd.Wait(); err != nil {
		logger.Errorf(err.Error())
	}
	return string(out[:n]), nil
}
