package main

import (
	"testing"
)

func Test_search(t *testing.T) {
	type args struct {
		filename string
		args     string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "1",
			args:    args{"dir", "hello"},
			wantErr: false,
		}, {
			name:    "2",
			args:    args{"test.txt", "hello"},
			wantErr: true,
		}, {
			name:    "3",
			args:    args{"random", "hello"},
			wantErr: true,
		},
	}
	wg.Add(len(tests))
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := search(tt.args.filename, tt.args.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("search() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
