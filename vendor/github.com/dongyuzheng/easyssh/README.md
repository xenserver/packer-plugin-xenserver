# easyssh

This fork of easyssh includes useful changes from other forks.

# Description

Package easyssh provides a simple implementation of some SSH protocol features in Go.
You can simply run command on remote server or upload a file even simple than native console SSH client.
Do not need to think about Dials, sessions, defers and public keys... Let easyssh will be think about it!

# Install

```go
go get github.com/dongyuzheng/easyssh
```

# Parameters

```go
type MakeConfig struct {
	// Required
	Server   string // Remote machine address (required)
	User     string // Remote user name (required)
	Password string // Password to user
	Key      string // Path to private key on local machine

	// Optional
	Port          string       // SSH port, default 22
	EnablePTY     bool         // PTY session, default false
	OutputHandler func(string) // Line-by-line handler for output, default does nothing
}
```

# Examples

### 1. [Execute a command with live output, and get output](examples/exec.go)
### 2. [Upload a file](examples/upload.go)
