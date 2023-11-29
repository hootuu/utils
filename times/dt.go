package times

import "time"

func GetToday(t time.Time) time.Time {
	now := time.Now()
	yesterday := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	return yesterday
}

func GetTomorrow(t time.Time) time.Time {
	return t.AddDate(0, 0, 1)
}

func GetYesterday(t time.Time) time.Time {
	now := time.Now()
	yesterday := time.Date(now.Year(), now.Month(), now.Day()-1, 0, 0, 0, 0, now.Location())
	return yesterday
}

func GetDtRange(t time.Time) (time.Time, time.Time) {
	start := time.Date(
		t.Year(), t.Month(), t.Day(),
		0, 0, 0, 0,
		t.Location(),
	)
	end := time.Date(
		t.Year(), t.Month(), t.Day(),
		23, 59, 59, int(time.Second-time.Nanosecond),
		t.Location(),
	)
	return start, end
}

func GetWeekRange(t time.Time) (time.Time, time.Time) {
	firstDay := t.AddDate(0, 0, -int(t.Weekday()))
	lastDay := firstDay.AddDate(0, 0, 6)
	start := time.Date(
		firstDay.Year(), firstDay.Month(), firstDay.Day(),
		0, 0, 0, 0, t.Location(),
	)
	end := time.Date(lastDay.Year(), lastDay.Month(), lastDay.Day(),
		23, 59, 59, int(time.Second-time.Nanosecond), t.Location(),
	)
	return start, end
}

func GetMonthRange(t time.Time) (time.Time, time.Time) {
	firstDay := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
	nextMonth := firstDay.AddDate(0, 1, 0)
	lastDay := nextMonth.AddDate(0, 0, -1)
	start := time.Date(
		firstDay.Year(), firstDay.Month(), firstDay.Day(),
		0, 0, 0, 0, t.Location(),
	)
	end := time.Date(
		lastDay.Year(), lastDay.Month(), lastDay.Day(),
		23, 59, 59, int(time.Second-time.Nanosecond), t.Location(),
	)
	return start, end
}

func GetQuarterRange(t time.Time) (time.Time, time.Time) {
	quarterStartMonth := time.Month(((t.Month()-1)/3)*3 + 1)
	firstDay := time.Date(t.Year(), quarterStartMonth, 1, 0, 0, 0, 0, t.Location())
	nextQuarter := firstDay.AddDate(0, 3, 0)
	lastDay := nextQuarter.AddDate(0, 0, -1)
	start := time.Date(firstDay.Year(), firstDay.Month(), firstDay.Day(),
		0, 0, 0, 0, t.Location())
	end := time.Date(lastDay.Year(), lastDay.Month(), lastDay.Day(),
		23, 59, 59, int(time.Second-time.Nanosecond), t.Location())
	return start, end
}

func GetYearRange(t time.Time) (time.Time, time.Time) {
	firstDay := time.Date(t.Year(), time.January, 1, 0, 0, 0, 0, t.Location())
	nextYear := firstDay.AddDate(1, 0, 0)
	lastDay := nextYear.AddDate(0, 0, -1)
	start := time.Date(firstDay.Year(), firstDay.Month(), firstDay.Day(),
		0, 0, 0, 0, t.Location())
	end := time.Date(lastDay.Year(), lastDay.Month(), lastDay.Day(),
		23, 59, 59, int(time.Second-time.Nanosecond), t.Location())
	return start, end
}
