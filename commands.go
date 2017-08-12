package cqrs

import (
	"errors"
	"github.com/satori/go.uuid"
)

type Command interface {
	AggregateId() uuid.UUID
}

type CommandFactory interface {
	CreateCommand(name string) (Command, error)
	GetCommandType(Command) (string, error)
}

type CommandSender interface {
	SendCommand(cmd Command) error
}

type CommandHandler interface {
	HandleCommand(Command) error
}

type CommandHandlerFunc func(Command) error

func (f CommandHandlerFunc) HandleCommand(cmd Command) error {
	return f(cmd)
}

var CommandHandlerAlreadyRegistered = errors.New("Command handler already registered")
