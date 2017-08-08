package cqrs

import (
	"github.com/satori/go.uuid"
)

type ApplyEventHandler func(Event)

type AggregateRoot struct {
	name    string
	id      uuid.UUID
	version int
	apply   ApplyEventHandler
	changes []*DomainEvent
}

func NewAggregateRoot(name string, id uuid.UUID, apply ApplyEventHandler) *AggregateRoot {
	return &AggregateRoot{
		name: name,
		id: id,
		apply: apply,
		version: -1,
		changes: []*DomainEvent{},
	}
}

func (a *AggregateRoot) Name() string { return a.name }

func (a *AggregateRoot) Id() uuid.UUID { return a.id }

func (a *AggregateRoot) Version() int { return a.version }

func (a *AggregateRoot) LoadHistory(events []*DomainEvent) {
	for _, e := range events {
		a.apply(e.event)
		a.version = e.version
	}
}

func (a *AggregateRoot) ApplyChange(event Event) {
	a.apply(event)
	a.version += 1
	a.changes = append(a.changes, &DomainEvent{
		aggregateId: a.id,
		id: uuid.NewV4(),
		version: a.version,
		event: event,
	})
}

func (a *AggregateRoot) GetUncommittedChanges() []*DomainEvent {
	return a.changes
}

func (a *AggregateRoot) MarkChangesAsCommitted() {
	a.changes = []*DomainEvent{}
}
