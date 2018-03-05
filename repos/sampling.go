package repos

import (
	"codexray/cxdig/types"
	"time"
)

func FilterCommitsByStep(commits []types.CommitInfo, freq SamplingFreq, limit int) []types.CommitInfo {
	commits = SortCommitByDateDecr(commits)
	if limit == 0 {
		limit = len(commits)
	}
	rtn := []types.CommitInfo{}
	if freq.Unit == FreqCommit {
		for i := 0; i < limit*freq.Value && i < len(commits); i += freq.Value {
			rtn = append(rtn, commits[i])
		}
	} else {
		t := bringToLastMoment(commits[0].DateTime, freq.Unit)
		j := 0
		for i := 0; j <= limit && i < len(commits); i++ {
			if commits[i].DateTime.Before(t) {
				rtn = append(rtn, commits[i])
				j++
				switch freq.Unit {
				case FreqDay:
					t = commits[i].DateTime.AddDate(0, 0, -freq.Value)
				case FreqWeek:
					t = commits[i].DateTime.AddDate(0, 0, -freq.Value*8)
				case FreqMonth:
					t = commits[i].DateTime
					for i := 0; i < freq.Value; i++ {
						temp := t.Month()
						t = t.AddDate(0, -1, 0)
						for t.Month() == temp {
							t = t.AddDate(0, 0, -1)
						}
					}
				case FreqQuarter:
					t = commits[i].DateTime
					for i := 0; i < freq.Value*3; i++ {
						temp := t.Month()
						t = t.AddDate(0, -1, 0)
						for t.Month() == temp {
							t = t.AddDate(0, 0, -1)
						}
					}
				case FreqYear:
					t = commits[i].DateTime.AddDate(-freq.Value, 0, 0)
				}
				t = bringToLastMoment(t, freq.Unit)
			}
		}
	}
	return rtn
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
