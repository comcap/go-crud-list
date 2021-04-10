package validator

import (
	"context"
	"meeting-room/service/event/eventin"

	"github.com/go-playground/validator/v10"
)

func (v *GoPlayGroundValidator) EventCreateStructLevelValidation(structLV validator.StructLevel) {
	ctx := context.Background()
	event := structLV.Current().Interface().(eventin.CreateInput)

	v.checkEventIDUnique(ctx, structLV, event.ID)
	v.checkEventCalendarIDUnique(ctx, structLV, event.CalendarId)
}
