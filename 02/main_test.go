package main

import (
	"reflect"
	"testing"
)

func TestUnpacker(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"Empty string", args{""}, "", true},
		{"Single character", args{"abcd"}, "abcd", false},
		{"Simple repetition", args{"a6b2c"}, "aaaaaabbc", false},
		{"Invalid input: digit first", args{"5ra"}, "", true},
		{"Complex repetition0", args{"a4b2c3d"}, "aaaabbcccd", false},
		{"Complex repetition1", args{"a4bc2d5e"}, "aaaabccddddde", false},
		{"Complex repetition2", args{"f3cn20"}, "fffcnnnnnnnnnnnnnnnnnnnn", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Unpacker(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Unpacker() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Unpacker() = %v, want %v", got, tt.want)
			}
		})
	}
}
