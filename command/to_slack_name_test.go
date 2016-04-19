package command

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestToSlackNameCommand_implement(t *testing.T) {
	var _ cli.Command = &ToSlackNameCommand{}
}
