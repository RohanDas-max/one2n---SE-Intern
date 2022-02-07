package main

import "testing"

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
			name: "foo",
			args: args{
				arg: "foo",
			},
			wantErr: false,
		},
	}
	wg.Add(1)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Searchstdin(tt.args.arg); (err != nil) != tt.wantErr {
				t.Errorf("Searchstdin() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
