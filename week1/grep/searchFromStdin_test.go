package main

import (
	"testing"
)

func TestSearchstdin(t *testing.T) {
	type args struct {
		arg string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "foomain",
			args: args{
				arg: "foo",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Searchstdin(tt.args.arg); (err != nil) != tt.wantErr {
				t.Errorf("Searchstdin() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestChecking(t *testing.T) {
	type args struct {
		input []string
		arg   string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "foo",
			args: args{
				input: []string{"foobar", "barbazfoo", "food"},
				arg:   "foo",
			},
			wantErr: false,
		}, {
			name: "hello",
			args: args{
				input: []string{"foobar", "barbazfoo", "food"},
				arg:   "hello",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Checking(tt.args.input, tt.args.arg); (err != nil) != tt.wantErr {
				t.Errorf("Checking() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
