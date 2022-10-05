package bqproto

import "time"

const (
	bitShiftSecond = 20
	bitShiftMinute = 26
	bitShiftHour   = 32
	bitShiftDay    = 37
	bitShiftMonth  = 42
	bitShiftYear   = 46
)

func encodeTime(t time.Time) uint64 {
	return uint64(t.Hour()<<bitShiftHour | t.Minute()<<bitShiftMinute | t.Second()<<bitShiftSecond)
}

func encodeTimeFraction(t time.Time) uint64 {
	return encodeTime(t) | uint64(t.Nanosecond()/1000)
}

func encodeDatetime(t time.Time) uint64 {
	return uint64(t.Year()<<bitShiftYear|int(t.Month())<<bitShiftMonth|t.Day()<<bitShiftDay) | encodeTime(t)
}

func encodeDatetimeFraction(t time.Time) uint64 {
	return encodeDatetime(t) | uint64(t.Nanosecond()/1000)
}
