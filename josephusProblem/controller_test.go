package main

import "testing"

func TestController(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{14, 2}, 13},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Controller(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("Controller() = %v, want %v", got, tt.want)
			}
		})
	}
}
