package main

import (
	"math"
)

// AD9910 DDS chip.
type AD9910 struct {
	// SysClock denotes the frequency in Hz of the system clock.
	SysClock float64
}

// LogAmplScaleToASF returns the amplitude scale factor (ASF) given an amplitude scale in dB.
func (a AD9910) LogAmplScaleToASF(ampl float64) uint16 {
	return uint16(math.Round(math.Pow(2, 14) * math.Pow(10.0, ampl/20)))
}

// FreqOutputToFTW returns the frequency tuning word (FTW) given a frequency in Hz.
func (a AD9910) FreqOutputToFTW(freq float64) uint32 {
	return uint32(math.Round(math.Pow(2, 32) * freq / a.SysClock))
}
