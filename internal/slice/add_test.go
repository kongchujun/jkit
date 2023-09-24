package slice

import (
	"github.com/kongchujun/jkit/internal/errs"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdd(t *testing.T) {
	testCase := []struct {
		name      string
		slice     []int
		addVal    int
		idx       int
		wantSlice []int
		wantErr   error
	}{
		{
			name:      "index middle",
			slice:     []int{100, 200, 300},
			addVal:    0,
			idx:       1,
			wantSlice: []int{100, 0, 200, 300},
			wantErr:   nil,
		}, {
			name:      "index 0",
			slice:     []int{100, 200, 300},
			addVal:    0,
			idx:       0,
			wantSlice: []int{0, 100, 200, 300},
			wantErr:   nil,
		},
		{
			name:      "index last",
			slice:     []int{100, 200, 300},
			addVal:    0,
			idx:       2,
			wantSlice: []int{100, 200, 0, 300},
			wantErr:   nil,
		},
		{
			name:    "index less than 0",
			slice:   []int{100, 200, 300},
			addVal:  0,
			idx:     -1,
			wantErr: errs.NewErrIndexOutOfRange(3, -1),
		},
		{
			name:    "index more than length",
			slice:   []int{100, 200, 300},
			addVal:  0,
			idx:     4,
			wantErr: errs.NewErrIndexOutOfRange(3, 4),
		},
	}
	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			res, err := Add[int](tc.slice, tc.addVal, tc.idx)
			assert.Equal(t, err, tc.wantErr)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantSlice, res)
		})
	}
}
