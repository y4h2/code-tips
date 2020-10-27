package eventstore

import "event-sourcing/model"

type EventStore struct {
	events []model.Event
}

func (es *EventStore) StoreEvent(events []model.Event) error {
	for _, evt := range events {
		evt.ID = len(es.events)
		es.events = append(es.events, evt)
	}

	return nil
}

func (es *EventStore) ListEvents(entity, entityID string) (events []model.Event) {
	for _, evt := range es.events {
		if evt.Entity == entity &&
			evt.EntityID == entityID {

			events = append(events, evt)
		}
	}
	return events
}

func NewEventStore() *EventStore {
	return &EventStore{
		events: []model.Event{},
	}
}
