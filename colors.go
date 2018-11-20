package main

import (
	"os"
)

var cs, _ = os.LookupEnv("cs")
var ce, _ = os.LookupEnv("ce")
var s, _ = os.LookupEnv("s")

var (
	ExitStatus       = cs + "\033[1;41m" + ce
	GitUnmodified    = cs + "\033[1;100m" + ce
	GitModified      = cs + "\033[1;33;100m" + ce
	GitRemoteMatch   = cs + "\033[0;1;42m" + ce
	GitRemotePush    = cs + "\033[0;1;43m" + ce
	GitRemotePull    = cs + "\033[0;1;45m" + ce
	GitRemoteDiverge = cs + "\033[0;1;41m" + ce

	Time = cs + "\033[1;35m" + ce

	Host = cs + "\033[1;30m" + ce

	PathError = cs + "\033[1;41m" + ce
	PathFull  = cs + "\033[1;33m" + ce
	PathShort = cs + "\033[0;33m" + ce

	PromptUser = cs + "\033[1;34m" + ce
	PromptRoot = cs + "\033[1;31m" + ce

	Reset = cs + "\033[0m" + ce
)
