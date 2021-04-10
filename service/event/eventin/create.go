package eventin

import (
	"github.com/uniplaces/carbon"
	"meeting-room/domain"
)

type CreateInput struct {
	ID          string      `json:"id"`
	CalendarId  string      `json:"calendarId" validate:"required"`
	HTMLLink    string      `json:"htmlLink" validate:"required"`
	Summary     string      `json:"summary" validate:"required"`
	Description string      `json:"description" validate:"required"`
	Start       DateTime    `json:"start" validate:"required"`
	End         DateTime    `json:"end" validate:"required"`
	Rooms       string      `json:"rooms" validate:"required"`
	Reminders   Reminders   `json:"reminders" validate:"required"`
	Attendees   []*Attendee `json:"attendees" validate:"required"`
	Creator     Creator     `json:"creator" validate:"required"`
	Organizer   Creator     `json:"organizer" validate:"required"`
} // @Name EventCreateInput
type Creator struct {
	Email string `json:"email" validate:"required"`
	Self  bool   `json:"self" validate:"required"`
}

type DateTime struct {
	DateTime int64 `json:"dateTime" validate:"required"`
}

type Reminders struct {
	Override []*Override `json:"override" validate:"required"`
}

func CreateInputToEventDomain(input *CreateInput) (event *domain.Event) {
	return &domain.Event{
		ID:          input.ID,
		CalendarId:  input.CalendarId,
		HTMLLink:    input.HTMLLink,
		Summary:     input.Summary,
		Description: input.Description,
		Start:       domain.DateTime{DateTime: input.Start.DateTime},
		End:         domain.DateTime{DateTime: input.End.DateTime},
		Rooms:       input.Rooms,
		Reminders:   domain.Reminders{Override: OverrideToDomain(input.Reminders.Override)},
		Creator: domain.Creator{
			Email: input.Creator.Email,
			Self:  input.Creator.Self,
		},
		Organizer: domain.Creator{
			Email: input.Organizer.Email,
			Self:  input.Organizer.Self,
		},
		Attendees: AttendeeToDomain(input.Attendees),
		Created:   carbon.Now().Unix(),
		Updated:   0,
	}
}

//func MakeTestCreateInput() (input *CreateInput) {
//	return &CreateInput{
//		ID:   "test",
//		Name: "test",
//	}
//}
//
