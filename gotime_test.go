package gotime

import (
	"testing"

	"time"
)

var (
	now    = time.Date(2022, 10, 6, 21, 4, 34, 0, time.UTC)
	format = "2006-01-02 15:04:05"
)

func TestStartOf(t *testing.T) {
	assert := assertT(t)

	assert("StartOf_Y", "2022-01-01 00:00:00", StartOf(now, Y))
	assert("StartOf_M", "2022-10-01 00:00:00", StartOf(now, M))
	assert("StartOf_D", "2022-10-06 00:00:00", StartOf(now, D))
	assert("StartOf_H", "2022-10-06 21:00:00", StartOf(now, H))
}

func TestEndOf(t *testing.T) {
	assert := assertT(t)

	assert("EndOf_Y", "2022-12-31 23:59:59", EndOf(now, Y))
	assert("EndOf_M", "2022-10-31 23:59:59", EndOf(now, M))
	assert("EndOf_D", "2022-10-06 23:59:59", EndOf(now, D))
	assert("EndOf_H", "2022-10-06 21:59:59", EndOf(now, H))
}

func assertT(t *testing.T) func(fn string, expected string, t time.Time) {
	return func(fn, expected string, actual time.Time) {
		actutalStr := actual.Format(format)
		if actutalStr != expected {
			t.Fatalf("Func %s\t Expect %s, but got: %s", fn, expected, actutalStr)
		}
	}
}
