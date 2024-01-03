// Package slice /*
/*
@File: contains_test.go
@Time: 2023/9/25-07:49
@Author: Lewis Kong
@Email: kongplusjun@163.com
@Desc: Some description of the file
*/
package slice

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContainsFunc(t *testing.T) {
	testCase := []struct {
		name  string
		slice []int

		element int
		wantRes bool
	}{
		{
			name:    "common test",
			slice:   []int{1, 2, 3, 4, 5},
			element: 5,
			wantRes: true,
		},
		{
			name:    "emtpy",
			slice:   []int{},
			element: 4,
			wantRes: false,
		},
		{
			name:    "src nil",
			element: 4,
			wantRes: false,
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			res := ContainsFunc[int](tc.slice, func(elem int) bool {
				return elem == tc.element
			})
			assert.Equal(t, res, tc.wantRes)
		})
	}
}

func TestContains(t *testing.T) {
	testCase := []struct {
		name    string
		slice   []int
		element int
		wantRes bool
	}{
		{
			name:    "common test",
			slice:   []int{1, 2, 3, 4, 5, 6},
			element: 5,
			wantRes: true,
		},
		{
			name:    "not in slice",
			slice:   []int{1, 2, 3, 4, 5, 6},
			element: 0,
			wantRes: false,
		},
		{
			name:    "slice emtpy",
			slice:   []int{},
			element: 0,
			wantRes: false,
		},
		{
			name:    "slice nil",
			element: 0,
			wantRes: false,
		},
		{
			name:    "element nil",
			slice:   []int{},
			wantRes: false,
		},
	}
	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			res := Contains[int](tc.slice, tc.element)
			assert.Equal(t, res, tc.wantRes)
		})
	}
}

func TestContainsAny(t *testing.T) {
	testCase := []struct {
		name    string
		slice   []int
		dist    []int
		wantRes bool
	}{
		{
			name:    "common",
			slice:   []int{1, 2, 3, 4},
			dist:    []int{4, 5, 6},
			wantRes: true,
		},
		{
			name:    "not in it",
			slice:   []int{1, 2, 3, 4},
			dist:    []int{5, 6},
			wantRes: false,
		},
		{
			name:    "many in it",
			slice:   []int{1, 2, 3, 4},
			dist:    []int{1, 2, 3},
			wantRes: true,
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			res := ContainsAny[int](tc.slice, tc.dist)
			assert.Equal(t, tc.wantRes, res)
		})
	}
}

func TestContainsAnyFunc(t *testing.T) {
	testCase := []struct {
		name    string
		slice   []int
		dst     []int
		wantRes bool
	}{
		{
			name:    "common",
			slice:   []int{1, 2, 3, 4},
			dst:     []int{4, 5, 6},
			wantRes: true,
		},
		{
			name:    "not in it",
			slice:   []int{1, 2, 3, 4},
			dst:     []int{5, 6},
			wantRes: false,
		},
		{
			name:    "many in it",
			slice:   []int{1, 2, 3, 4},
			dst:     []int{1, 2, 3},
			wantRes: true,
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			res := ContainsAnyFunc[int](tc.slice, tc.dst, func(src, dst int) bool {
				return src == dst
			})
			assert.Equal(t, tc.wantRes, res)
		})
	}
}

func TestContainsAll(t *testing.T) {
	testCase := []struct {
		name    string
		slice   []int
		dst     []int
		wantRes bool
	}{
		{
			name:    "contains all",
			slice:   []int{1, 2, 3, 4, 5, 6, 7},
			dst:     []int{1, 2, 3, 4},
			wantRes: true,
		},
		{
			name:    "contains part",
			slice:   []int{1, 2, 3, 4, 5, 6, 7},
			dst:     []int{1, 2, 3, 4, 8},
			wantRes: false,
		},
		{
			name:    "contains none of dst",
			slice:   []int{1, 2, 3, 4, 5, 6, 7},
			dst:     []int{8, 9, 10},
			wantRes: false,
		},
	}
	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			res := ContainsAll[int](tc.slice, tc.dst)
			assert.Equal(t, res, tc.wantRes)
		})
	}
}

func TestContainsAllFunc(t *testing.T) {
	testCase := []struct {
		name    string
		slice   []int
		dst     []int
		wantRes bool
	}{
		{
			name:    "contains all",
			slice:   []int{1, 2, 3, 4, 5, 6, 7},
			dst:     []int{1, 2, 3, 4},
			wantRes: true,
		},
		{
			name:    "contains part",
			slice:   []int{1, 2, 3, 4, 5, 6, 7},
			dst:     []int{1, 2, 3, 4, 8},
			wantRes: false,
		},
		{
			name:    "contains none of dst",
			slice:   []int{1, 2, 3, 4, 5, 6, 7},
			dst:     []int{8, 9, 10},
			wantRes: false,
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			res := ContainsAllFunc[int](tc.slice, tc.dst, func(src, dst int) bool {
				return src == dst
			})
			assert.Equal(t, res, tc.wantRes)
		})
	}
}

func ExampleContainsAny() {
	res := ContainsAny[int]([]int{1, 2, 3, 4, 5}, []int{5, 6, 7})
	fmt.Println(res)
	res = ContainsAny[int]([]int{1, 2, 3, 4, 5}, []int{6, 7})
	fmt.Println(res)
	//Output:
	//true
	//false
}

func ExampleContainsAnyFunc() {
	res := ContainsAnyFunc[int]([]int{1, 2, 3, 4, 5}, []int{5, 6, 7}, func(src, dst int) bool {
		return src == dst
	})
	fmt.Println(res)
	res = ContainsAnyFunc[int]([]int{1, 2, 3, 4, 5}, []int{6, 7}, func(src, dst int) bool {
		return src == dst
	})
	fmt.Println(res)
	//Output:
	//true
	//false
}
