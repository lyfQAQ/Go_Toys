package error

import (
	"errors"
	"testing"
)

var MyError = errors.New("n must in [2, 100)")

func GetFibinacci(n int) ([]int, error) {
	if n < 2 || n > 100 {
		return nil, MyError
	}
	ret := []int{1, 1}

	for i := 2; i < n; i++ {
		ret = append(ret, ret[i-1]+ret[i-2])
	}
	return ret, nil
}

func TestError(t *testing.T) {
	t.Log(GetFibinacci(-1))
}
