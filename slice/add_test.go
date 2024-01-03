// Package slice /*
/*
@File: add_test.go
@Time: 2023/9/23-07:39
@Author: Lewis Kong
@Email: kongplusjun@163.com
@Desc: Some description of the file
*/
package slice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdd(t *testing.T) {
	testCase := []struct {
		name    string
		slice   []int
		elem    int
		index   int
		wantRes []int
		wantErr error
	}{
		{
			name:    "test for common",
			slice:   []int{100, 200, 300},
			elem:    1,
			index:   1,
			wantRes: []int{100, 1, 200, 300},
			wantErr: nil,
		},
	}
	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			res, err := Add[int](tc.slice, tc.elem, tc.index)
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantRes, res)
		})
	}
}
