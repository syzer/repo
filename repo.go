package main

import (
	"log"
	"os/exec"
	"regexp"
	"strings"
)

// TODO extract real domain
func main() {
	_, err := exec.Command("ping", "-c 1", "gitlab.liip.ch").Output()
	if err != nil {
		log.Fatal("Network is down... or you forgot VPN")
	}

	remote, err := exec.Command("git", "remote", "-v").Output()
	if err != nil {
		log.Fatal("Go inside git repo!")
	}

	url := GetUrl(string(remote))

	exec.Command("open", url).Output()

	log.Println("(╯°□°）╯︵ ┻━┻ ", url)
}

// ¯\_(ツ)_/¯
func GetUrl(url string) string {
	r, _ := regexp.Compile(`origin\s*http([a-z]://.+)*.git`)

	switch r.MatchString(url) {
	case true:
		url = strings.Replace(r.FindString(url), ".git", "", 1)
	default:
		url = strings.Replace(url, ":", "/", -1)
		url = strings.Replace(url, `git@`, "https://", 1)
		r, _ := regexp.Compile(`.git\s*\(([a-z]+)\)`)
		url = r.ReplaceAllString(url, "")
	}

	url = strings.Replace(url, "origin", "", 1)
	url = strings.TrimSpace(url)

	return url
}
