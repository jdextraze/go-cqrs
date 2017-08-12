package example

import (
	"errors"
	"github.com/jdextraze/go-cqrs"
	"github.com/satori/go.uuid"
)

type InventoryItem struct {
	*cqrs.AggregateRoot
	name    string
	created bool
}

func NewInventoryItem(id uuid.UUID) *InventoryItem {
	obj := &InventoryItem{}
	obj.AggregateRoot = cqrs.NewAggregateRoot("InventoryItem", id, obj.apply)
	return obj
}

func (i *InventoryItem) apply(e cqrs.Event) {
	switch e.(type) {
	case *InventoryItemCreated:
		i.onCreated(e.(*InventoryItemCreated))
	}
}

func (i *InventoryItem) Create(name string) error {
	if i.created {
		return errors.New("Already created")
	}
	return i.Apply(&InventoryItemCreated{
		Name: name,
	})
}

func (i *InventoryItem) onCreated(e *InventoryItemCreated) {
	i.created = true
}
