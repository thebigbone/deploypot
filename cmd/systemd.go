package main

import (
	"fmt"
	"os"
	"os/exec"
	"text/template"
)

const unitFileTemplate = `
[Unit]
Description={{.Description}}
After=network.target

[Service]
Type=simple
ExecStart={{.ExecStart}}
Restart=always

[Install]
WantedBy=multi-user.target
`

type ServiceData struct {
	Description string
	ExecStart   string
}

func CreateAndStartService(data ServiceData) error {
	tmpl, err := template.New("unitFile").Parse(unitFileTemplate)
	if err != nil {
		return err
	}

	_, err = os.Stat("/etc/systemd/system/deploy.service")
	if err == nil {
		return nil
	}

	file, err := os.Create("/etc/systemd/system/deploy.service")
	if err != nil {
		return err
	}

	defer func() {
		err := file.Close()
		if err != nil {
			return
		}
	}()

	err = tmpl.Execute(file, data)
	if err != nil {
		return err
	}

	err = runCommand("systemctl", "daemon-reload")
	if err != nil {
		return err
	}

	err = runCommand("systemctl", "enable", "deploy.service")
	if err != nil {
		return err
	}

	err = runCommand("systemctl", "start", "deploy.service")
	if err != nil {
		return err
	}

	return nil
}

func runCommand(command string, args ...string) error {
	cmd := exec.Command(command, args...)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error running command '%s %s': %w", command, args, err)
	}
	return nil
}
