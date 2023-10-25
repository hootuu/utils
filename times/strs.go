package times

import "time"

const (
	DateTimeLayout  = "2006-01-02"
	TimestampLayout = "2006-01-02 15:04:05"
)

func ToDate(t time.Time) string {
	return t.Format(DateTimeLayout)
}

func ToTimestamp(t time.Time) string {
	return t.Format(TimestampLayout)
}
