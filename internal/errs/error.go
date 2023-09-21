package errs

import "fmt"

// create an err shows out of list
func NewErrIndexOutOfRange(length int, idx int) error {
	return fmt.Errorf("jkit: out of range, the length is %d, index is %d", length, idx)
}
