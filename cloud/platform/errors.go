package platform

import "errors"

const (
	ServerErrorMessage  = "Internal server error"
	BadRequestMessage   = "Invalid request"
	UnauthorizedMessage = "Unauthorized"
)

var (
	ErrorInvalidAuthHeader   = errors.New("invalid or missing authorization header")
	ErrorMissingAuthToken    = errors.New("no available tokens in db")
	ErrorMissingContextValue = errors.New("missing context value")
	ErrorInvalidStatus       = errors.New("invalid tip status")
	ErrorInvalidSignature    = errors.New("invalid signature")
	ErrorUnknownMessageType  = errors.New("unknown message type")
	ErrorUnknownEventType    = errors.New("unknown event type")
)
