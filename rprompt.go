package main

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"gopkg.in/src-d/go-git.v4/plumbing"

	"gopkg.in/src-d/go-git.v4"
)

func main() {
	stmp := s
	s = ""

	//////////
	// TIME //
	//////////

	if len(os.Args) > 1 && os.Args[1] != "0" {
		os.Stdout.WriteString(ExitStatus + " " + os.Args[1] + " " + Reset)
		s = stmp
	}

	/////////
	// GIT //
	/////////

	dir, err := os.Getwd()
	if err != nil {
		os.Stdout.WriteString(s)
		return
	}
	err = errors.New("")
	var repo *git.Repository

	for err != nil {
		repo, err = git.PlainOpen(dir)
		if err == nil {
			break
		}
		dir = filepath.Dir(dir)
		if dir == "." || dir == "/" {
			os.Stdout.WriteString(s)
			return
		}
	}

	head, err := repo.Head()
	if err != nil {
		return
	}
	color := GitUnmodified

	// Use "git status" for checks as go-git is very slow here:
	cmd := exec.Command("git", "status")
	cmd.Env = []string{"LANG=C"}
	output, err := cmd.Output()
	if err == nil {
		if !strings.HasSuffix(string(output), "working tree clean\n") {
			color = GitModified
		}
	}

	if head.Name().Short() == "HEAD" {
		os.Stdout.WriteString(color + " " + head.Hash().String()[:7] + " " + Reset)
		s = stmp
	} else {
		os.Stdout.WriteString(color + " " + head.Name().Short() + " " + Reset)
		s = stmp

		if !head.Name().IsBranch() {
			os.Stdout.WriteString(s)
			return
		}

		localRef, err := repo.Reference(head.Name(), true)
		if err != nil {
			os.Stdout.WriteString(s)
			return
		}

		branch, err := repo.Branch(head.Name().Short())
		if err != nil {
			os.Stdout.WriteString(s)
			return
		}

		remoteRef, err := repo.Reference(plumbing.NewRemoteReferenceName(branch.Remote, branch.Merge.Short()), true)
		if err != nil {
			os.Stdout.WriteString(s)
			return
		}

		if localRef.Hash().String() == remoteRef.Hash().String() {
			os.Stdout.WriteString(GitRemoteMatch + " = " + Reset)
		} else {
			localCommit, _ := repo.CommitObject(head.Hash())
			for _, c := range localCommit.ParentHashes {
				if c.String() == remoteRef.Hash().String() {
					os.Stdout.WriteString(GitRemotePush + " + " + Reset)
					os.Stdout.WriteString(s)
					return
				}
			}
			remoteCommit, _ := repo.CommitObject(remoteRef.Hash())
			for _, c := range remoteCommit.ParentHashes {
				if c.String() == localRef.Hash().String() {
					os.Stdout.WriteString(GitRemotePull + " - " + Reset)
					os.Stdout.WriteString(s)
					return
				}
			}

			os.Stdout.WriteString(GitRemoteDiverge + " - " + Reset)
		}
	}
	os.Stdout.WriteString(s)
}
