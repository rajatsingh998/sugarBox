package Models

import (
	"sync"
)

type Event struct {
	id           string
	eventType    string
	creationTime int64
	userID       int
	rwMutex      sync.RWMutex
}

func (e *Event) GetID() string {
	defer e.rwMutex.RUnlock()
	e.rwMutex.RLock()
	return e.id
}

func (e *Event) GetUserID() int {
	defer e.rwMutex.RUnlock()
	e.rwMutex.RLock()
	return e.userID
}
func (e *Event) GetEventType() string {
	defer e.rwMutex.RUnlock()
	e.rwMutex.RLock()
	return e.eventType
}

func (e *Event) GetCreationTime() int64 {
	defer e.rwMutex.RUnlock()
	e.rwMutex.RLock()
	return e.creationTime
}

func (e *Event) SetID(id string) {
	defer e.rwMutex.RUnlock()
	e.rwMutex.RLock()
	e.id = id
}

func (e *Event) SetUserID(userID int) {
	defer e.rwMutex.RUnlock()
	e.rwMutex.RLock()
	e.userID = userID
}

func (e *Event) SetEventType(eventType string) {
	defer e.rwMutex.RUnlock()
	e.rwMutex.RLock()
	e.eventType = eventType
}

func (e *Event) SetCreationTime(creationTime int64) {
	defer e.rwMutex.RUnlock()
	e.rwMutex.RLock()
	e.creationTime = creationTime
}
