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

func CreateEvent(userID string, eventType string, creationTime int) (string, error) {
	var (
		err       error
		event     Models.Event
		eventDate time.Time
		ok        bool
	)

	if !createEventReqValidation(eventType) {
		err = errors.New("Validation Failure. Wrong Event Type ")
		return "", err
	}

	if ok, _ = getUserWithID(userID); !ok {
		if _, err = createUser(userID); err != nil {
			err = errors.New("User Creation Failed ")
			return "", err
		}
	}

	eventDate, err = Utility.ConvertTimestampTo_YYYY_MM_DD(creationTime)

	if err != nil {
		return "", err
	}

	event = createNewEvent(eventType, eventDate)

	if _, err = addEventToUser(userID, event.GetID()); err != nil {
		err = errors.New("Encountered Error While Adding Event To The User ")
		return "", err
	}

	return event.GetID(), nil
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
