package enrollment

import (
	"errors"
)

var ErrUserIDRequired = errors.New("user id is required")
var ErrCourseIDRequired = errors.New("course id is required")
