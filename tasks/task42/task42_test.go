package task42

import (
	"testing"
)

func Test_trap(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		name     string
		args     args
		wantRain int
	}{
		// TODO: Add test cases.
		{"c1", args{[]int{5, 4, 1, 2}}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRain := trap(tt.args.height); gotRain != tt.wantRain {
				t.Errorf("trap() = %v, want %v", gotRain, tt.wantRain)
			}
		})
	}
}
