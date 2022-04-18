package first_last_position

import (
	"reflect"
	"testing"
)

func Test_searchRange(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example1",
			args: args{nums: []int{5, 7, 7, 8, 8, 10}, target: 8},
			want: []int{3, 4},
		},
		{
			name: "Example2",
			args: args{nums: []int{5, 7, 7, 8, 8, 10}, target: 6},
			want: []int{-1, -1},
		},
		{
			name: "Example3",
			args: args{nums: []int{}, target: 0},
			want: []int{-1, -1},
		},
		{
			name: "Example4",
			args: args{nums: []int{1, 2, 3, 3, 3, 3, 4, 5, 9}, target: 3},
			want: []int{2, 5},
		},
		{
			name: "Example5",
			args: args{nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 3, 4, 4, 4, 4, 5, 5, 6, 6, 6, 8, 10, 10}, target: 4},
			want: []int{10, 13},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchRange(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("searchRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
