package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/fatih/color"
	"github.com/google/uuid"
)

var (
	infoLog  = log.New(os.Stdout, color.GreenString("INFO\t"), log.Ldate|log.Ltime)
	errorLog = log.New(os.Stderr, color.RedString("ERROR\t"), log.Ldate|log.Ltime|log.Lshortfile)
)

func main() {

	config, err := parseConfig("../config.yaml")
	if err != nil {
		errorLog.Fatalf("Error parsing config: %v", err)
	}

	repo := &Repo{
		URL: config.App.Repo_URL,
		ID:  uuid.NewString(),
		Dir: config.App.Directory,
	}

	infoLog.Println("Cloning the repository.")
	repo.cloneRepo()

	infoLog.Printf("Language is: %s\n", config.App.Language)

	infoLog.Println("Installing dependencies.")
	err = installDependencies(config.App.Language)

	if err != nil {
		log.Fatal(err)
	}

	infoLog.Println("Building application.")
	// fmt.Println(config.App.Directory)
	err = buildApplication(config.App.Directory, config.App.Language, config.App.Name)

	if err != nil {
		log.Fatal(err)
	}

	infoLog.Println("Application built.")

	infoLog.Println(config.App.Arguments)
	if err != nil {
		log.Fatal(err)
	}

	path, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	full_path := path + "/" + config.App.Directory + "/" + config.App.Name

	infoLog.Println(full_path)
	data := ServiceData{
		Description:   config.App.Name,
		ExecStart:     full_path,
		ExecStartArgs: config.App.Arguments,
	}

	err = CreateAndStartService(data)
	if err != nil {
		log.Fatal("Error creating and starting service:", err)
		os.Exit(1)
	}

	infoLog.Println("Application running as systemd service.")
	infoLog.Println("Running caddy server for deploying application.")

	_, err = exec.LookPath("caddy")
	if err != nil {
		fmt.Println("caddy command not found in $PATH")
		os.Exit(1)
	}

	infoLog.Println(config.App.Domain)

	caddy := CaddyData{
		Domain:       config.App.Domain,
		ReverseProxy: config.App.Proxy,
	}

	err = startCaddy(caddy)

	if err != nil {
		log.Fatal("Error starting caddy:", err)
	}
}
