package cqrs

type CommandSender interface {
	SendCommand(cmd Command) error
}

type EventPublisher interface {
	PublishEvent(event *DomainEvent) error
}
