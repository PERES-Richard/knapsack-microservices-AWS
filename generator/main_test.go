package main

import "testing"

func Test_handleArgs(t *testing.T) {
	tests := []struct {
		name  string
		want  int
		want1 int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := handleArgs()
			if got != tt.want {
				t.Errorf("handleArgs() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("handleArgs() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
