package main

import "testing"

func TestFind(t *testing.T) {
	type args struct {
		name  string
		names []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{name: "Finding a String",
		args: args{
			name: "Manu",
			names: ["Manu","Ajit","Pratim","Rohini",]string
		},
		want: "String found",
		//refactoring...
	},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Find(tt.args.name, tt.args.names...); got != tt.want {
				t.Errorf("Find() = %v, want %v", got, tt.want)
			}
		})
	}
}
