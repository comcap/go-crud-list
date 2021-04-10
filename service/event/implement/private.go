package implement

import (
	"fmt"
)

//const (
//	betweenFilterFormat = "%s%s:between:%d|%d"
//)

func makeEventIDFilters(eventID string) (filters []string) {
	return []string{
		fmt.Sprintf("calendarId:eq:%s", eventID),
	}
}

//func MakeDurationBetweenString(start, end string, durationFrom, durationTo int64) (filter string) {
//	return fmt.Sprintf(betweenFilterFormat, start, end, durationFrom, durationTo)
//}

func MakeDurationFromFilterString(start int64) (filter string) {
	return fmt.Sprintf("start.dateTime:gte:%d", start)
}

func MakeDurationToFilterString(end int64) (filter string) {
	return fmt.Sprintf("end.dateTime:lte:%d", end)
}
