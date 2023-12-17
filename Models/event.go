package Models

import (
	"sync"
	"time"
)

type Event struct {
	id           string
	eventType    string
	creationTime time.Time
	rwMutex      sync.RWMutex
}

func (e *Event) GetID() string {
	defer e.rwMutex.RUnlock()
	e.rwMutex.RLock()
	return e.id
}

func (e *Event) GetEventType() string {
	defer e.rwMutex.RUnlock()
	e.rwMutex.RLock()
	return e.eventType
}

func (e *Event) GetCreationTime() time.Time {
	defer e.rwMutex.RUnlock()
	e.rwMutex.RLock()
	return e.creationTime
}

func (e *Event) SetID(id string) {
	defer e.rwMutex.RUnlock()
	e.rwMutex.RLock()
	e.id = id
}

func (e *Event) SetEventType(eventType string) {
	defer e.rwMutex.RUnlock()
	e.rwMutex.RLock()
	e.eventType = eventType
}

func (e *Event) SetCreationTime(creationTime time.Time) {
	defer e.rwMutex.RUnlock()
	e.rwMutex.RLock()
	e.creationTime = creationTime
}
