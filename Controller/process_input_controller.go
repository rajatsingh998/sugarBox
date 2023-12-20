package Controller

import (
	"fmt"
	"sugarBox/Payload"
	"sugarBox/Service"
)

func ProcessInput(eventCreationReq Payload.EventCreationReq) error {
	var (
		eventID string
		err     error
	)

	eventID, err = Service.CreateEvent(eventCreationReq.UserID,
		eventCreationReq.EventType, eventCreationReq.TimeStamp)

	if err != nil {
		return err
	}

	fmt.Println("Event ID:", eventID)
	return nil
}
