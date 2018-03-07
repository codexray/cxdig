package repos

import (
	"codexray/cxdig/types"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBringToLastMoment(t *testing.T) {
	tInitial := time.Date(2000, time.Month(5), 23, 15, 32, 20, 0, &time.Location{})
	tDay := bringToLastMoment(tInitial, FreqDay)
	assert.Equal(t, time.Date(tInitial.Year(), tInitial.Month(), tInitial.Day(), 23, 59, 59, 0, &time.Location{}), tDay)

	tMonth := bringToLastMoment(tInitial, FreqMonth)
	assert.Equal(t, time.Date(tInitial.Year(), tInitial.Month(), 31, 23, 59, 59, 0, &time.Location{}), tMonth)

	tQuarter := bringToLastMoment(tInitial, FreqQuarter)
	assert.Equal(t, time.Date(tInitial.Year(), time.Month(6), 30, 23, 59, 59, 0, &time.Location{}), tQuarter)

	tYear := bringToLastMoment(tInitial, FreqYear)
	assert.Equal(t, time.Date(tInitial.Year(), time.Month(12), 31, 23, 59, 59, 0, &time.Location{}), tYear)
}

func TestGetCommitByStep(t *testing.T) {
	commits := []types.CommitInfo{}
	var incr time.Duration
	incr = time.Duration(0)
	for i := 0; i < 366; i++ {
		commit1 := types.CommitInfo{
			Number:   i * 2,
			DateTime: time.Date(2001, time.Month(1), 1, 8, 0, 0, 0, &time.Location{}).Add(incr),
		}
		incr += time.Hour * time.Duration(12)
		commit2 := types.CommitInfo{
			Number:   (i * 2) + 1,
			DateTime: time.Date(2001, time.Month(1), 1, 8, 0, 0, 0, &time.Location{}).Add(incr),
		}
		incr += time.Hour * time.Duration(12)
		commits = append(commits, commit1, commit2)
	}

	commits = SortCommitByDateDecr(commits)

	freq := SamplingFreq{
		Value: 1,
		Unit:  FreqDay,
	}
	commits2 := FilterCommitsByStep(commits, freq, 0)
	assert.Equal(t, 366, len(commits2))

	freq.Value = 10
	commits2 = FilterCommitsByStep(commits, freq, 0)
	assert.Equal(t, 37, len(commits2))

	freq.Unit = FreqMonth
	freq.Value = 1
	commits2 = FilterCommitsByStep(commits, freq, 0)
	assert.Equal(t, 13, len(commits2))

	freq.Value = 2
	commits2 = FilterCommitsByStep(commits, freq, 0)
	assert.Equal(t, 7, len(commits2))

	freq.Unit = FreqYear
	freq.Value = 1
	commits2 = FilterCommitsByStep(commits, freq, 0)
	assert.Equal(t, 2, len(commits2))

	freq.Value = 2
	commits2 = FilterCommitsByStep(commits, freq, 0)
	assert.Equal(t, 1, len(commits2))

	freq.Unit = FreqQuarter
	freq.Value = 1
	commits2 = FilterCommitsByStep(commits, freq, 0)
	assert.Equal(t, 5, len(commits2))

	freq.Value = 2
	commits2 = FilterCommitsByStep(commits, freq, 0)
	assert.Equal(t, 3, len(commits2))

	freq.Unit = FreqCommit
	freq.Value = 10
	commits2 = FilterCommitsByStep(commits, freq, 0)
	assert.Equal(t, 74, len(commits2))

	freq.Value = 2
	commits2 = FilterCommitsByStep(commits, freq, 0)
	assert.Equal(t, 366, len(commits2))

	temp := []types.CommitInfo{}
	for _, com := range commits {
		if com.DateTime.Weekday() == time.Tuesday && com.DateTime.Month() != time.April && com.DateTime.Month() != time.May && com.DateTime.Month() != time.June {
			temp = append(temp, com)
		}
	}
	commits = temp

	freq = SamplingFreq{
		Value: 1,
		Unit:  FreqDay,
	}
	commits2 = FilterCommitsByStep(commits, freq, 0)
	assert.Equal(t, 365, len(commits2))

	freq.Value = 10
	commits2 = FilterCommitsByStep(commits, freq, 0)
	assert.Equal(t, 37, len(commits2))

	freq.Unit = FreqMonth
	freq.Value = 1
	commits2 = FilterCommitsByStep(commits, freq, 0)
	assert.Equal(t, 13, len(commits2))

	freq.Value = 2
	commits2 = FilterCommitsByStep(commits, freq, 0)
	assert.Equal(t, 7, len(commits2))

	freq.Unit = FreqYear
	freq.Value = 1
	commits2 = FilterCommitsByStep(commits, freq, 0)
	assert.Equal(t, 2, len(commits2))

	freq.Value = 2
	commits2 = FilterCommitsByStep(commits, freq, 0)
	assert.Equal(t, 1, len(commits2))

	freq.Unit = FreqQuarter
	freq.Value = 1
	commits2 = FilterCommitsByStep(commits, freq, 0)
	assert.Equal(t, 5, len(commits2))

	freq.Value = 2
	commits2 = FilterCommitsByStep(commits, freq, 0)
	assert.Equal(t, 3, len(commits2))

	freq.Unit = FreqCommit
	freq.Value = 10
	commits2 = FilterCommitsByStep(commits, freq, 0)
	assert.Equal(t, 8, len(commits2))

	freq.Value = 2
	commits2 = FilterCommitsByStep(commits, freq, 0)
	assert.Equal(t, 40, len(commits2))

}
