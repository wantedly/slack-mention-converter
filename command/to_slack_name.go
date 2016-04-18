package command

import (
	"strings"
)

type To_slack_nameCommand struct {
	Meta
}

func (c *To_slack_nameCommand) Run(args []string) int {
	// Write your code here

	return 0
}

func (c *To_slack_nameCommand) Synopsis() string {
	return ""
}

func (c *To_slack_nameCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
