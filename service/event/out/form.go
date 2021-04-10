package out

import (
	"meeting-room/domain"
)

type EventView struct {
	CalendarID string   `json:"calendarId"`
	Status     string   `json:"status"`
	HTMLLink   string   `json:"htmlLink"`
	Created    int64    `json:"created"`
	Updated    int64    `json:"updated"`
	Summary    string   `json:"summary"`
	Creator    Creator  `json:"creator"`
	Organizer  Creator  `json:"organizer"`
	Start      DateTime `json:"start"`
	End        DateTime `json:"end"`
} // @Name EventView

type Creator struct {
	Email string `json:"email"`
	Self  bool   `json:"self"`
}

type DateTime struct {
	DateTime int64 `json:"dateTime"`
}

func EventToView(event *domain.Event) (view *EventView) {
	return &EventView{
		CalendarID: event.CalendarId,
		HTMLLink:   event.HTMLLink,
		Created:    event.Created,
		Updated:    event.Updated,
		Summary:    event.Summary,
		Creator: Creator{
			Email: event.Creator.Email,
			Self:  event.Creator.Self,
		},
		Organizer: Creator{
			Email: event.Organizer.Email,
			Self:  event.Organizer.Self,
		},
		Start: DateTime{
			DateTime: event.Start.DateTime,
		},
		End: DateTime{
			DateTime: event.End.DateTime,
		},
	}
}
