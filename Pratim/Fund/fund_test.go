package funding

import (
	"reflect"
	"testing"
)

func TestNewFund(t *testing.T) {
	type args struct {
		initialBalance int
	}
	tests := []struct {
		name string
		args args
		want *Fund
	}{
		{
			name: "funding",
			args: args{
				initialBalance: 30000,
			},
			//want : *Fund
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFund(tt.args.initialBalance); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFund() = %v, want %v", got, tt.want)
			}
		})
	}
}
