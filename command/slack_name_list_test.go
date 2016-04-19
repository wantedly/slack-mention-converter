package command

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestSlackNameListCommand_implement(t *testing.T) {
	var _ cli.Command = &SlackNameListCommand{}
}
