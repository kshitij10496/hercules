package common

import "errors"

var ErrDataEncoding = errors.New("Unable to encode the data for you")

var ErrCourseNotFound = errors.New("The specified course was not found in the database.")
