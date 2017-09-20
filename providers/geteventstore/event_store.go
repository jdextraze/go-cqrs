package geteventstore

import (
	"encoding/json"
	"fmt"
	"github.com/jdextraze/go-cqrs"
	"github.com/jdextraze/go-gesclient/client"
	"github.com/satori/go.uuid"
)

const (
	eventsReadBatchSize = 100
)

type eventStore struct {
	eventStoreClient client.Connection
	eventFactory     cqrs.EventFactory
}

func NewEventStore(
	eventStoreClient client.Connection,
	eventFactory cqrs.EventFactory,
) *eventStore {
	return &eventStore{
		eventStoreClient: eventStoreClient,
		eventFactory:     eventFactory,
	}
}

func (r *eventStore) GetEvents(aggregateType string, aggregateId uuid.UUID) ([]*cqrs.DomainEvent, error) {
	domainEvents := []*cqrs.DomainEvent{}
	stream := fmt.Sprintf("%s-%s", aggregateType, aggregateId)
	start := 0

	for {
		task, err := r.eventStoreClient.ReadStreamEventsForwardAsync(stream, start, eventsReadBatchSize, true, nil)
		if err != nil {
			return nil, err
		} else if err := task.Error(); err != nil {
			return nil, err
		}
		slice := task.Result().(*client.StreamEventsSlice)

		resolvedEvents := slice.Events()
		nbEvents := len(resolvedEvents)
		for _, resolvedEvent := range resolvedEvents {
			recordedEvent := resolvedEvent.Event()
			evt, err := r.eventFactory.CreateEvent(recordedEvent.EventType())
			if err != nil {
				return nil, err
			}
			if err := json.Unmarshal(recordedEvent.Data(), evt); err != nil {
				return nil, err
			}

			domainEvents = append(domainEvents, cqrs.NewDomainEvent(
				aggregateId,
				recordedEvent.EventId(),
				recordedEvent.EventNumber(),
				evt,
			))
		}

		if nbEvents < eventsReadBatchSize {
			break
		}
		start += eventsReadBatchSize
	}

	return domainEvents, nil
}

func (r *eventStore) SaveEvents(aggregateType string, aggregateId uuid.UUID, domainEvents []*cqrs.DomainEvent) error {
	stream := fmt.Sprintf("%s-%s", aggregateType, aggregateId)

	domainEventsLength := len(domainEvents)
	if domainEventsLength == 0 {
		return nil
	}

	expectedVersion := int(domainEvents[0].Version()) - 1
	events := make([]*client.EventData, domainEventsLength)
	for i, domainEvent := range domainEvents {
		eventType, err := r.eventFactory.GetEventType(domainEvent.Event())
		if err != nil {
			return err
		}
		data, err := json.Marshal(domainEvent.Event())
		if err != nil {
			return err
		}
		events[i] = client.NewEventData(domainEvent.Id(), eventType, true, data, []byte(""))
	}

	task, err := r.eventStoreClient.AppendToStreamAsync(stream, expectedVersion, events, nil)
	if err != nil {
		return err
	} else if err := task.Error(); err != nil {
		return err
	}

	return err
}
