// Package slice /*
/*
@File: find
@Time: 2023/9/24-21:54
@Author: Lewis Kong
@Email: kongplusjun@163.com
@Desc: Some description of the file
*/
package slice

// Find find the element from slice
func Find[T any](src []T, match matchFunc[T]) (T, bool) {
	for _, elem := range src {
		if match(elem) {
			return elem, true
		}
	}
	var t T
	return t, false
}

// FindAll find elements which satisfy match function
// set the new slice for 1/8, so just trigger 3 times as worst situation.
func FindAll[T any](src []T, match matchFunc[T]) []T {
	res := make([]T, 0, len(src)>>3+1)
	for _, val := range src {
		if match(val) {
			res = append(res, val)
		}
	}
	return res
}
