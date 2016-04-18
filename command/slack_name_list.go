package command

import (
	"strings"
)

type Slack_name_listCommand struct {
	Meta
}

func (c *Slack_name_listCommand) Run(args []string) int {
	// Write your code here

	return 0
}

func (c *Slack_name_listCommand) Synopsis() string {
	return ""
}

func (c *Slack_name_listCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
