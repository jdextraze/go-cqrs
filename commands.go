package cqrs

import "github.com/satori/go.uuid"

type Command interface {
	AggregateId() uuid.UUID
}

type CommandFactory interface {
	CreateCommand(name string) (interface{}, error)
	GetCommandType(interface{}) (string, error)
}
