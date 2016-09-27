package main

import (
	"log"
	"os/exec"
	"regexp"
	"strings"
)

// TODO real remote
func main() {
	_, err := exec.Command("ping", "-c 1", "gitlab.liip.ch").Output()
	if err != nil {
		log.Fatal("Network is down... or you forgot VPN")
	}
	remote := getRemote()

	exec.Command("open", remote).Output()

	log.Println("(╯°□°）╯︵ ┻━┻ ", remote)
}

func getRemote() string {
	remote, err := exec.Command("git", "remote", "-v").Output()
	if err != nil {
		log.Fatal("Go inside git repo!")
	}
	url := string(remote)
	r, _ := regexp.Compile(`origin\s*http([a-z]://.+)*.git`)
	url = strings.Replace(r.FindString(url), ".git", "", 1)
	url = strings.Replace(url, "origin", "", 1)
	url = strings.TrimSpace(url)
	return url
}
