package errs

import "fmt"

type ErrType string

type LocalizedError struct {
	Code ErrType
	Err  error
	Data map[string]interface{}
}

func (e *LocalizedError) Error() string {
	if e.Err != nil {
		return e.Err.Error()
	}
	return fmt.Sprintf("LocalizedError code: %s", e.Code)
}
