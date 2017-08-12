package example

import (
	"github.com/jdextraze/go-cqrs"
	"log"
)

type EventsHandler struct{}

func (h *EventsHandler) HandleEvent(e *cqrs.DomainEvent) error {
	evt := e.Event()
	switch evt.(type) {
	case *InventoryItemCreated:
		return h.handleCreated(evt.(*InventoryItemCreated))
	}
	return nil
}

func (h *EventsHandler) handleCreated(e *InventoryItemCreated) error {
	log.Printf("InventoryItemCreated: %v", e)
	return nil
}
