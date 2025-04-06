package ALGORITHM

import (
	"reflect"
	"testing"
)

func TestReverseList1(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseList1(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseList1() = %v, want %v", got, tt.want)
			}
		})
	}
}
