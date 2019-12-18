package helper

import "errors"

// errors : custom types
var (
	ErrBadRouting          = errors.New("Inconsistent mapping between route and handler")
	ErrInternalServerError = errors.New("Internal Server Error")
	ErrNotFoundError       = errors.New("Your requested item is not found")
)

var (
	InternalServerError         = "Internal Server Error"
	StatusUnKnown               = "UNKNOWN"
	StatusOk                    = "OK"
	SuccessCodeType             = "Success"
	MessageRegisterSuccess      = "Register Successfully"
	MessageLoginSuccess         = "Login Successfully"
	MessageCreatedSuccess       = "Created Successfully"
	MessageRetrievedSuccess     = "Retrieved Successfully"
	MessageRetrievedListSuccess = "Retrieved List Successfully"
	MessageUpdatedSuccess       = "Updated Successfully"
	MessageDeletedSuccess       = "Deleted Successfully"
	MessageValidatedSuccess     = "Validated Successfully"
	MessageSentSuccess          = "Sent Successfully"
)
