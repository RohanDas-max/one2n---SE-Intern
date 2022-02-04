package main

import "testing"

func TestSearchstdin(t *testing.T) {
	type args struct {
		input []string
		arg   string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"random", args{
			input: []string{},
			arg:   "",
		}, false},
	}
	wg.Add(len(tests))
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Searchstdin(tt.args.input, tt.args.arg); (err != nil) != tt.wantErr {
				t.Errorf("Searchstdin() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
