package models

type ErrNoRecord struct{}

type ClientError struct{
	Message string
}

func NewClientError(message string) *ClientError {
	return &ClientError{Message: message}
}

func (err *ErrNoRecord) Error() string {
	return "No matching record found"
}

func (err *ClientError) Error() string {
	return err.Message
}
