package Service

import (
	"errors"
	"fmt"
	"sugarBox/DataHub"
	"sugarBox/Models"
	"sugarBox/Payload"
	"time"
)

var (
	userDataSetInstance  = DataHub.GetUserDataSetInstance()
	eventDataSetInstance = DataHub.GetEventDataSetInstance()
)

func createUser(id int) (*Models.User, error) {
	newUser := Models.User{}

	newUser.SetID(id)
	userDataSetInstance.SetUsers(&newUser)
	return &newUser, nil
}

func addEventToUser(user *Models.User, event *Models.Event) (int, error) {
	var (
		err error
		ok  bool
	)

	/*Todo:: We can make updating event counter and setting event id to user in single
	Todo::  transaction to avoid any inconsistency*/
	if ok = updateEventCounterForUser(user, event); !ok {
		err = errors.New("Error Encountered While Updating The Count ")
		return 0, err
	}

	user.SetEventID(event.GetID())

	return user.GetID(), nil
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
		if fmt.Sprint(userEventCounter.GetDate().Date()) == fmt.Sprint(eventDate.Date()) {
			return userEventCounter, true
		}
	}
	return &Models.UserEventCounter{}, false
}

func getUserWithID(userID int) (bool, *Models.User) {

	users := userDataSetInstance.GetUsers()

	for _, eventUser := range users {
		if eventUser.GetID() == userID {
			return true, eventUser
		}
	}

	return false, &Models.User{}
}

func ListOfUsersWithEventCounters() []Payload.UserDetailsResp {
	var (
		userDetails []Payload.UserDetailsResp
	)

	for _, user := range userDataSetInstance.GetUsers() {
		allCounterForUser := getAllEventCountersForUser(user)
		//fmt.Println("user:", user.GetID(), " ", allCounterForUser)
		userDetails = append(userDetails, allCounterForUser...)
	}
	return userDetails
}

func getAllEventCountersForUser(user *Models.User) []Payload.UserDetailsResp {
	var (
		userDetails []Payload.UserDetailsResp
	)

	for _, eventCounter := range user.GetUserEventCounters() {
		fmt.Println("===================================")

		userDetail := Payload.UserDetailsResp{
			UserID:       user.GetID(),
			Date:         fmt.Sprint(eventCounter.GetDate().Date()),
			LikeReceived: eventCounter.GetLikeCount(),
			Comment:      eventCounter.GetCommentCount(),
			Post:         eventCounter.GetPostCount(),
		}
		fmt.Println("UserID", userDetail.UserID)
		fmt.Println("Date", userDetail.Date)
		fmt.Println("like", userDetail.LikeReceived)
		fmt.Println("comment", userDetail.Comment)
		fmt.Println("post", userDetail.Post)
		userDetails = append(userDetails, userDetail)
	}

	return userDetails
}
