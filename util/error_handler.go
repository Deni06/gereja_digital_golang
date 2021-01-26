package util

type UnhandledError struct {
	ErrorMessage string
}

func (ce UnhandledError) Error()string{
	return ce.ErrorMessage
}
