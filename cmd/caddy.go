package main

import (
	"os"
	"text/template"
)

const caddyFileTemplate = `
{{.Domain}} {
    reverse_proxy {{.ReverseProxy}}
    encode zstd gzip
}
`

type CaddyData struct {
	Domain       string
	ReverseProxy string
}

func startCaddy(data CaddyData) error {
	tmpl, err := template.New("caddyFile").Parse(caddyFileTemplate)
	if err != nil {
		return err
	}

	_, err = os.Stat("/etc/caddy/Caddyfile")
	if err == nil {
		return nil
	}

	file, err := os.Create("/etc/caddy/Caddyfile")
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

	err = runCommand("caddy", "start")
	if err != nil {
		return err
	}

	return nil
}
