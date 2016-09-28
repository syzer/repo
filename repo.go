package main

import (
	"log"
	"os/exec"
	"regexp"
	"strings"
)

// TODO extract real domain
// ಠ_ಠ
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
	r := regexp.MustCompile(`origin\s*http([a-z]://.+)`)

	switch r.MatchString(url) {
	case true:
		url = r.FindString(url)
		words := strings.Fields(url)
		url = words[1]
	default:
		r := regexp.MustCompile(`^origin(.*)git@(.*)git`)
		url = strings.Replace(r.FindString(url), ":", "/", -1)
		url = strings.Replace(url, `git@`, "https://", 1)
	}
	url = strings.Replace(url, ".git", "", 1)
	url = strings.Replace(url, "origin", "", 1)
	url = strings.TrimSpace(url)

	return url
}
