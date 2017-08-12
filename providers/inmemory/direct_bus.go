package inmemory

import (
	"github.com/jdextraze/go-cqrs"
)

type bus struct {
	commandFactory  cqrs.CommandFactory
	commandHandlers map[string]cqrs.CommandHandler
	eventFactory    cqrs.EventFactory
	eventHandlers   map[string][]cqrs.EventHandler
}

func NewBus(
	commandFactory cqrs.CommandFactory,
	eventFactory cqrs.EventFactory,
) *bus {
	return &bus{
		commandFactory:  commandFactory,
		commandHandlers: map[string]cqrs.CommandHandler{},
		eventFactory:    eventFactory,
		eventHandlers:   map[string][]cqrs.EventHandler{},
	}
}

func (b *bus) RegisterCommandHandler(
	cmd cqrs.Command,
	handler cqrs.CommandHandler,
) error {
	name, err := b.commandFactory.GetCommandType(cmd)
	if err != nil {
		return err
	}
	if _, found := b.commandHandlers[name]; found {
		return cqrs.CommandHandlerAlreadyRegistered
	}
	b.commandHandlers[name] = handler
	return nil
}

func (b *bus) RegisterEventHandler(
	evt cqrs.Event,
	handler cqrs.EventHandler,
) error {
	name, err := b.eventFactory.GetEventType(evt)
	if err != nil {
		return err
	}
	if _, found := b.eventHandlers[name]; !found {
		b.eventHandlers[name] = []cqrs.EventHandler{}
	}
	b.eventHandlers[name] = append(b.eventHandlers[name], handler)
	return nil
}

func (b *bus) SendCommand(cmd cqrs.Command) error {
	name, err := b.commandFactory.GetCommandType(cmd)
	if err != nil {
		return err
	}
	handler := b.commandHandlers[name]
	return handler.HandleCommand(cmd)
}

func (b *bus) PublishEvent(evt *cqrs.DomainEvent) error {
	name, err := b.eventFactory.GetEventType(evt.Event())
	if err != nil {
		return err
	}
	handlers := b.eventHandlers[name]
	for _, h := range handlers {
		if err := h.HandleEvent(evt); err != nil {
			return err
		}
	}
	return nil
}
