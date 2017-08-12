package example

import "github.com/satori/go.uuid"

type CreateInventoryItem struct {
	InventoryItemId uuid.UUID
	Name            string
}

func (c *CreateInventoryItem) AggregateId() uuid.UUID { return c.InventoryItemId }
