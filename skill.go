package main

import (
	"flag"
	"fmt"
	"io"

	skilldoc "github.com/phobologic/repoguide/internal/skill"
)

func runSkill(args []string, stdout, stderr io.Writer) error {
	fs := flag.NewFlagSet("repoguide skill", flag.ContinueOnError)
	fs.SetOutput(stderr)

	fs.Usage = func() {
		_, _ = fmt.Fprintf(stderr, `Usage: repoguide skill

Print the bundled OpenCode skill markdown to stdout.

Examples:
  repoguide skill
  repoguide skill > .opencode/skills/repoguide/SKILL.md
`)
	}

	if err := fs.Parse(args); err != nil {
		return err
	}
	if fs.NArg() != 0 {
		return fmt.Errorf("skill takes no positional arguments")
	}

	content := skilldoc.Content()
	_, _ = io.WriteString(stdout, content)
	if len(content) == 0 || content[len(content)-1] != '\n' {
		_, _ = io.WriteString(stdout, "\n")
	}
	return nil
}
