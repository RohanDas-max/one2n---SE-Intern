package main

import "testing"

func Test_flags(t *testing.T) {
	tests := []struct {
		name  string
		want  int
		want1 int
	}{
		{
			name:  "1",
			want:  14,
			want1: 2,
		},
		// {
		// 	name:  "1",
		// 	want:  14,
		// 	want1: 2,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := flags()
			if got == tt.want && got1 == tt.want1 {
				t.Errorf("flags() got = %v, %v, want %v,%v", got, got1, tt.want, tt.want1)
			}
		})
	}
}
