package command

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestTo_slack_nameCommand_implement(t *testing.T) {
	var _ cli.Command = &To_slack_nameCommand{}
}
