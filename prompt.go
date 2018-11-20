package main

import (
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func main() {

	//////////
	// TIME //
	//////////

	os.Stdout.WriteString(Time + time.Now().Format("15:04:05") + " ")

	/////////////////
	// USER & HOST //
	/////////////////

	//   0      1     2     3
	// [... username uid homedir]
	home, _ := ioutil.ReadFile("/etc/passwd")
	homes := regexp.MustCompile(`(?:^|\n)([^:]+):[^:]+:([^:]+):[^:]+:[^:]+:([^:]+):`).FindAllStringSubmatch(string(home), -1)
	euid := os.Geteuid()
	host := ""
	for _, user := range homes {
		if user[2] == strconv.Itoa(euid) {
			host += user[1]
			break
		}
	}
	if host == "" {
		host += strconv.Itoa(euid)
	}
	hostname, _ := os.Hostname()
	if hostname != "" {
		host += "@" + hostname
	}

	os.Stdout.WriteString(Host + host + " ")

	///////////////////////
	// WORKING DIRECTORY //
	///////////////////////

	dir, err := os.Getwd()
	d := PathFull

	if err != nil { // Can't read directory (permissions, or doesn't exist)

		d = PathError + " ? " + Reset

	} else {

		// Replace home paths with ~user (if shorter)
		for _, user := range homes {
			if user[2] == strconv.Itoa(euid) && strings.HasPrefix(dir, user[3]) { // current user
				dir = "~/" + strings.TrimPrefix(strings.TrimPrefix(dir, user[3]), "/")
				break
			} else if strings.HasPrefix(dir, user[3]) && (user[2] == "0" || len(user[1]) < len(strings.Trim(user[3], "/"))) { // at least one characters saved, or username is root
				dir = "~" + user[1] + "/" + strings.TrimPrefix(strings.TrimPrefix(dir, user[3]), "/")
				break
			}
		}

		// Shorten path elements and color everything
		dirParts := strings.Split(dir, "/")
		for i, part := range dirParts {
			if i == 0 && part != "" { // User
				d += PathShort + part + PathFull + "/"
			} else if i > 0 && i < len(dirParts)-2 && len(part) > 3 { // Keep last two parts and parts with up to three characters full-length
				d += PathShort + part[:2] + "*" + PathFull + "/"
			} else if i < len(dirParts)-1 {
				d += part + "/"
			} else {
				d += part
			}
		}

	}

	os.Stdout.WriteString(d)

	//////////////////////
	// PROMPT INDICATOR //
	//////////////////////

	if euid == 0 {
		os.Stdout.WriteString(" " + PromptRoot + "#" + Reset + " ")
	} else {
		os.Stdout.WriteString(" " + PromptUser + "$" + Reset + " ")
	}

}
