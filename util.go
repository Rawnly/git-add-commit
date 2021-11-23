package main

import (
	"os"
	"os/exec"
	"strings"
)

func GitDiff() error {
	return RunCommand("git", "diff")
}

func GitAddAll() error {
	return RunCommand("git", "add", "-A", ".")
}

func RunCommand(command string, args ...string) error {
	cmd := exec.Command(command, args...)

	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func GitCommit(message string) error {
	return RunCommand("git", "commit", "-n", "-a", "-m", message)
}

func GitStatus() ([]string, error) {
	b, err := exec.Command("git", "status", "-s").Output()

	if err != nil {
		return nil, err
	}

	out := string(b)

	output := strings.Split(out, "\n")

	return output, nil
}

func Clear() error {
	return RunCommand("clear")
}

func writeFile(filename string, content string) error {
	f, err := os.Create(filename)

	if err != nil {
		return err
	}

	defer f.Close()

	_, err2 := f.WriteString(content)

	if err2 != nil {
		return err
	}

	return nil
}


func OpenEditor(content string) (string, error) {
	const FileName = ".commit"

	if e := writeFile(FileName, content); e != nil {
		return "", e
	}

	err := RunCommand("vim",  FileName)

	if err != nil {
		return content, err
	}

	data, err := os.ReadFile(FileName)

	if err != nil {
		return content, err
	}

	if err = os.Remove(FileName); err != nil {
		return "", err
	}

	return strings.TrimRight(string(data), "\n"), nil
}