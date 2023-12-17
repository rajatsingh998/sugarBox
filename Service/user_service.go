package Service

import (
	"errors"
	"sugarBox/DataHub"
	"sugarBox/Models"
)

var (
	userDataSetInstance  = DataHub.GetUserDataSetInstance()
	eventDataSetInstance = DataHub.GetEventDataSetInstance()
)

func createUser(id string) (Models.User, error) {
	newUser := Models.User{}

	newUser.SetID(id)
	userDataSetInstance.SetUsers(&newUser)
	return newUser, nil
}

func addEventToUser(userId string, eventID string) (string, error) {
	var (
		err  error
		user Models.User
		ok   bool
	)

	if ok, user = getUserWithID(userId); !ok {
		err = errors.New("User Not Found With Given ID ")
		return "", err
	}

	user.SetEventID(eventID)
	return userId, nil
}

func getUserWithID(userID string) (bool, Models.User) {

	users := userDataSetInstance.GetUsers()

	for _, eventUser := range users {
		if eventUser.GetID() == userID {
			return true, *eventUser
		}
	}

	return false, Models.User{}
}
