package times

import "time"

func ToDate(t time.Time) string {
	return t.Format("2006-01-02")
}

func ToTimestamp(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}
