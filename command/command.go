package command

import (
	"os/exec"
	"time"
	"errors"
)

// Executes the command and return the combined output.
func RunCommand(command string, args ...string) (output string, err error) {
	cmd := exec.Command(command, args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	output = string(out)
	return output, nil
}

// Run a command, kill it if timeout
func RunCommandWithTimeout(timeout time.Duration, command string, args ...string) (err error) {
	cmd := exec.Command(command, args...)
	err = cmd.Start()
	if err != nil {
		return
	}

	done := make(chan error, 1)
	t := time.After(timeout)
	go func() {
		done <- cmd.Wait()
	}()
	select {
	case err = <-done:
		return
	case <-t:
		cmd.Process.Kill()
		err = errors.New("command timed out")
		return
	}
	return
}
