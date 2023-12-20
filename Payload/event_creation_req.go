package Payload

type EventCreationReq struct {
	UserID    int    `json:"userId"`
	EventType string `json:"eventType"`
	TimeStamp int64  `json:"timeStamp"`
}
