package main

import (
	"github.com/jdextraze/go-cqrs"
	"github.com/jdextraze/go-cqrs/providers/inmemory"
	. "github.com/jdextraze/go-cqrs/example"
	"github.com/satori/go.uuid"
	"log"
)

func main() {
	factory := &Factory{}
	bus := inmemory.NewBus(factory, factory)
	commandsHandler := NewCommandsHandler(cqrs.NewAggregateRepository(inmemory.NewEventStore(bus)))
	eventsHandler := &EventsHandler{}

	bus.RegisterCommandHandler(&CreateInventoryItem{}, commandsHandler)
	bus.RegisterEventHandler(&InventoryItemCreated{}, eventsHandler)

	id := uuid.NewV4()
	if err := bus.SendCommand(&CreateInventoryItem{
		InventoryItemId: id,
		Name:            "Test",
	}); err != nil {
		log.Printf("%v", err)
	}
	if err := bus.SendCommand(&CreateInventoryItem{
		InventoryItemId: id,
		Name:            "Test",
	}); err != nil {
		log.Printf("%v", err)
	}
}
