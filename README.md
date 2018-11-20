# prompt

A really fast command line prompt written in Go.

![](https://raw.githubusercontent.com/moqmar/prompt/master/screenshot.png)

## Setup

```bash
sudo mkdir /opt/prompt && sudo chown $(id -u):$(id -g) /opt/prompt
git clone https://github.com/moqmar/prompt.git /opt/prompt

# Additional steps for development
cd /opt/prompt
go get -u "gopkg.in/src-d/go-git.v4/..."
CGO_ENABLED=0 go build -ldflags '-s -w' prompt.go colors.go
CGO_ENABLED=0 go build -ldflags '-s -w' rprompt.go colors.go
```

### bash (add to `/etc/bash.bashrc` or `~/.bashrc`)
```bash
# Small version without git and exit status
PS1='$(cs="\[" ce="\]" /opt/prompt/prompt)'

# Extended version, adds the rprompt at the front (as bash doesn't support a prompt on the right side)
PS1='$(cs="\[" ce="\]" s=" " /opt/prompt/rprompt "$?")$(cs="\[" ce="\]" /opt/prompt/prompt)'
```

### zsh (add to `~/.zshrc`)
```bash
PROMPT='$(cs="%{" ce="%}" /opt/prompt/prompt)'
RPROMPT='$(cs="%{" ce="%}" /opt/prompt/rprompt "$?")'
```

### fish
```bash
# add to ~/.config/fish/functions/fish_prompt.fish
function fish_prompt
    /opt/prompt/prompt
end

# add to ~/.config/fish/functions/fish_right_prompt.fish
function fish_right_prompt
    /opt/prompt/rprompt "$status"
end
```

## Environment

| Variable | Meaning |
| -------- | ------- |
| `cs`     | Color start marker, for text that shouldn't take up any width. Required by bash and zsh for correct cursor handling. |
| `ce`     | Color end marker (see `cs`) |
| `s`      | Suffix for `rprompt` - will only be appended if the result isn't empty. |

Colors can be adjusted in `colors.go`. See https://stackoverflow.com/a/33206814 for more information on color codes.
