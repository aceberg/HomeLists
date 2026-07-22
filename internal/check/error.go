package check

import (
	"fmt"
	"log"
)

// IfError prints error, if it is not nil and returns true, otherwise returns false
func IfError(err error) bool {
	if err == nil {
		return false
	}

	log.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("ERROR: %s", err))
	return true
}
