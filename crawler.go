package main

import (
	"os"
	"os/exec"
)

const directory string = "websites"

func ScrapeWebsite(url string) {
	cmd := exec.Command("wget", "-r", "-A html,pdf", "-l 10", "--no-parent", "--directory-prefix="+directory, url)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
