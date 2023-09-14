package ad9910

import (
	"math"
)

// FreqOutToFTW returns the frequency-tuning word (FTW) given an output frequency and system clock in Hz.
func FreqOutToFTW(freqOut, sysClk float64) uint32 {
	if freqOut < 0.0 {
		panic("output frequency cannot be less than zero")
	}
	if freqOut > sysClk/2 {
		panic("output frequency cannot be greater than half the system clock")
	}
	return uint32(math.Round(math.Pow(2, 32) * freqOut / sysClk))
}

// LinAmplScaleToASF returns the amplitude scale factor (ASF) given a linear amplitude scale from 0.0 to 1.0.
func LinAmplScaleToASF(amplScale float64) uint16 {
	if amplScale < 0.0 {
		panic("amplitude scale cannot be less than zero")
	}
	if amplScale > 1.0 {
		panic("amplitude scale cannot be greater than one")
	}
	return uint16(math.Round(math.Pow(2, 14) * amplScale))
}

// LogAmplScaleToASF returns the amplitude scale factor (ASF) given an amplitude scale in dB relative to maximum output.
func LogAmplScaleToASF(amplScale float64) uint16 {
	if amplScale > 0.0 {
		panic("amplitude scale cannot be greater than zero")
	}
	if amplScale < -84.2884 {
		panic("amplitude scale cannot be less than -84.2884")
	}
	return uint16(math.Round(math.Pow(2, 14) * math.Pow(10.0, amplScale/20)))
}
