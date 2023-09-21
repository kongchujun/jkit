// Package slice /*
/*
@File: delete_test.go
@Time: 2023/9/21-12:40
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
		name      string
		slice     []int
		index     int
		wantSlice []int
		wantElem  int
		wantErr   error
	}{
		{
			name:      "common test",
			slice:     []int{1, 2, 3, 4, 5},
			index:     3,
			wantSlice: []int{1, 2, 3, 5},
			wantElem:  4,
			wantErr:   nil,
		},
		{
			name:      "delete first",
			slice:     []int{1, 2, 3, 4, 5},
			index:     0,
			wantSlice: []int{2, 3, 4, 5},
			wantElem:  1,
			wantErr:   nil,
		},
		{
			name:      "delete last",
			slice:     []int{1, 2, 3, 4, 5},
			index:     4,
			wantSlice: []int{1, 2, 3, 4},
			wantElem:  5,
			wantErr:   nil,
		},
		{
			name:  "delete error less than 0",
			slice: []int{1, 2, 3, 4, 5},
			index: -1,

			wantErr: errs.NewErrIndexOutOfRange(5, -1),
		},
		{
			name:  "delete error out of range",
			slice: []int{1, 2, 3, 4, 5},
			index: 5,

			wantErr: errs.NewErrIndexOutOfRange(5, 5),
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			resSlice, elem, err := Delete(tc.slice, tc.index)
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantSlice, resSlice)
			assert.Equal(t, tc.wantElem, elem)
		})
	}
}
