package main

import "sugarBox/Handler"

func main() {
	//tr := Payload.EventCreationReq{
	//	UserID:    12,
	//	EventType: "post",
	//	TimeStamp: 1672444800,
	//}
	//
	//Controller.ProcessInput(tr)
	Handler.HandleAndExecuteInput()
}
