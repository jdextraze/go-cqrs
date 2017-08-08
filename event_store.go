package cqrs

import (
	"github.com/satori/go.uuid"
	"errors"
)

type EventStore interface {
	GetEvents(aggregateType string, aggregateId uuid.UUID) ([]*DomainEvent, error)
	SaveEvents(aggregateType string, aggregateId uuid.UUID, events []*DomainEvent) error
}

var AggregateNotFound = errors.New("Aggregate not found")
