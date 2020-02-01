package controllers

type errorMessage struct {
	message string
}

func NewError(err error) errorMessage {
	return errorMessage{err.Error()}
}