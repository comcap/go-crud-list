package eventin

import (
	"github.com/uniplaces/carbon"
	"meeting-room/domain"
)

type UpdateInput struct {
	ID          string      `json:"id" validate:"required"`
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

//func MakeTestUpdateInput() (input *UpdateInput) {
//	return &UpdateInput{
//		ID:   "test",
//		Name: "test",
//	}
//}

func UpdateInputToEventDomain(input *UpdateInput) (company *domain.Event) {
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
		Updated:   carbon.Now().Unix(),
	}
}
