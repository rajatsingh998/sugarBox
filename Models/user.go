package Models

import "sync"

type User struct {
	id       string
	eventIds []string
	rwMutex  sync.RWMutex
}

func (u *User) GetID() string {
	defer u.rwMutex.RUnlock()
	u.rwMutex.RLock()
	return u.id
}

func (u *User) GetEventIDs() []string {
	defer u.rwMutex.RUnlock()
	u.rwMutex.RLock()
	return u.eventIds
}

func (u *User) SetID(id string) {
	defer u.rwMutex.RUnlock()
	u.rwMutex.RLock()
	u.id = id
}

func (u *User) SetEventID(eventID string) {
	defer u.rwMutex.RUnlock()
	u.rwMutex.RLock()
	u.eventIds = append(u.eventIds, eventID)
}
