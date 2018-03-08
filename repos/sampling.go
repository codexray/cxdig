package repos

import (
	"codexray/cxdig/types"
	"time"
)

func FilterCommitsByStep(commits []types.CommitInfo, freq SamplingFreq, limit int) []types.SampleInfo {
	if limit == 0 {
		limit = len(commits)
	}
	rtn := []types.SampleInfo{}
	sampleNumber := 1
	if freq.Unit == FreqCommit {
		for i := 0; i < limit*freq.Value && i < len(commits); i += freq.Value {
			rtn = append(rtn, types.SampleInfo{
				Number:   types.SampleID(sampleNumber),
				DateTime: commits[i].DateTime,
				CommitID: commits[i].CommitID,
			})
			sampleNumber++
		}
	} else {
		step := bringToLastMoment(commits[0].DateTime, freq.Unit)
		j := 0
		for i := 0; j <= limit && i < len(commits); i++ {
			if commits[i].DateTime.Before(step) {
				var t time.Time
				firstAdded := 0
				for t = step; commits[i].DateTime.Before(t); t = getNextStep(t, freq) {
					temp := types.SampleInfo{
						Number:   types.SampleID(sampleNumber),
						DateTime: t,
						CommitID: commits[i].CommitID,
					}
					if firstAdded == 0 {
						firstAdded = sampleNumber
					} else {
						temp.AliasOf = types.SampleID(firstAdded)
					}
					rtn = append(rtn, temp)
					sampleNumber++
				}
				j++
				step = t
			}
		}
	}
	return rtn
}

func getNextStep(step time.Time, freq SamplingFreq) time.Time {
	switch freq.Unit {
	case FreqDay:
		step = step.AddDate(0, 0, -freq.Value)
	case FreqWeek:
		step = step.AddDate(0, 0, -freq.Value*8)
	case FreqMonth:
		for i := 0; i < freq.Value; i++ {
			temp := step.Month()
			step = step.AddDate(0, -1, 0)
			for step.Month() == temp {
				step = step.AddDate(0, 0, -1)
			}
		}
	case FreqQuarter:
		for i := 0; i < freq.Value*3; i++ {
			temp := step.Month()
			step = step.AddDate(0, -1, 0)
			for step.Month() == temp {
				step = step.AddDate(0, 0, -1)
			}
		}
	case FreqYear:
		step = step.AddDate(-freq.Value, 0, 0)
	}
	return bringToLastMoment(step, freq.Unit)
}

func bringToLastMoment(t time.Time, freqUnit FreqUnit) time.Time {
	t = t.Add((time.Hour * time.Duration(23-t.Hour())) + (time.Minute * time.Duration(59-t.Minute())) + (time.Second * time.Duration(59-t.Second())))

	if freqUnit == FreqWeek {
		i := 0
		for t.AddDate(0, 0, i).Weekday() == time.Weekday(7) {
			i++
		}
		t = t.AddDate(0, 0, i)
	}

	if freqUnit == FreqYear {
		i := 0
		for t.AddDate(0, i, 0).Month() != time.Month(12) {
			i++
		}
		t = t.AddDate(0, i, 0)
	}

	if freqUnit == FreqQuarter {
		i := 0
		for t.AddDate(0, i, 0).Month()%3 != 0 {
			i++
		}
		t = t.AddDate(0, i, 0)
	}

	if freqUnit == FreqMonth || freqUnit == FreqQuarter || freqUnit == FreqYear {
		i := 0
		for t.AddDate(0, 0, i).Month() == t.Month() {
			i++
		}
		t = t.AddDate(0, 0, i-1)
	}

	return t
}
