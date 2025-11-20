package inp

import (
	"errors"
	"os"
	"os/exec"
	"strings"
)

var ErrNoPath = errors.New("path required \n")

func ExecInput(input string) error {
	// Remove the newline character
	input = strings.TrimSuffix(input, "\n")

	// Split the input to separet the command and the arguements
	args := strings.Split(input, " ")

	switch args[0] {
	case "cd":
		if len(args[1:]) == 0 {
			return ErrNoPath
		}

		return os.Chdir(args[1])
	case "exit":
		os.Exit(0)
	}

	// Pass the program and the arguments separately
	cmd := exec.Command(args[0], args[1:]...)

	// Set the correct output device
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	// Execute the command and return the error if there's any
	return cmd.Run()
}
