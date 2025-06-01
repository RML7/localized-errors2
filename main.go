package main

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"localized-errors2/errs"
)

const (
	ErrTypeUserNotFound errs.ErrType = "UserNotFound"
)

func main() {
	if err := process(); err != nil {
		errorInterceptor(err)
	}
}

func process() error {
	return &errs.LocalizedError{
		Code: ErrTypeUserNotFound,
		Err:  errors.New("sdfdsfsdfsdf"),
		Data: map[string]interface{}{
			"UserID": uuid.New().String(),
		},
	}
}

func errorInterceptor(err error) {
	localizer := errs.NewErrorLocalizer()

	var locErr *errs.LocalizedError
	if errors.As(err, &locErr) {
		fmt.Println(localizer.Localize(locErr, "ru"))
	} else {
		// ...
	}
}
