package util

import (
	"log"
	"math"
	"time"
)

func GetDateFromDateString(format string, dateString string) time.Time {
	date, error := time.Parse(format, dateString)
	if error != nil {
		log.Printf("error formatting date string %s date, %s format , %v error ", dateString, format, error)
	}
	return date
}

func AddMonthsToDate(date time.Time, noOfMonths int) time.Time {
	return date.AddDate(0, noOfMonths, 0)
}

func GetDaysBetweenDates(start time.Time, end time.Time) float64 {
	return math.Ceil(end.Sub(start).Hours() / 24)
}
