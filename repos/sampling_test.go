package repos

import (
	"codexray/cxdig/types"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBringToLastMoment(t *testing.T) {
	tInitial := time.Date(2000, time.Month(5), 23, 15, 32, 20, 0, &time.Location{})
	tDay := bringToLastMoment(tInitial, RateDay)
	assert.Equal(t, time.Date(tInitial.Year(), tInitial.Month(), tInitial.Day(), 23, 59, 59, 0, &time.Location{}), tDay)

	tMonth := bringToLastMoment(tInitial, RateMonth)
	assert.Equal(t, time.Date(tInitial.Year(), tInitial.Month(), 31, 23, 59, 59, 0, &time.Location{}), tMonth)

	tQuarter := bringToLastMoment(tInitial, RateQuarter)
	assert.Equal(t, time.Date(tInitial.Year(), time.Month(6), 30, 23, 59, 59, 0, &time.Location{}), tQuarter)

	tYear := bringToLastMoment(tInitial, RateYear)
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

	rate := SamplingRate{
		Value: 1,
		Unit:  RateDay,
	}
	commits2 := FilterCommitsByStep(commits, rate, 0)
	assert.Equal(t, 366, len(commits2))

	rate.Value = 10
	commits2 = FilterCommitsByStep(commits, rate, 0)
	assert.Equal(t, 37, len(commits2))

	rate.Unit = RateMonth
	rate.Value = 1
	commits2 = FilterCommitsByStep(commits, rate, 0)
	assert.Equal(t, 13, len(commits2))

	rate.Value = 2
	commits2 = FilterCommitsByStep(commits, rate, 0)
	assert.Equal(t, 7, len(commits2))

	rate.Unit = RateYear
	rate.Value = 1
	commits2 = FilterCommitsByStep(commits, rate, 0)
	assert.Equal(t, 2, len(commits2))

	rate.Value = 2
	commits2 = FilterCommitsByStep(commits, rate, 0)
	assert.Equal(t, 1, len(commits2))

	rate.Unit = RateQuarter
	rate.Value = 1
	commits2 = FilterCommitsByStep(commits, rate, 0)
	assert.Equal(t, 5, len(commits2))

	rate.Value = 2
	commits2 = FilterCommitsByStep(commits, rate, 0)
	assert.Equal(t, 3, len(commits2))

	rate.Unit = RateCommit
	rate.Value = 10
	commits2 = FilterCommitsByStep(commits, rate, 0)
	assert.Equal(t, 74, len(commits2))

	rate.Value = 2
	commits2 = FilterCommitsByStep(commits, rate, 0)
	assert.Equal(t, 366, len(commits2))

	temp := []types.CommitInfo{}
	for _, com := range commits {
		if com.DateTime.Weekday() == time.Tuesday && com.DateTime.Month() != time.April && com.DateTime.Month() != time.May && com.DateTime.Month() != time.June {
			temp = append(temp, com)
		}
	}
	commits = temp

	rate = SamplingRate{
		Value: 1,
		Unit:  RateDay,
	}
	commits2 = FilterCommitsByStep(commits, rate, 0)
	assert.Equal(t, 365, len(commits2))

	rate.Value = 10
	commits2 = FilterCommitsByStep(commits, rate, 0)
	assert.Equal(t, 37, len(commits2))

	rate.Unit = RateMonth
	rate.Value = 1
	commits2 = FilterCommitsByStep(commits, rate, 0)
	assert.Equal(t, 13, len(commits2))

	rate.Value = 2
	commits2 = FilterCommitsByStep(commits, rate, 0)
	assert.Equal(t, 7, len(commits2))

	rate.Unit = RateYear
	rate.Value = 1
	commits2 = FilterCommitsByStep(commits, rate, 0)
	assert.Equal(t, 2, len(commits2))

	rate.Value = 2
	commits2 = FilterCommitsByStep(commits, rate, 0)
	assert.Equal(t, 1, len(commits2))

	rate.Unit = RateQuarter
	rate.Value = 1
	commits2 = FilterCommitsByStep(commits, rate, 0)
	assert.Equal(t, 5, len(commits2))

	rate.Value = 2
	commits2 = FilterCommitsByStep(commits, rate, 0)
	assert.Equal(t, 3, len(commits2))

	rate.Unit = RateCommit
	rate.Value = 10
	commits2 = FilterCommitsByStep(commits, rate, 0)
	assert.Equal(t, 8, len(commits2))

	rate.Value = 2
	commits2 = FilterCommitsByStep(commits, rate, 0)
	assert.Equal(t, 40, len(commits2))

}
