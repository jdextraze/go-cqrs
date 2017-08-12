package example

import (
	"errors"
	"github.com/jdextraze/go-cqrs"
)

type Factory struct{}

func (f *Factory) CreateCommand(name string) (cqrs.Command, error) {
	return nil, errors.New("Not implemented")
}

func (f *Factory) GetCommandType(cmd cqrs.Command) (string, error) {
	switch cmd.(type) {
	case *CreateInventoryItem:
		return "CreateInventoryItem", nil
	}
	return "", nil
}

func (f *Factory) CreateEvent(name string) (cqrs.Event, error) {
	switch name {
	case "InventoryItemCreated":
		return &InventoryItemCreated{}, nil
	}
	return nil, nil
}

func (f *Factory) GetEventType(evt cqrs.Event) (string, error) {
	switch evt.(type) {
	case *InventoryItemCreated:
		return "InventoryItemCreated", nil
	}
	return "", nil
}
