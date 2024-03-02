package platform

import "errors"

const (
	ServerErrorMessage  = "Internal server error"
	BadRequestMessage   = "Invalid request"
	UnauthorizedMessage = "Unauthorized"
)

var (
	ErrorMissingAuthToken    = errors.New("no available tokens in db")
	ErrorMissingContextValue = errors.New("missing context value")
	ErrorMissingModuleDep    = errors.New("missing module dependency")
	ErrorInvalidStatus       = errors.New("invalid tip status")
	ErrorInvalidSignature    = errors.New("invalid signature")
	ErrorUnknownEventType    = errors.New("unknown event type")
	ErrorInvalidPayload      = errors.New("invalid payload")
	ErrorObjectNotFound      = errors.New("object not found")
	ErrorBucketNotFound      = errors.New("bucket not found")
	ErrorUnknownEnumValue    = errors.New("unknown enum value")
	ErrorMissingSample       = errors.New("missing sample")
	ErrorResponseNotOk       = errors.New("response not ok")
)
