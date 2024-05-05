package main

import (
	"log"
	"os"
	"text/template"
)

const caddyFileTemplate = `{{.Domain}} {
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

	path, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	caddyfile := "Caddyfile"

	// _, err = os.Stat(path + "/" + caddyfile)
	// if err == nil {
	// 	return nil
	// }

	file, err := os.Create(path + "/" + caddyfile)
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

	err = runCommand("caddy", "stop")
	if err != nil {
		return err
	}

	err = runCommand("caddy", "start", "--config", path+"/"+caddyfile)
	if err != nil {
		return err
	}

	return nil
}
