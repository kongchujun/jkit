// Package slice /*
/*
@File: aggregate_test.go
@Time: 2023/9/24-23:19
@Author: Lewis Kong
@Email: kongplusjun@163.com
@Desc: Some description of the file
*/
package slice

import (
	"github.com/kongchujun/jkit"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMax(t *testing.T) {
	testCase := []struct {
		name    string
		slice   []int
		wantMax int
	}{
		{
			name:    "get max value",
			slice:   []int{3, 21, 4, 22, 55, 98},
			wantMax: 98,
		},
	}
	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			res := Max[int](tc.slice)
			assert.Equal(t, res, tc.wantMax)
		})
	}
	testMaxTypes[uint](t)
	testMaxTypes[uint8](t)
	testMaxTypes[uint16](t)
	testMaxTypes[uint32](t)
	testMaxTypes[uint64](t)
	testMaxTypes[int](t)
	testMaxTypes[int8](t)
	testMaxTypes[int16](t)
	testMaxTypes[int32](t)
	testMaxTypes[int64](t)
	testMaxTypes[float32](t)
	testMaxTypes[float64](t)

}

func testMaxTypes[T jkit.RealNumber](t *testing.T) {
	res := Max[T]([]T{1, 2, 3})
	assert.Equal(t, T(3), res)
}

func TestMin(t *testing.T) {
	testCase := []struct {
		name    string
		slice   []int
		wantMin int
	}{
		{
			name:    "get max value",
			slice:   []int{3, 21, 4, 22, 55, 98},
			wantMin: 3,
		},
	}
	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			res := Min[int](tc.slice)
			assert.Equal(t, res, tc.wantMin)
		})
	}
	testMinTypes[uint](t)
	testMinTypes[uint8](t)
	testMinTypes[uint16](t)
	testMinTypes[uint32](t)
	testMinTypes[uint64](t)
	testMinTypes[int](t)
	testMinTypes[int8](t)
	testMinTypes[int16](t)
	testMinTypes[int32](t)
	testMinTypes[int64](t)
	testMinTypes[float32](t)
	testMinTypes[float64](t)
}

func testMinTypes[T jkit.RealNumber](t *testing.T) {
	res := Min[T]([]T{1, 2, 3})
	assert.Equal(t, T(1), res)
}

func TestSum(t *testing.T) {
	testCase := []struct {
		name    string
		slice   []int
		wantRes int
	}{
		{
			name:    "empty",
			slice:   []int{},
			wantRes: 0,
		},
		{
			name:    "common",
			slice:   []int{1, 2, 3},
			wantRes: 6,
		},
	}
	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			res := Sum[int](tc.slice)
			assert.Equal(t, res, tc.wantRes)
		})
	}
}
