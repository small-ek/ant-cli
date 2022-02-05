package aexec

import (
	"bytes"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

type Process struct {
	exec.Cmd
	PPid int
}

func ShellExec(cmd string, environment ...[]string) (result string, err error) {
	var (
		buf = bytes.NewBuffer(nil)
		p   = NewProcess(
			getShell(),
			append([]string{getShellOption()}, parseCommand(cmd)...),
			environment...,
		)
	)
	p.Stdout = buf
	p.Stderr = buf
	err = p.Run()
	result = buf.String()
	return
}

func parseCommand(cmd string) (args []string) {
	if runtime.GOOS != "windows" {
		return []string{cmd}
	}
	// Just for "cmd.exe" in windows.
	var argStr string
	var firstChar, prevChar, lastChar1, lastChar2 byte
	array := strings.Split(cmd, " ")
	for _, v := range array {
		if len(argStr) > 0 {
			argStr += " "
		}
		firstChar = v[0]
		lastChar1 = v[len(v)-1]
		lastChar2 = 0
		if len(v) > 1 {
			lastChar2 = v[len(v)-2]
		}
		if prevChar == 0 && (firstChar == '"' || firstChar == '\'') {
			// It should remove the first quote char.
			argStr += v[1:]
			prevChar = firstChar
		} else if prevChar != 0 && lastChar2 != '\\' && lastChar1 == prevChar {
			// It should remove the last quote char.
			argStr += v[:len(v)-1]
			args = append(args, argStr)
			argStr = ""
			prevChar = 0
		} else if len(argStr) > 0 {
			argStr += v
		} else {
			args = append(args, v)
		}
	}
	return
}

// getShellOption returns the shell option depending on current working operating system.
// It returns "/c" for windows, and "-c" for others.
func getShellOption() string {
	switch runtime.GOOS {
	case "windows":
		return "/c"
	default:
		return "-c"
	}
}

func NewProcess(path string, args []string, environment ...[]string) *Process {
	env := os.Environ()
	if len(environment) > 0 {
		env = append(env, environment[0]...)
	}
	process := &Process{
		PPid: os.Getpid(),
		Cmd: exec.Cmd{
			Args:       []string{path},
			Path:       path,
			Stdin:      os.Stdin,
			Stdout:     os.Stdout,
			Stderr:     os.Stderr,
			Env:        env,
			ExtraFiles: make([]*os.File, 0),
		},
	}
	process.Dir, _ = os.Getwd()
	if len(args) > 0 {
		// Exclude of current binary path.
		start := 0
		if strings.EqualFold(path, args[0]) {
			start = 1
		}
		process.Args = append(process.Args, args[start:]...)
	}
	return process
}

// getShell returns the shell command depending on current working operating system.
// It returns "cmd.exe" for windows, and "bash" or "sh" for others.
func getShell() string {
	switch runtime.GOOS {
	case "windows":
		return SearchBinary("cmd.exe")
	default:
		// Check the default binary storage path.
		if Exists("/bin/bash") {
			return "/bin/bash"
		}
		if Exists("/bin/sh") {
			return "/bin/sh"
		}
		// Else search the env PATH.
		path := SearchBinary("bash")
		if path == "" {
			path = SearchBinary("sh")
		}
		return path
	}
}

// SearchBinary searches the binary `file` in current working folder and PATH environment.
func SearchBinary(file string) string {
	// Check if it is absolute path of exists at current working directory.
	if Exists(file) {
		return file
	}
	return SearchBinaryPath(file)
}

// Exists checks whether given `path` exist.
func Exists(path string) bool {
	if stat, err := os.Stat(path); stat != nil && !os.IsNotExist(err) {
		return true
	}
	return false
}

func SearchBinaryPath(file string) string {
	array := ([]string)(nil)
	switch runtime.GOOS {
	case "windows":
		envPath, _ := os.LookupEnv("PATH")
		if strings.Contains(envPath, ";") {
			array = strings.Split(envPath, ";")
		} else if strings.Contains(envPath, ":") {
			array = strings.Split(envPath, ":")
		}
		if Ext(file) != ".exe" {
			file += ".exe"
		}
	default:
		v, _ := os.LookupEnv("PATH")
		array = strings.Split(v, ":")
	}
	if len(array) > 0 {
		path := ""
		for _, v := range array {
			path = v + string(os.PathSeparator) + file
			if Exists(path) && IsFile(path) {
				return path
			}
		}
	}
	log.Println(array)
	return ""
}

// IsFile checks whether given `path` a file, which means it's not a directory.
// Note that it returns false if the `path` does not exist.
func IsFile(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !s.IsDir()
}

func Ext(path string) string {
	ext := filepath.Ext(path)
	if p := strings.IndexByte(ext, '?'); p != -1 {
		ext = ext[0:p]
	}
	return ext
}
