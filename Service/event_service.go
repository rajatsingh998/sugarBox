package Service

import (
	"errors"
	"sugarBox/DataHub"
	"sugarBox/Models"
	"sugarBox/Utility"
)

const (
	PostEventType         string = "post"
	LikeReceivedEventType        = "likeReceived"
	CommentEventType             = "comment"
)

func CreateEvent(userID int, eventType string, creationTime int64) (string, error) {
	var (
		err   error
		event *Models.Event
		user  *Models.User
		ok    bool
	)

	if !createEventReqValidation(eventType) {
		err = errors.New("Validation Failure. Wrong Event Type ")
		return "", err
	}

	if ok, user = getUserWithID(userID); !ok {
		if user, err = createUser(userID); err != nil {
			err = errors.New("User Creation Failed ")
			return "", err
		}
	}

	ok, eventID := isDuplicateEvent(eventType, creationTime, userID)
	if ok {
		err = errors.New("Duplicate Event Details ")
		return eventID, err
	}

	event = createNewEvent(eventType, creationTime, userID)

	if _, err = addEventToUser(user, event); err != nil {
		err = errors.New("Encountered Error While Adding Event To The User ")
		return "", err
	}

	return event.GetID(), nil
}
func isDuplicateEvent(eventType string, creationTime int64, userID int) (bool, string) {
	events := DataHub.GetEventDataSetInstance().GetEvents()

	for _, event := range events {
		if event.GetEventType() == eventType && event.GetUserID() == userID && event.GetCreationTime() == creationTime {
			return true, event.GetID()
		}
	}
	return false, ""
}
func createNewEvent(eventType string, creationTime int64, userID int) *Models.Event {
	newEvent := Models.Event{}
	newEventID, _ := Utility.GenerateUUID()
	newEvent.SetID(newEventID)
	newEvent.SetEventType(eventType)
	newEvent.SetCreationTime(creationTime)
	newEvent.SetUserID(userID)
	eventDataSetInstance.SetEvent(&newEvent)

	return &newEvent
}

func createEventReqValidation(eventType string) bool {
	switch eventType {
	case PostEventType:
		fallthrough
	case CommentEventType:
		fallthrough
	case LikeReceivedEventType:
		return true
	default:
		return false
	}
}

func getEventByID(eventID string) (*Models.Event, bool) {
	events := DataHub.GetEventDataSetInstance().GetEvents()

	for _, event := range events {
		if event.GetID() == eventID {
			return event, true
		}
	}
	return &Models.Event{}, false
}
