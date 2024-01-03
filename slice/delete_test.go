// Package slice /*
/*
@File: delete_test.go
@Time: 2023/9/24-09:49
@Author: Lewis Kong
@Email: kongplusjun@163.com
@Desc: Some description of the file
*/
package slice

import (
	"github.com/kongchujun/jkit/internal/errs"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDelete(t *testing.T) {
	testCase := []struct {
		name    string
		slice   []int
		index   int
		wantRes []int
		wantErr error
	}{
		{
			name:    "delete common",
			slice:   []int{100, 200, 300, 400},
			index:   2,
			wantRes: []int{100, 200, 400},
			wantErr: nil,
		},
		{
			name:    "delete first",
			slice:   []int{100, 200, 300, 400},
			index:   0,
			wantRes: []int{200, 300, 400},
			wantErr: nil,
		},
		{
			name:    "delete last",
			slice:   []int{100, 200, 300, 400},
			index:   3,
			wantRes: []int{100, 200, 300},
			wantErr: nil,
		},
		{
			name:    "delete error",
			slice:   []int{100, 200, 300, 400},
			index:   -1,
			wantErr: errs.NewErrIndexOutOfRange(4, -1),
		},
		{
			name:    "delete out of list",
			slice:   []int{100, 200, 300, 400},
			index:   6,
			wantErr: errs.NewErrIndexOutOfRange(4, 6),
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			res, err := Delete[int](tc.slice, tc.index)
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, res, tc.wantRes)
		})
	}
}

func TestFilterDelete(t *testing.T) {
	testCase := []struct {
		name          string
		src           []int
		conditionFunc func(idx int, src []int) bool
		wantSrc       []int
	}{
		{
			name: "delete middle one",
			src:  []int{0, 1, 2, 3, 4, 5, 6, 7},
			conditionFunc: func(idx int, src []int) bool {
				return idx == 3
			},
			wantSrc: []int{0, 1, 2, 4, 5, 6, 7},
		},
		{
			name: "delete sequence two ",
			src:  []int{0, 1, 2, 3, 4, 5, 6, 7},
			conditionFunc: func(idx int, src []int) bool {
				return idx == 3 || idx == 4
			},
			wantSrc: []int{0, 1, 2, 5, 6, 7},
		},
		{
			name: "delete different two ",
			src:  []int{0, 1, 2, 3, 4, 5, 6, 7},
			conditionFunc: func(idx int, src []int) bool {
				return idx == 3 || idx == 5
			},
			wantSrc: []int{0, 1, 2, 4, 6, 7},
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			res := FilterDelete[int](tc.src, tc.conditionFunc)
			assert.Equal(t, tc.wantSrc, res)
		})
	}
}
