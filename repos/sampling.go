package repos

import (
	"codexray/cxdig/types"
	"time"
)

func FilterCommitsByStep(commits []types.CommitInfo, rate SamplingRate, limit int) []types.SampleInfo {
	if limit == 0 {
		limit = len(commits)
	}
	rtn := []types.SampleInfo{}
	sampleNumber := 1
	if rate.Unit == RateCommit {
		for i := 0; i < limit*rate.Value && i < len(commits); i += rate.Value {
			rtn = append(rtn, types.SampleInfo{
				Number:   types.SampleID(sampleNumber),
				DateTime: commits[i].DateTime,
				CommitID: commits[i].CommitID,
			})
			sampleNumber++
		}
	} else {
		step := bringToLastMoment(commits[0].DateTime, rate.Unit)
		j := 0
		for i := 0; j <= limit && i < len(commits); i++ {
			if commits[i].DateTime.Before(step) {
				var t time.Time
				firstAdded := 0
				for t = step; commits[i].DateTime.Before(t); t = getNextStep(t, rate) {
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

func getNextStep(step time.Time, rate SamplingRate) time.Time {
	switch rate.Unit {
	case RateDay:
		step = step.AddDate(0, 0, -rate.Value)
	case RateWeek:
		step = step.AddDate(0, 0, -rate.Value*8)
	case RateMonth:
		for i := 0; i < rate.Value; i++ {
			temp := step.Month()
			step = step.AddDate(0, -1, 0)
			for step.Month() == temp {
				step = step.AddDate(0, 0, -1)
			}
		}
	case RateQuarter:
		for i := 0; i < rate.Value*3; i++ {
			temp := step.Month()
			step = step.AddDate(0, -1, 0)
			for step.Month() == temp {
				step = step.AddDate(0, 0, -1)
			}
		}
	case RateYear:
		step = step.AddDate(-rate.Value, 0, 0)
	}
	return bringToLastMoment(step, rate.Unit)
}

func bringToLastMoment(t time.Time, rateUnit RateUnit) time.Time {
	t = t.Add((time.Hour * time.Duration(23-t.Hour())) + (time.Minute * time.Duration(59-t.Minute())) + (time.Second * time.Duration(59-t.Second())))

	if rateUnit == RateWeek {
		i := 0
		for t.AddDate(0, 0, i).Weekday() == time.Weekday(7) {
			i++
		}
		t = t.AddDate(0, 0, i)
	}

	if rateUnit == RateYear {
		i := 0
		for t.AddDate(0, i, 0).Month() != time.Month(12) {
			i++
		}
		t = t.AddDate(0, i, 0)
	}

	if rateUnit == RateQuarter {
		i := 0
		for t.AddDate(0, i, 0).Month()%3 != 0 {
			i++
		}
		t = t.AddDate(0, i, 0)
	}

	if rateUnit == RateMonth || rateUnit == RateQuarter || rateUnit == RateYear {
		i := 0
		for t.AddDate(0, 0, i).Month() == t.Month() {
			i++
		}
		t = t.AddDate(0, 0, i-1)
	}

	return t
}
