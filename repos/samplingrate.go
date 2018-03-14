package repos

import (
	"errors"
	"strconv"
)

// RateUnit defines a rate unit
type RateUnit string

// SamplingRate defines a rate to be used for taking regular samples of a repository
type SamplingRate struct {
	Value int
	Unit  RateUnit
}

func (rate *SamplingRate) String() string {
	var unit string
	switch rate.Unit {
	case RateCommit:
		unit = "c"
	case RateDay:
		unit = "d"
	case RateWeek:
		unit = "w"
	case RateMonth:
		unit = "m"
	case RateQuarter:
		unit = "q"
	case RateYear:
		unit = "y"
	}
	return strconv.Itoa(rate.Value) + unit
}

const (
	RateCommit  = RateUnit("commit")
	RateDay     = RateUnit("day")
	RateWeek    = RateUnit("week")
	RateMonth   = RateUnit("month")
	RateQuarter = RateUnit("quarter")
	RateYear    = RateUnit("year")
)

// DecodeSamplingRate decodes a sampling rate from a given text
func DecodeSamplingRate(text string) (SamplingRate, error) {
	if len(text) < 2 {
		return SamplingRate{}, errors.New("invalid sampling rate: expect an integer followed by an unit (example: '1w' for each week)")
	}
	// extract the rate value
	value, err := strconv.Atoi(text[:len(text)-1])
	if err != nil {
		return SamplingRate{}, errors.New("invalid sampling rate: expect an integer followed by an unit (example: '1w' for each week)")
	}

	// extract the rate unit
	lastChar := text[len(text)-1]
	if lastChar < 'a' || lastChar > 'z' {
		return SamplingRate{}, errors.New("invalid sampling rate: unit is missing (example: '1w' for each week)")
	}

	var unit RateUnit
	switch lastChar {
	case 'c':
		unit = RateCommit
	case 'd':
		unit = RateDay
	case 'w':
		unit = RateWeek
	case 'm':
		unit = RateMonth
	case 'q':
		unit = RateQuarter
	case 'y':
		unit = RateYear
	default:
		return SamplingRate{}, errors.New("invalid sampling rate unit: valid units are c, d, w, m, q, y")
	}

	return SamplingRate{value, unit}, nil
}
