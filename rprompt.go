package main

import (
	"os"
	"os/exec"
	"strings"
)

func main() {
	stmp := s
	s = ""

	defer func() {
		recover()
		os.Stdout.WriteString(s)
	}()

	//////////
	// EXIT //
	//////////

	if len(os.Args) > 1 && os.Args[1] != "0" {
		os.Stdout.WriteString(ExitStatus + " " + os.Args[1] + " " + Reset)
		s = stmp
	}

	/////////
	// GIT //
	/////////

	color := GitUnmodified

	// Use "git status" for checks as go-git is very slow here:
	cmd := exec.Command("git", "status")
	cmd.Env = []string{"LANG=C"}
	output, err := cmd.Output()
	split := strings.SplitN(string(output), "\n", 4)
	splitFirstLine := strings.Split(split[0], " ")
	branch := splitFirstLine[len(splitFirstLine)-1]
	statusLine := ""
	if err != nil {
		return
	}
	if !strings.HasSuffix(string(output), "working tree clean\n") {
		color = GitModified
	}
	if len(split) > 1 {
		statusLine = split[1]
	}

	os.Stdout.WriteString(color + " " + branch + " " + Reset)

	if strings.HasPrefix(statusLine, "Your branch is up to date") {
		os.Stdout.WriteString(GitRemoteMatch + " = " + Reset)
	} else if strings.HasPrefix(statusLine, "Your branch is ahead of") { // ' by <n>
		n := strings.Split(strings.Split(statusLine, "' by ")[1], " ")[0]
		os.Stdout.WriteString(GitRemotePush + " +" + n + " " + Reset)
	} else if strings.HasPrefix(statusLine, "Your branch is behind") { // ' by <m>
		m := strings.Split(strings.Split(statusLine, "' by ")[1], " ")[0]
		os.Stdout.WriteString(GitRemotePull + " -" + m + " " + Reset)
	} else if strings.HasSuffix(statusLine, "have diverged,") { // \nand have <n> and <m> different commits
		divergeLine := split[2]
		n := strings.Split(strings.TrimPrefix(divergeLine, "and have "), " ")[0]
		m := strings.Split(strings.Split(divergeLine, " and ")[1], " ")[0]
		os.Stdout.WriteString(GitRemoteDiverge + " +" + n + " -" + m + " " + Reset)
	}
}
