package ALGORITHM

import (
	"reflect"
	"testing"
)

func Test_findKClosestElements(t *testing.T) {
	type args struct {
		nums []int
		k    int
		x    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		
		{
			name: "test1",
			args: args{
				nums: []int{1, 2, 3, 4, 5},
				k:    4,
				x:    3,
			},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "test2",
			args: args{
				nums: []int{1, 2, 3, 4, 5},
				k:    4,
				x:    -1,
			},
			want: []int{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findKClosestElements(tt.args.nums, tt.args.k, tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findKClosestElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
