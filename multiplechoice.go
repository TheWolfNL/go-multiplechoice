// Package MultipleChoice is intended for use in CLI applications
package MultipleChoice

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

// Selection allows the program to ask a question
// and give the user several options to chose from,
// then returning the selected option.
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
