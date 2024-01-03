// Package jkit /*
/*
@File: constrain
@Time: 2023/9/24-23:17
@Author: Lewis Kong
@Email: kongplusjun@163.com
@Desc: Some description of the file
*/
package jkit

type RealNumber interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~float32 | ~float64
}
