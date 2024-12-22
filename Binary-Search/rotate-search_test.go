package ALGORITHM

import "testing"

func Test_rotateSearch(t *testing.T) {
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
				nums:   []int{4, 5, 6, 7, 0, 1, 2},
				target: 0,
			},	
			want: 4,
		},
		{
			name: "test2",
			args: args{
				nums:   []int{4, 5, 6, 7, 0, 1, 2},
				target: 3,
			},	
			want: -1,
		},
		{
			name: "test3",
			args: args{
				nums:   []int{1},
				target: 0,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotateSearch(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("rotateSearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
