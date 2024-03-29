package Models

import "sync"

type User struct {
	id                int
	eventIds          []string
	userEventCounters []*UserEventCounter
	rwMutex           sync.RWMutex
}

func (u *User) GetID() int {
	defer u.rwMutex.RUnlock()
	u.rwMutex.RLock()
	return u.id
}

func (u *User) GetUserEventCounters() []*UserEventCounter {
	defer u.rwMutex.RUnlock()
	u.rwMutex.RLock()
	return u.userEventCounters
}

func (u *User) GetEventIDs() []string {
	defer u.rwMutex.RUnlock()
	u.rwMutex.RLock()
	return u.eventIds
}

func (u *User) SetUserEventCounter(userEventCounter *UserEventCounter) {
	defer u.rwMutex.RUnlock()
	u.rwMutex.RLock()
	u.userEventCounters = append(u.userEventCounters, userEventCounter)
}
func (u *User) SetID(id int) {
	defer u.rwMutex.RUnlock()
	u.rwMutex.RLock()
	u.id = id
}

func (u *User) SetEventID(eventID string) {
	defer u.rwMutex.RUnlock()
	u.rwMutex.RLock()
	u.eventIds = append(u.eventIds, eventID)
}
