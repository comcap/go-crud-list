package validator

import (
	"context"
	"fmt"
	"meeting-room/domain"
	"strings"

	"github.com/go-playground/validator/v10"
)

func (v *GoPlayGroundValidator) pageOptionFilterValidation(structLV validator.StructLevel, filter string) {
	fragments := strings.Split(filter, ":")
	if len(fragments) < 2 {
		structLV.ReportError(filter, "filters", "filters", "length", "")
		return
	}

	support := map[string]bool{
		"ne":        true,
		"like":      true,
		"eq":        true,
		"eqInt":     true,
		"isNull":    true,
		"isNotNull": true,
		"id":        true,
	}
	if _, ok := support[fragments[1]]; !ok {
		structLV.ReportError(filter, "filters", "filters", "support", "")
		return
	}

	switch fragments[1] {
	case "isNull":
		if len(fragments) == 2 {
			return
		}
		structLV.ReportError(filter, "filters", "filters", "length", "")
	case "isNotNull":
		if len(fragments) == 2 {
			return
		}
		structLV.ReportError(filter, "filters", "filters", "length", "")
	default:
		if len(fragments) == 3 {
			return
		}
		structLV.ReportError(filter, "filters", "filters", "length", "")
	}
}

func (v *GoPlayGroundValidator) pageOptionSortValidation(structLV validator.StructLevel, sort string) {
	fragments := strings.Split(sort, ":")
	if len(fragments) < 2 {
		structLV.ReportError(sort, "sorts", "sorts", "length", "")
		return
	}

	support := map[string]bool{
		"asc":  true,
		"desc": true,
	}
	if _, ok := support[fragments[1]]; !ok {
		structLV.ReportError(sort, "sorts", "sorts", "support", "")
		return
	}
}

func (v *GoPlayGroundValidator) checkEventIDUnique(ctx context.Context, structLV validator.StructLevel, ID string) (event *domain.Event) {
	if err := v.eventRepo.Read(ctx, []string{fmt.Sprintf("id:eq:%s", ID)}, &domain.Event{}); err == nil {
		structLV.ReportError(ID, "id", "ID", "unique", "")
	}
	return event
}

func (v *GoPlayGroundValidator) checkEventCalendarIDUnique(ctx context.Context, structLV validator.StructLevel, ID string) (event *domain.Event) {
	if err := v.eventRepo.Read(ctx, []string{fmt.Sprintf("calendarId:eq:%s", ID)}, &domain.Event{}); err == nil {
		structLV.ReportError(ID, "calendarId", "CalendarId", "unique", "")
	}
	return event
}
