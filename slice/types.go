// Package slice /*
/*
@File: types
@Time: 2023/9/24-21:53
@Author: Lewis Kong
@Email: kongplusjun@163.com
@Desc: Some description of the file
*/
package slice

type matchFunc[T any] func(src T) bool
type equalFunc[T any] func(src, dst T) bool
