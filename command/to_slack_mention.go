package command

import (
	"strings"
)

type To_slack_mentionCommand struct {
	Meta
}

func (c *To_slack_mentionCommand) Run(args []string) int {
	// Write your code here

	return 0
}

func (c *To_slack_mentionCommand) Synopsis() string {
	return ""
}

func (c *To_slack_mentionCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
