// Package slice /*
/*
@File: aggregate
@Time: 2023/9/24-23:11
@Author: Lewis Kong
@Email: kongplusjun@163.com
@Desc: Some description of the file
*/
package slice

import "github.com/kongchujun/jkit"

// Max provide the max value of slice
// assume at least one element
func Max[T jkit.RealNumber](src []T) T {
	max := src[0]
	for i := 1; i < len(src); i++ {
		if max < src[i] {
			max = src[i]
		}
	}
	return max
}

// Min provide the min value of slice
// assume at least one element
func Min[T jkit.RealNumber](src []T) T {
	min := src[0]
	for i := 1; i < len(src); i++ {
		if min > src[i] {
			min = src[i]
		}
	}
	return min
}

// Sum is to add all value of slice
func Sum[T jkit.RealNumber](src []T) T {
	var res T
	for _, n := range src {
		res = res + n
	}
	return res
}
