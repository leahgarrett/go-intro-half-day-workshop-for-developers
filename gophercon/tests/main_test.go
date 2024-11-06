package main

import (
	"fmt"
	"testing"
)

func TestNDISAgeCheck(t *testing.T) {

	var agetests = []struct {
		in    int
		label string
		out   bool
	}{
		{in: 64, out: true},
		{in: 65, out: true},
		{in: 6, out: false},
		{in: 7, out: true},
		{in: 24, out: false},
	}

	for _, tt := range agetests {
		label := fmt.Sprintf("%v", tt.in)

		t.Run(label, func(t *testing.T) {
			s := checkNDISAge(tt.in)

			if s != tt.out {
				t.Errorf("%v: got %v, want %v", label, s, tt.out)
			}
		})
	}
}

func TestDrivingAgeCheck(t *testing.T) {

	var agetests = []struct {
		age     int
		country string
		out     bool
	}{
		{age: 16, country: "Australia", out: false},
		{age: 18, country: "Australia", out: true},
		{age: 15, country: "USA", out: false},
		{age: 16, country: "USA", out: true},
		{age: 14, country: "USA", out: false},
	}

	for _, tt := range agetests {
		label := fmt.Sprintf("%v %v", tt.country, tt.age)

		t.Run(label, func(t *testing.T) {
			s := checkDrivingAge(tt.age, tt.country)

			if s != tt.out {
				t.Errorf("%v: got %v, want %v", label, s, tt.out)
			}
		})
	}
}
