package main

import (
	"testing"
)

func TestAD9910_LogAmplScaleToASF(t *testing.T) {
	ad9910 := &AD9910{}

	tests := []struct {
		in  float64
		out uint16
	}{
		{0.0, uint16(0x4000)},
		{-84.2884, uint16(0x1)},
	}

	for _, test := range tests {
		got := ad9910.LogAmplScaleToASF(test.in)
		want := test.out

		if got != want {
			t.Errorf("got %x; want %x", got, want)
		}
	}
}

func TestAD9910_FreqOutToFTW(t *testing.T) {
	ad9910 := &AD9910{
		SysClock: 122.88e6,
	}

	tests := []struct {
		in  float64
		out uint32
	}{
		{41e6, uint32(0x556aaaab)},
	}

	for _, test := range tests {
		got := ad9910.FreqOutputToFTW(test.in)
		want := test.out

		if got != want {
			t.Errorf("got %x; want %x", got, want)
		}
	}
}
