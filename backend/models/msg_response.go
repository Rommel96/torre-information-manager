package models

type MsgResponse struct {
	Status  StatusResponse `json:"status"`
	Message interface{}    `json:"message"`
}

type StatusResponse string
type MessageData string

const (
	StatusOk           StatusResponse = "OK"
	StatusError        StatusResponse = "ERROR"
	BodyRequest        MessageData    = "Empty or malformed body"
	UserNotFound       MessageData    = "That user was not found. Please try again"
	UserIsTaken        MessageData    = "That username is taken. Try another"
	PasswordIncorrect  MessageData    = "That password was incorrect. Please try again"
	AuthHeaderNotFound MessageData    = "No Authorization header found"
	InvalidToken       MessageData    = "Not Valid Token"
	GroupNotFound      MessageData    = "That Group was not found. Please try again"
	NotPermissions     MessageData    = "You don't have permissions for this process"
	ParamsNotFound     MessageData    = "Params not found"
	UserToAddNotFound  MessageData    = "User to add not found. Please check username"
	ErrorUploadImages  MessageData    = "Error to upload images"
)
