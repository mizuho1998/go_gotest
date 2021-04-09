package tddbc

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestString(t *testing.T) {
	tests := map[string]struct {
		ci       ClosedInterval
		excepted string
		msg      string
	}{
		"[3,8]":  {ClosedInterval{3, 8}, "[3, 8]", "閉区間[3,8]の文字列表現は[3, 8]"},
		"[1,10]": {ClosedInterval{1, 10}, "[1, 10]", "閉区間[1,10]の文字列表現は[1, 10]"},
	}

	for _, test := range tests {
		t.Run(test.msg, func(t *testing.T) {
			assert.Equal(t, test.excepted, test.ci.String(), test.msg)
		})
	}
}

func TestIsIncludeDot(t *testing.T) {
	tests := map[string]struct {
		ci       ClosedInterval
		val      int
		excepted bool
		msg      string
	}{
		"[3,8], 2": {ClosedInterval{3, 8}, 2, false, "2は閉区間[3,8]に含まれない"},
		"[3,8], 3": {ClosedInterval{3, 8}, 3, true, "3は閉区間[3,8]に含まれる"},
		"[3,8], 8": {ClosedInterval{3, 8}, 8, true, "8は閉区間[3,8]に含まれる"},
		"[3,8], 9": {ClosedInterval{3, 8}, 9, false, "9は閉区間[3,8]に含まれない"},
	}

	for _, test := range tests {
		t.Run(test.msg, func(t *testing.T) {
			assert.Equal(t, test.excepted, test.ci.IsIncludeDot(test.val), test.msg)
		})
	}
}

func TestIsSame(t *testing.T) {
	tests := map[string]struct {
		ci1      ClosedInterval
		ci2      ClosedInterval
		excepted bool
		msg      string
	}{
		"[3,8], [3,8]": {ClosedInterval{3, 8}, ClosedInterval{3, 8}, true, "閉区間[3,8]は閉区間[3,8]と等価"},
		"[3,8], [1,5]": {ClosedInterval{3, 8}, ClosedInterval{1, 5}, false, "閉区間[3,8]は閉区間[1,5]と等価でない"},
	}

	for _, test := range tests {
		t.Run(test.msg, func(t *testing.T) {
			assert.Equal(t, test.excepted, test.ci1.IsSame(test.ci2), test.msg)
		})
	}
}

func TestInclude(t *testing.T) {
	tests := map[string]struct {
		ci1      ClosedInterval
		ci2      ClosedInterval
		excepted bool
		msg      string
	}{
		"[3,8], [1,9]": {ClosedInterval{3, 8}, ClosedInterval{1, 9}, true, "閉区間[3,8]は閉区間[1,9]に完全に含まれる"},
		"[3,8], [1,6]": {ClosedInterval{3, 8}, ClosedInterval{1, 6}, false, "閉区間[3,8]は閉区間[1,6]に完全に含まれない"},
		"[3,8], [5,9]": {ClosedInterval{3, 8}, ClosedInterval{5, 9}, false, "閉区間[3,8]は閉区間[5,9]に完全に含まれない"},
		"[3,8], [4,6]": {ClosedInterval{3, 8}, ClosedInterval{4, 6}, false, "閉区間[3,8]は閉区間[4,6]に完全に含まれない"},
	}

	for _, test := range tests {
		t.Run(test.msg, func(t *testing.T) {
			assert.Equal(t, test.excepted, test.ci1.IsInclude(test.ci2), test.msg)
		})
	}
}

func TestInitOk(t *testing.T) {
	tests := map[string]struct {
		lower int
		upper int
		msg   string
	}{
		"[3,8]": {3, 8, "上端点より下端点が小さい閉区間を作ろうとした場合はok"},
	}

	for _, test := range tests {
		t.Run(test.msg, func(t *testing.T) {
			_, err := NewClosedInterval(test.lower, test.upper)
			assert.Nil(t, err, test.msg)
		})
	}
}
func TestInitError(t *testing.T) {
	tests := map[string]struct {
		lower int
		upper int
		msg   string
	}{
		"[9,8]": {9, 8, "上端点より下端点が大きい閉区間を作ろうとした場合はエラー"},
	}

	for _, test := range tests {
		t.Run(test.msg, func(t *testing.T) {
			_, err := NewClosedInterval(test.lower, test.upper)
			assert.Error(t, err, test.msg)
		})
	}
}
