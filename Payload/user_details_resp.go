package Payload

type UserDetailsResp struct {
	UserID       int         `json:"userId,omitempty"`
	Date         interface{} `json:"date,omitempty"`
	Post         int64       `json:"post,omitempty"`
	Comment      int64       `json:"comment,omitempty"`
	LikeReceived int64       `json:"likeReceived,omitempty"`
}
