package enrollment

import (
	"errors"
	"fmt"
)

var ErrUserIDRequired = errors.New("user id is required")
var ErrCourseIDRequired = errors.New("course id is required")
var ErrStatusRequired = errors.New("status is required")

type ErrNotFound struct {
	EnrollmentID string
}

func (e ErrNotFound) Error() string {
	return fmt.Sprintf("enrollment with id '%s' not found or doesn't exist", e.EnrollmentID)
}

type ErrInvalidStatus struct {
	Status string
}

func (e ErrInvalidStatus) Error() string {
	return fmt.Sprintf("invalid '%s' status", e.Status)
}
