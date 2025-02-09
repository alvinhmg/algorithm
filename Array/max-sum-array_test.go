package ALGORITHM

import (
	"testing"
)

func Test_maxSumArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			},
			want: 6,
		},
		{
			name: "test2",
			args: args{
				nums: []int{1},
			},
			want: 1,
		},
		{
			name: "test3",
			args: args{
				nums: []int{5, 4, -1, 7, 8},
			},
			want: 23,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSumArray(tt.args.nums); got != tt.want {
				t.Errorf("maxSumArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxSumArray_Dynamic(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			},
			want: 6,
		},
		{
			name: "test2",
			args: args{
				nums: []int{1},
			},
			want: 1,
		},
		{
			name: "test3",
			args: args{
				nums: []int{5, 4, -1, 7, 8},
			},
			want: 23,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSumArray_Dynamic(tt.args.nums); got != tt.want {
				t.Errorf("maxSumArray_Dynamic() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxSumArray_PrefixSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			},
			want: 6,
		},
		{
			name: "test2",
			args: args{
				nums: []int{1},
			},
			want: 1,
		},
		{
			name: "test3",
			args: args{
				nums: []int{5, 4, -1, 7, 8},
			},
			want: 23,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSumArray_PrefixSum(tt.args.nums); got != tt.want {
				t.Errorf("maxSumArray_PrefixSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
