// Package slice /*
/*
@File: find_test.go
@Time: 2023/9/24-22:00
@Author: Lewis Kong
@Email: kongplusjun@163.com
@Desc: Some description of the file
*/
package slice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFind(t *testing.T) {
	testCase := []struct {
		name          string
		slice         []int
		conditionFunc matchFunc[int]
		wantRes       int
		wantFlag      bool
	}{
		{
			name:  "common logic",
			slice: []int{1, 2, 3, 4, 5, 6},
			conditionFunc: func(src int) bool {
				if src == 5 {
					return true
				}
				return false
			},
			wantRes:  5,
			wantFlag: true,
		},

		{
			name:  "no find",
			slice: []int{1, 2, 3, 4, 5, 6},
			conditionFunc: func(src int) bool {
				return src == 20
			},
			wantRes:  0,
			wantFlag: false,
		},
	}
	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			res, flag := Find(tc.slice, tc.conditionFunc)
			assert.Equal(t, tc.wantFlag, flag)
			if !flag {
				return
			}
			assert.Equal(t, tc.wantRes, res)
		})
	}
}

func TestFindAll(t *testing.T) {
	testCase := []struct {
		name           string
		slice          []int
		conditionMatch matchFunc[int]
		wantRes        []int
	}{
		{
			name:  "findAll_empty",
			slice: []int{2, 4},
			conditionMatch: func(src int) bool {
				return src%2 == 1
			},
			wantRes: []int{},
		},
		{
			name:  "findAll part of them",
			slice: []int{1, 2, 3, 4, 5, 6, 7},
			conditionMatch: func(src int) bool {
				return src%2 == 1
			},
			wantRes: []int{1, 3, 5, 7},
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			res := FindAll[int](tc.slice, tc.conditionMatch)
			assert.Equal(t, res, tc.wantRes)
		})
	}
}
