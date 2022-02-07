package main

import (
	"os"
	"testing"
)

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
		{
			name: "",
			args: args{
				filename: "dir",
				arg:      "hello",
				dest:     "out.txt",
			},
			wantErr: false,
		},
	}
	wg.Add(len(tests) + 1)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := oflag(tt.args.filename, tt.args.arg, tt.args.dest); (err != nil) != tt.wantErr {
				t.Errorf("oflag() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
		os.Remove(tt.args.dest)
	}
}
