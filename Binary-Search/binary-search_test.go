package ALGORITHM

import (
	"testing"
)

func TestBinarySearch_1(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				nums:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				target: 5,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinarySearch_1(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("BinarySearch_1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinarySearch_2(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{

		{
			name: "test1",
			args: args{
				nums:   []int{12, 23, 34, 45, 56, 67, 78, 89, 90},
				target: 78,
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinarySearch_2(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("BinarySearch_2() = %v, want %v", got, tt.want)
			}
		})
	}
}
