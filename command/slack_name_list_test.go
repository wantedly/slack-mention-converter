package command

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestSlack_name_listCommand_implement(t *testing.T) {
	var _ cli.Command = &Slack_name_listCommand{}
}
