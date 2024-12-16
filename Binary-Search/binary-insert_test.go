package ALGORITHM

import "testing"

func Test_binary_insert(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test1", args: args{nums: []int{1, 2, 4, 7}, target: 9}, want: 4},
		{name: "test2", args: args{nums: []int{1, 3, 6, 9}, target: 3}, want: 1},
		{name: "test3", args: args{nums: []int{1, 4, 8, 9}, target: 5}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binary_insert(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("binary_insert() = %v, want %v", got, tt.want)
			}
		})
	}
}
