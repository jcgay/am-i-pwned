package pwned

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

var logger = log.New(os.Stderr, "", 0)

type executor func(output io.Writer, command ...string) error

type password struct {
	name  string
	value string
}

type jsonPassword struct {
	Password string `json:"password"`
}

func CheckAllPasswords() {
	passwords := listAllPasswords(ls, show)
	for _, password := range passwords {
		result, err := CheckPassword(password.value)
		if err != nil {
			logger.Printf("Error while verifying %s\n  %s", password.name, err)
		}
		fmt.Printf("%s: %d\n", password.name, result)
	}
}

func listAllPasswords(ls func() (map[string]string, error), show func(id string, exec executor) (string, error)) map[string]password {
	websiteByIds, err := ls()
	if err != nil {
		return map[string]password{}
	}

	result := make(map[string]password)
	for id, name := range websiteByIds {
		output, err := show(id, runCommand)
		if err != nil {
			result[id] = password{name: name}
		} else {
			result[id] = password{name: name, value: output}
		}
	}

	return result
}

func ls() (map[string]string, error) {
	output := new(bytes.Buffer)
	err := runCommand(output, "lpass", "ls")
	if err != nil {
		return nil, err
	}
	return parseManagedPasswords(output.String()), nil
}

func show(id string, exec executor) (string, error) {
	output := new(bytes.Buffer)
	err := exec(output, "lpass", "show", id, "--json")
	if err != nil {
		return "", err
	}

	result := make([]jsonPassword, 0)
	json.Unmarshal(output.Bytes(), &result)

	return result[0].Password, nil
}

func parseManagedPasswords(output string) map[string]string {
	result := make(map[string]string)
	password := regexp.MustCompile(`(.+) \[id: ([0-9]+)]`)

	lines := strings.Split(output, "\n")
	for _, line := range lines {
		matches := password.FindStringSubmatch(strings.TrimSpace(line))
		if matches != nil {
			result[matches[2]] = matches[1]
		}
	}

	return result
}

func runCommand(output io.Writer, command ...string) error {
	toExecute := exec.Command(command[0], command[1:]...)
	toExecute.Stdout = output
	toExecute.Stderr = output
	return toExecute.Run()
}
