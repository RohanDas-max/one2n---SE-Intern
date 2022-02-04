package main

import "testing"

func Test_flags(t *testing.T) {
	tests := []struct {
		name  string
		want  int
		want1 int
	}{
		{
			name:  "",
			want:  0,
			want1: 0,
		},
		// {
		// 	name:  "",
		// 	want:  14,
		// 	want1: 3,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := flags()
			if got != tt.want {
				t.Errorf("flags() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("flags() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
