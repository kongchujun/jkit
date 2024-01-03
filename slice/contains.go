// Package slice /*
/*
@File: contains
@Time: 2023/9/25-07:42
@Author: Lewis Kong
@Email: kongplusjun@163.com
@Desc: Some description of the file
*/
package slice

// ContainsFunc can let you build your rule to check the result in slice or not
func ContainsFunc[T any](src []T, equals func(elem T) bool) bool {
	for _, v := range src {
		if equals(v) {
			return true
		}
	}
	return false
}

// Contains common method find element whether in slice
func Contains[T comparable](src []T, dst T) bool {
	return ContainsFunc(src, func(elem T) bool {
		return elem == dst
	})
}

// ContainsAny can check whether slice src contains any elem of dst
func ContainsAny[T comparable](src, dst []T) bool {
	dataMap := toMap[T](src)
	for _, v := range dst {
		if _, exists := dataMap[v]; exists {
			return true
		}
	}
	return false
}

// ContainsAnyFunc check whether slice src contains any elem of dst by func
// ContainsAny should be priority
func ContainsAnyFunc[T any](src, dst []T, equal equalFunc[T]) bool {
	for _, valDst := range dst {
		for _, valSrc := range src {
			if equal(valSrc, valDst) {
				return true
			}
		}
	}
	return false
}

// ContainsAll check whether src contains all elem of dst
func ContainsAll[T comparable](src, dst []T) bool {
	srcMap := toMap[T](src)
	for _, v := range dst {
		if _, exist := srcMap[v]; !exist {
			return false
		}
	}
	return true
}

// ContainsAllFunc check whether src contains all elem of dst by func
// ContainsAll should be priority
func ContainsAllFunc[T any](src, dst []T, equals equalFunc[T]) bool {
	for _, valDst := range dst {
		if !ContainsFunc[T](src, func(elem T) bool {
			return equals(elem, valDst)
		}) {
			return false
		}
	}
	return true
}
