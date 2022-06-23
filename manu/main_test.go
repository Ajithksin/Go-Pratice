package main

import (
	"reflect"
	"testing"
)

func Test_spec_tata(t *testing.T) {
	tests := []struct {
		name string
		s    *spec
		want spec
	}{
		// TODO: Add test cases.

		{name: "TATA Specifications",
			s: &spec{
				company: "TATA",
				model:   "NEXA",
				drive:   "Four-wheel Drive",
				fuel:    "EV",
				price:   750000,
			},
			want: spec{
				company: "TATA",
				model:   "NEXA",
				drive:   "Four-wheel Drive",
				fuel:    "EV",
				price:   750000,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.tata(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("spec.tata() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_spec_mahindra(t *testing.T) {
	tests := []struct {
		name string
		s    *spec
		want spec
	}{
		// TODO: Add test cases.
		{name: "Mahindra Specifications",
			s: &spec{
				company: "MAHINDRA",
				model:   "XUV",
				drive:   "Four-wheel Drive",
				fuel:    "EV",
				price:   1500000,
			},
			want: spec{
				company: "MAHINDRA",
				model:   "XUV",
				drive:   "Four-wheel Drive",
				fuel:    "EV",
				price:   1500000,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.mahindra(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("spec.mahindra() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCarType(t *testing.T) {
	tests := []struct {
		name string
		want spec
	}{
		// TODO: Add test cases.
		{name: "CAR TYPE",

			
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CarType(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CarType() = %v, want %v", got, tt.want)
			}
		})
	}
}
