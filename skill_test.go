package main

import (
	"bytes"
	"strings"
	"testing"

	skilldoc "github.com/phobologic/repoguide/internal/skill"
)

func TestRunSkill(t *testing.T) {
	t.Parallel()

	var stdout, stderr bytes.Buffer
	if err := runSkill(nil, &stdout, &stderr); err != nil {
		t.Fatalf("runSkill: %v", err)
	}

	got := stdout.String()
	if got != skilldoc.Content()+"\n" && got != skilldoc.Content() {
		t.Fatalf("unexpected skill output:\n%s", got)
	}
	if !strings.Contains(got, "name: repoguide") {
		t.Fatal("missing skill frontmatter")
	}
	if !strings.Contains(got, "## Recommended Workflow") {
		t.Fatal("missing skill body")
	}
	if stderr.Len() != 0 {
		t.Fatalf("unexpected stderr: %q", stderr.String())
	}
}

func TestRunSkillRejectsArgs(t *testing.T) {
	t.Parallel()

	err := runSkill([]string{"extra"}, &bytes.Buffer{}, &bytes.Buffer{})
	if err == nil {
		t.Fatal("expected error for positional arguments")
	}
}

func TestRunDispatchesSkillSubcommand(t *testing.T) {
	t.Parallel()

	var stdout, stderr bytes.Buffer
	if err := run([]string{"skill"}, &stdout, &stderr); err != nil {
		t.Fatalf("run: %v", err)
	}
	if !strings.Contains(stdout.String(), "name: repoguide") {
		t.Fatal("missing skill output")
	}
}
