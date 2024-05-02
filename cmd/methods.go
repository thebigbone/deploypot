package main

import (
	"log"
	"os"
	"os/exec"
	"syscall"

	"github.com/go-git/go-git/v5"
)

type Repo struct {
	URL string
	ID  string
	Dir string
}

func (r *Repo) cloneRepo() error {

	_, err := git.PlainClone(r.Dir, false, &git.CloneOptions{
		URL:      r.URL,
		Progress: os.Stdout,
	})

	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func installDependencies(language_name string) error {
	if language_name == "go" {
		cmd := exec.Command("go", "mod", "tidy")

		err := cmd.Run()

		if err != nil {
			log.Fatal(err)
			return err
		}
	}

	return nil
}

func buildApplication(directory, language_name, app_name string) error {
	if language_name == "go" {
		cmd := exec.Command("go", "build", "-o", app_name)

		cmd.Dir = directory

		err := cmd.Run()

		if err != nil {
			log.Fatal(err)
			return err
		}
	}

	return nil
}

func runApplication(directory, language_name, app_name string, arguments ...string) error {
	if language_name == "go" {
		cmd := exec.Command("./"+app_name, arguments...)

		cmd.Dir = directory

		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}

		err := cmd.Start()

		if err != nil {
			log.Fatal(err)
			return err
		}

		err = cmd.Process.Release()

		if err != nil {
			log.Fatal(err)
			return err
		}
	}

	return nil
}
