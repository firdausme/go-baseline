package exception

type ValidationError struct {
	Message string
}

func (ValidationError ValidationError) Error() string {
	return ValidationError.Message
}
