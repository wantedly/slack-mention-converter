package command

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestTo_slack_mentionCommand_implement(t *testing.T) {
	var _ cli.Command = &To_slack_mentionCommand{}
}
