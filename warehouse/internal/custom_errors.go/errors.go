package custom_errors

import "errors"

var (
	CategoryNotExist = errors.New("the category does not exist yet")
)
