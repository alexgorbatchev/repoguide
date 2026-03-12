package skill

import _ "embed"

//go:embed repoguide.md
var content string

func Content() string {
	return content
}
