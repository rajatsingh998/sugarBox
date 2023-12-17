package Service

import (
	"errors"
	"sugarBox/Models"
	"sugarBox/Utility"
	"time"
)

const (
	PostEventType         string = "post"
	LikeReceivedEventType        = "likeReceived"
	CommentEventType             = "comment"
)

func CreateEvent(userID string, eventType string, creationTime time.Time) (string, error) {
	var (
		err   error
		event Models.Event
		user  Models.User
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

	event = createNewEvent(eventType, creationTime)

	AddEventToUser(userID, event.GetEventType())

}

func createNewEvent(eventType string, creationTime time.Time) Models.Event {
	newEvent := Models.Event{}
	newEventID, _ := Utility.GenerateUUID()
	newEvent.SetID(newEventID)
	newEvent.SetEventType(eventType)
	newEvent.SetCreationTime(creationTime)
	eventDataSetInstance.SetEvent(&newEvent)
	return newEvent
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
