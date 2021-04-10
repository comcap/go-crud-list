package eventin

import "meeting-room/domain"

type Override struct {
	Method  string `json:"method" validate:"required"`
	Minutes string `json:"minutes" validate:"required"`
}

func OverrideToDomain(input []*Override) (override []*domain.Override) {
	override = make([]*domain.Override, len(input))
	for i, value := range input {
		override[i] = InputToOverride(value)
	}

	return override
}

func InputToOverride(input *Override) (attendee *domain.Override) {
	return &domain.Override{
		Method:  input.Method,
		Minutes: input.Minutes,
	}
}
