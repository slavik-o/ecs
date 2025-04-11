package ecs

// Event represents something that happened in the game world
type Event interface {
	Type() EventType
}

// EventType identifies different types of events
type EventType uint32

// EventHandler is a callback function that processes events
type EventHandler func(event Event)

// EventManager handles the registration and distribution of events
type EventManager struct {
	// Handlers map to store all event handlers
	handlers map[EventType][]EventHandler

	// Queue of events to be processed
	queue []Event

	// Pending events to be processed
	pending []Event
}

// NewEventManager creates a new event manager
func NewEventManager() *EventManager {
	return &EventManager{
		handlers: make(map[EventType][]EventHandler),
		queue:    []Event{},
		pending:  []Event{},
	}
}

// Subscribe registers a handler for a specific event type
func (em *EventManager) Subscribe(eventType EventType, handler EventHandler) {
	em.handlers[eventType] = append(em.handlers[eventType], handler)
}

// Publish adds an event to the queue to be processed on the next update
func (em *EventManager) Publish(event Event) {
	em.queue = append(em.queue, event)
}

// PublishImmediate immediately dispatches an event to all subscribed handlers
func (em *EventManager) PublishImmediate(event Event) {
	if handlers, exists := em.handlers[event.Type()]; exists {
		for _, handler := range handlers {
			handler(event)
		}
	}
}

// Update processes all events in the queue
func (em *EventManager) Update() {
	// Swap queue and pending to prevent infinite loops if handlers publish new events
	em.pending, em.queue = em.queue, em.pending
	em.queue = em.queue[:0] // Clear queue but retain capacity

	for _, event := range em.pending {
		if handlers, exists := em.handlers[event.Type()]; exists {
			for _, handler := range handlers {
				handler(event)
			}
		}
	}
}
