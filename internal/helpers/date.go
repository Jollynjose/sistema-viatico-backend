package helpers

import "time"

func FormatTimeToYYYYMMDD(t time.Time) string {
	return t.Format("2006-01-02")
}

func FormatTimeToHHMMAMOrPM(t time.Time) string {
	return t.Format("03:04 PM")
}
