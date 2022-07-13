package other

import (
	"reflect"
	"testing"
)

func Test_asteroidCollision(t *testing.T) {
	type args struct {
		asteroids []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case1",
			args: args{
				asteroids: []int{5, 10, -5},
			},
			want: []int{5, 10},
		},
		{
			name: "case2",
			args: args{
				asteroids: []int{8, -8},
			},
			want: []int{},
		},
		{
			name: "case3",
			args: args{
				asteroids: []int{10, 2, -5},
			},
			want: []int{10},
		},
		{
			name: "case4",
			args: args{
				asteroids: []int{1, -1, -2, -2},
			},
			want: []int{-2, -2},
		},
		{
			name: "case5",
			args: args{
				asteroids: []int{-2, -2, 1, -2},
			},
			want: []int{-2, -2, -2},
		},
		{
			name: "case6",
			args: args{
				asteroids: []int{1, -2, -2, -2},
			},
			want: []int{-2, -2, -2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := asteroidCollision(tt.args.asteroids); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("asteroidCollision() = %v, want %v", got, tt.want)
			}
		})
	}
}
