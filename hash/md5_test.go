package hash

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringMd5(t *testing.T) {
	testCase := []struct {
		name     string
		inputStr string
		wantStr  string
	}{
		{name: "common test",
			inputStr: "kong",
			wantStr:  "153ddfb15ae1e37b7cf004b201c3e3fd",
		},
	}
	for _, tc := range testCase {
		res := StringMd5(tc.inputStr)
		assert.Equal(t, tc.wantStr, res)
	}
}
