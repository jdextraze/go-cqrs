package cqrs

type AggregateRepository struct {
	storage EventStore
}

func (r *AggregateRepository) Load(a AggregateRoot) error {
	events, err := r.storage.GetEvents(a.name, a.id)
	if err != nil {
		return err
	}
	a.LoadHistory(events)
	return nil
}

func (r *AggregateRepository) Save(a AggregateRoot) error {
	return r.storage.SaveEvents(a.name, a.id, a.GetUncommittedChanges())
}
