package MultipleChoice

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

/*
This currently requires the user to have fzf installed

Usage:
	selection := MultipleChoice.Selection("Select one: ", []string{"option1", "option2", "option3"}])
*/
func Selection(question string, options []string) string {
	fmt.Println(question)

	input := func(in io.WriteCloser) {
		for i := 0; i < len(options); i++ {
			fmt.Fprintln(in, options[i])
		}
	}

	shell := os.Getenv("SHELL")
	if len(shell) == 0 {
		shell = "sh"
	}
	cmd := exec.Command(shell, "-c", "fzf -m --height 40% --reverse --border")
	cmd.Stderr = os.Stderr
	in, _ := cmd.StdinPipe()
	go func() {
		input(in)
		in.Close()
	}()
	result, _ := cmd.Output()
	return strings.TrimSpace(string(result))
}
