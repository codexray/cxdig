package repos

import (
	"errors"
	"strconv"
)

// FreqUnit defines a frequency unit
type FreqUnit string

// SamplingFreq defines a frequency to be used for taking regular samples of a repository
type SamplingFreq struct {
	Value int
	Unit  FreqUnit
}

const (
	FreqCommit  = FreqUnit("commit")
	FreqDay     = FreqUnit("day")
	FreqWeek    = FreqUnit("week")
	FreqMonth   = FreqUnit("month")
	FreqQuarter = FreqUnit("quarter")
	FreqYear    = FreqUnit("year")
)

// DecodeSamplingFreq decodes a sampling frequency from a given text
func DecodeSamplingFreq(text string) (SamplingFreq, error) {
	if len(text) < 2 {
		return SamplingFreq{}, errors.New("invalid sampling frequency: expect an integer followed by an unit (example: '1w' for each week)")
	}
	// extract the frequency value
	value, err := strconv.Atoi(text[:len(text)-1])
	if err != nil {
		return SamplingFreq{}, errors.New("invalid sampling frequency: expect an integer followed by an unit (example: '1w' for each week)")
	}

	// extract the frequency unit
	lastChar := text[len(text)-1]
	if lastChar < 'a' || lastChar > 'z' {
		return SamplingFreq{}, errors.New("invalid sampling frequency: unit is missing (example: '1w' for each week)")
	}

	var unit FreqUnit
	switch lastChar {
	case 'c':
		unit = FreqCommit
	case 'd':
		unit = FreqDay
	case 'w':
		unit = FreqWeek
	case 'm':
		unit = FreqMonth
	case 'q':
		unit = FreqQuarter
	case 'y':
		unit = FreqYear
	default:
		return SamplingFreq{}, errors.New("invalid sampling frequency unit: valid units are c, d, w, m, q, y")
	}

	return SamplingFreq{value, unit}, nil
}
