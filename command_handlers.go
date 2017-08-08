package cqrs

import "errors"

type CommandHandler interface {
	HandleCommand(Command) error
}

type CommandHandlerFunc func(Command) error

func (f CommandHandlerFunc) HandleCommand(cmd Command) error {
	return f(cmd)
}

var CommandHandlerAlreadyRegistered = errors.New("Command handler already registered")
