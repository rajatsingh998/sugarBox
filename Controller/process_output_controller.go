package Controller

import (
	"sugarBox/Payload"
	"sugarBox/Service"
)

func ProcessOutput() []Payload.UserDetailsResp {
	return Service.ListOfUsersWithEventCounters()
}
