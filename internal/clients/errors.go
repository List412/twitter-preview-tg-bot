package clients

import (
	"fmt"
	"github.com/pkg/errors"
	"tweets-tg-bot/internal/commands"
)

type ErrServiceDisabled struct {
	command commands.Cmd
}

func NewErrServiceDisabled(command commands.Cmd) *ErrServiceDisabled {
	return &ErrServiceDisabled{command: command}
}

func (e *ErrServiceDisabled) Error() string {
	return fmt.Sprintf("%s service is disabled right now", e.command)
}

var ErrServiceNotFound = errors.New("service not found")
