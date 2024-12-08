package ALGORITHM

import (
	"reflect"
	"testing"
)

func Test_fourSum_backtrack(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// Add test cases.
		{name: "test1", args: args{nums: []int{1, 0, -1, 0, -2, 2}, target: 0}, want: [][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fourSum_backtrack(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fourSum_backtrack() = %v, want %v", got, tt.want)
			}
		})
	}
}
