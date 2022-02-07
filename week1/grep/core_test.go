package main

import "testing"

func Test_core(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := core(); (err != nil) != tt.wantErr {
				t.Errorf("core() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
