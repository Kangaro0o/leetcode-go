package slidingwindow

import "testing"

func TestLengthOfLongestSubstringTwoDistinct(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				s: "ecebab",
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LengthOfLongestSubstringTwoDistinct(tt.args.s); got != tt.want {
				t.Errorf("LengthOfLongestSubstringTwoDistinct() = %v, want %v", got, tt.want)
			}
		})
	}
}
