package Service

import (
	"errors"
	"sugarBox/DataHub"
	"sugarBox/Models"
	"time"
)

var (
	userDataSetInstance  = DataHub.GetUserDataSetInstance()
	eventDataSetInstance = DataHub.GetEventDataSetInstance()
)

func createUser(id string) (*Models.User, error) {
	newUser := Models.User{}

	newUser.SetID(id)
	userDataSetInstance.SetUsers(&newUser)
	return &newUser, nil
}

func addEventToUser(userId string, event *Models.Event) (string, error) {
	var (
		err  error
		user *Models.User
		ok   bool
	)

	if ok, user = getUserWithID(userId); !ok {
		err = errors.New("User Not Found With Given ID ")
		return "", err
	}

	/*Todo:: We can make updating event counter and setting event id to user in single
	Todo::  transaction to avoid any inconsistency*/
	if ok = updateEventCounterForUser(user, event); !ok {
		err = errors.New("Error Encountered While Updating The Count ")
		return "", err
	}

	user.SetEventID(event.GetID())

	return userId, nil
}

func updateEventCounterForUser(user *Models.User, event *Models.Event) bool {
	var (
		userEventCounter *Models.UserEventCounter
		ok               bool
	)

	userEventCounter, ok = isEventCounterAvailableForTheDate(user, event.GetCreationTime())

	if !ok {
		userEventCounter.SetDate(event.GetCreationTime())
		user.SetUserEventCounter(userEventCounter)
	}

	switch event.GetEventType() {
	case LikeReceivedEventType:
		userEventCounter.IncreaseLikeCountByOne()
	case PostEventType:
		userEventCounter.IncreasePostCountByOne()
	case CommentEventType:
		userEventCounter.IncreaseCommentCountByOne()
	default:
		return false
	}

	return true
}

func isEventCounterAvailableForTheDate(user *Models.User, eventDate time.Time) (*Models.UserEventCounter, bool) {
	userEventCounters := user.GetUserEventCounters()

	for _, userEventCounter := range userEventCounters {
		if userEventCounter.GetDate().Date() == eventDate.Date() {
			return userEventCounter, true
		}
	}
	return &Models.UserEventCounter{}, false
}

func getUserWithID(userID string) (bool, *Models.User) {

	users := userDataSetInstance.GetUsers()

	for _, eventUser := range users {
		if eventUser.GetID() == userID {
			return true, eventUser
		}
	}

	return false, &Models.User{}
}
