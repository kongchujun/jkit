// Package slice /*
/*
@File: delete
@Time: 2023/9/21-11:11
@Author: Lewis Kong
@Email: kongplusjun@163.com
@Desc: Some description of the file
*/
package slice

import "github.com/kongchujun/jkit/internal/errs"

func Delete[T any](src []T, index int) ([]T, T, error) {
	length := len(src)
	if index < 0 || index >= length {
		var zeroVal T
		return src, zeroVal, errs.NewErrIndexOutOfRange(length, index)
	}
	// delete element and move the rest forward
	res := src[index]
	for i := index; i+1 < length; i++ {
		src[i] = src[i+1]
	}
	src = src[:length-1]
	return src, res, nil
}
