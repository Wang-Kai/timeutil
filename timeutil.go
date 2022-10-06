package timeutil

import "time"

type TimeUnit uint8

const (
	Y TimeUnit = iota
	M
	D
	H
)

func StartOf(t time.Time, unit TimeUnit) time.Time {
	var nt time.Time
	switch unit {
	case Y:
		y := t.Year()
		nt = time.Date(y, 1, 1, 0, 0, 0, 0, t.Location())
	case M:
		y, m := t.Year(), t.Month()
		nt = time.Date(y, m, 1, 0, 0, 0, 0, t.Location())
	case D:
		y, m, d := t.Year(), t.Month(), t.Day()
		nt = time.Date(y, m, d, 0, 0, 0, 0, t.Location())
	case H:
		y, m, d, h := t.Year(), t.Month(), t.Day(), t.Hour()
		nt = time.Date(y, m, d, h, 0, 0, 0, t.Location())
	}
	return nt
}

func EndOf(t time.Time, unit TimeUnit) time.Time {
	var nt time.Time
	switch unit {
	case Y:
		nt = StartOf(t, Y).AddDate(1, 0, 0).Add(time.Nanosecond * -1)
	case M:
		nt = StartOf(t, M).AddDate(0, 1, 0).Add(time.Nanosecond * -1)
	case D:
		nt = StartOf(t, D).AddDate(0, 0, 1).Add(time.Nanosecond * -1)
	case H:
		nt = StartOf(t, H).Add(time.Hour - time.Nanosecond)
	}
	return nt
}
