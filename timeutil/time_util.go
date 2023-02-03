package timeutil

import (
	"database/sql"
	"github.com/xq-libs/go-utils/stringutil"
	"log"
	"time"
)

const (
	RFCTime     = "15:04:05"
	RFCDate     = "2006-01-02"
	RFCDateTime = "2006-01-02 15:04:05"
)

func TimeToString(t time.Time) string {
	return ToString(t, RFCTime)
}

func DateToString(t time.Time) string {
	return ToString(t, RFCDate)
}

func DateTimeToString(t time.Time) string {
	return ToString(t, RFCDateTime)
}

func ToString(t time.Time, f string) string {
	if t.IsZero() {
		return ""
	} else {
		return t.Format(f)
	}
}

func StringToTime(s string) time.Time {
	return ToTime(s, RFCTime)
}

func StringToDate(s string) time.Time {
	return ToTime(s, RFCDate)
}

func StringToDateTime(s string) time.Time {
	return ToTime(s, RFCDateTime)
}

func ToTime(s string, f string) time.Time {
	if stringutil.IsNotBlank(s) {
		t, err := time.Parse(f, s)
		if err != nil {
			log.Printf("Format time string %s with format %s failure: %v \n", s, f, err)
		}
		return t
	}
	return time.Time{}
}

func TimeToSqlTime(t time.Time) sql.NullTime {
	return sql.NullTime{
		Time:  t,
		Valid: t.IsZero(),
	}
}

func SqlTimeToTime(t sql.NullTime) time.Time {
	if t.Valid {
		return t.Time
	} else {
		return time.Time{}
	}
}
