package slice

import "github.com/kongchujun/jkit/internal/errs"

func Add[T any](src []T, element T, index int) ([]T, error) {
	length := len(src)
	if index < 0 || index >= length {
		return nil, errs.NewErrIndexOutOfRange(length, index)
	}
	var zeroValue T
	src = append(src, zeroValue)
	// move one position back from index
	for i := len(src) - 1; i > index; i-- {
		src[i] = src[i-1]
	}
	src[index] = element
	return src, nil
}
