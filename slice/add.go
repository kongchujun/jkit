// Package slice /*
/*
@File: add
@Time: 2023/9/21-12:54
@Author: Lewis Kong
@Email: kongplusjun@163.com
@Desc: Some description of the file
*/
package slice

import "github.com/kongchujun/jkit/internal/slice"

func Add[Src any](src []Src, element Src, index int) ([]Src, error) {
	res, err := slice.Add[Src](src, element, index)
	return res, err
}
