package eventin

import "meeting-room/domain"

type Attendee struct {
	Email string `json:"email" validate:"required"`
}

func AttendeeToDomain(input []*Attendee) (attendee []*domain.Attendee) {
	attendee = make([]*domain.Attendee, len(input))
	for i, value := range input {
		attendee[i] = InputToAttendee(value)
	}

	return attendee
}

func InputToAttendee(input *Attendee) (attendee *domain.Attendee) {
	return &domain.Attendee{Email: input.Email}
}
