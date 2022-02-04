package main

import "testing"

func Test_oflag(t *testing.T) {
	type args struct {
		filename string
		arg      string
		dest     string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := oflag(tt.args.filename, tt.args.arg, tt.args.dest); (err != nil) != tt.wantErr {
				t.Errorf("oflag() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
