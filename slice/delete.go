// Package slice /*
/*
@File: delete
@Time: 2023/9/24-09:40
@Author: Lewis Kong
@Email: kongplusjun@163.com
@Desc: provide delete method for slice by internal delete method
*/
package slice

import "github.com/kongchujun/jkit/internal/slice"

// Delete delete element from slice by index
func Delete[T any](src []T, index int) ([]T, error) {
	res, _, err := slice.Delete(src, index)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// FilterDelete delete by condition with quick and slow pointer
func FilterDelete[T any](src []T, m func(idx int, src []T) bool) []T {
	emptyPos := 0          // slow pointer
	for idx := range src { // quick pointer
		if m(idx, src) {
			continue
		}
		src[emptyPos] = src[idx]
		emptyPos++
	}
	return src[:emptyPos]
}
