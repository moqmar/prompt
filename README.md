# prompt

A really fast command line prompt written in Go.

## Setup
```bash
go get -u "gopkg.in/src-d/go-git.v4/..."
git clone https://github.com/moqmar/prompt.git /opt/prompt
CGO_ENABLED=0 go build -ldflags '-s -w' prompt.go colors.go
CGO_ENABLED=0 go build -ldflags '-s -w' rprompt.go colors.go
```

### bash (`/etc/bash.bashrc` or `~/.bashrc`)
```bash
PS1='$(cs="\[" ce="\]" s=" " /opt/prompt/rprompt "$?")$(cs="\[" ce="\]" /opt/prompt/prompt)'
```

### zsh (`~/.zshrc`)
```bash
PROMPT='$(cs="%{" ce="%}" /opt/prompt/prompt)'
RPROMPT='$(cs="%{" ce="%}" /opt/prompt/rprompt "$?")'
```

### fish
```bash
function fish_prompt
    /opt/prompt/prompt
end
function fish_right_prompt
    /opt/prompt/rprompt
end
```

## Environment

| Variable | Meaning |
| -------- | ------- |
| `cs`     | Color start marker, for text that shouldn't take up any width. Required by bash and zsh for correct cursor handling. |
| `ce`     | Color end marker (see `cs`) |
| `s`      | Suffix for `rprompt` - will only be appended if the result isn't empty. |

Colors can be adjusted in `colors.go`. See https://stackoverflow.com/a/33206814 for more information on color codes.
