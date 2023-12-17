package DataHub

import (
	"sugarBox/Models"
	"sync"
)

var (
	eventDataSetInstance EventDataSet
)

type EventDataSet struct {
	events  []*Models.Event
	rwMutex sync.RWMutex
}

func (ed *EventDataSet) GetEvents() []*Models.Event {
	defer ed.rwMutex.RUnlock()
	ed.rwMutex.RLock()
	return ed.events
}

func (ed *EventDataSet) SetEvent(event *Models.Event) {
	defer ed.rwMutex.RUnlock()
	ed.rwMutex.RLock()
	ed.events = append(ed.events, event)
}

func initEventDataSet() {
	eventDataSetInstance = EventDataSet{
		events:  make([]*Models.Event, 0),
		rwMutex: sync.RWMutex{},
	}
}

func GetEventDataSetInstance() *EventDataSet {
	once.Do(initUserDataSet)
	return &eventDataSetInstance
}
