package main

import "testing"

func Test_core(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "",
		},
	}
	wg.Add(1)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := core(); err != nil {
				t.Errorf("got an error %v", err)
			}
		})
	}
}
