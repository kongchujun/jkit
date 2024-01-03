// Package slice /*
/*
@File: map
@Time: 2023/9/30-15:54
@Author: Lewis Kong
@Email: kongplusjun@163.com
@Desc: Some description of the file
*/
package slice

// toMap makes slice into map construct
func toMap[T comparable](src []T) map[T]struct{} {
	var dataMap = make(map[T]struct{}, len(src))
	for _, v := range src {
		// use struct{}{} as value in order to save memory
		dataMap[v] = struct{}{}
	}
	return dataMap
}
