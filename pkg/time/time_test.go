package time

import (
	"testing"
	"time"
)

const hoursInADay = 24
const hoursInAWeek = hoursInADay * 7

func TestDay(t *testing.T) {
	timeNow := time.Now()

	timeDay := timeNow.Add(Day)
	timeDayWithHours := timeNow.Add(hoursInADay * time.Hour)

	if timeDay != timeDayWithHours {
		t.Errorf("Day duration seems invalid. got: %s, want: %s", timeDay.String(), timeDayWithHours.String())
	}
}

func TestWeek(t *testing.T) {
	timeNow := time.Now()

	timeWeek := timeNow.Add(Week)
	timeWeekWithHours := timeNow.Add(hoursInAWeek * time.Hour)

	if timeWeek != timeWeekWithHours {
		t.Errorf("Day duration seems invalid. got: %s, want: %s", timeWeek.String(), timeWeekWithHours.String())
	}
}
