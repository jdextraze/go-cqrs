package example

import (
	"github.com/jdextraze/go-cqrs"
)

type CommandsHandler struct {
	repo *cqrs.AggregateRepository
}

func NewCommandsHandler(repo *cqrs.AggregateRepository) *CommandsHandler {
	return &CommandsHandler{repo}
}

func (h *CommandsHandler) HandleCommand(cmd cqrs.Command) (err error) {
	i := NewInventoryItem(cmd.AggregateId())
	h.repo.Load(i.AggregateRoot)
	switch cmd.(type) {
	case *CreateInventoryItem:
		err = h.handleCreate(i, cmd.(*CreateInventoryItem))
	}
	if err != nil {
		return
	}
	err = h.repo.Save(i.AggregateRoot)
	return
}

func (h *CommandsHandler) handleCreate(i *InventoryItem, cmd *CreateInventoryItem) error {
	return i.Create(cmd.Name)
}
