package main

import (
	"os"
	"os/exec"
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
	home, _ := os.LookupEnv("HOME")
	euid := os.Geteuid()

	hostname, _ := os.Hostname()
	username, _ := exec.Command("/usr/bin/id", "-un").Output()

	os.Stdout.WriteString(Host + strings.TrimSpace(string(username)) + "@" + hostname + " ")

	///////////////////////
	// WORKING DIRECTORY //
	///////////////////////

	dir, err := os.Getwd()
	d := PathFull

	if err != nil { // Can't read directory (permissions, or doesn't exist)

		d = PathError + " ? " + Reset

	} else {

		// Replace home path with ~
		if home != "" && strings.HasPrefix(dir, home) { // current user
			dir = "~/" + strings.TrimPrefix(strings.TrimPrefix(dir, home), "/")
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
