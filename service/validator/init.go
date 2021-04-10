package validator

import (
	"meeting-room/service/util"

	"meeting-room/service/event/eventin"

	"github.com/go-playground/validator/v10"

	"meeting-room/domain"
)

type GoPlayGroundValidator struct {
	validate    *validator.Validate
	companyRepo util.Repository
	staffRepo   util.Repository
	eventRepo   util.Repository
}

func New(eventRepo util.Repository) (v *GoPlayGroundValidator) {
	v = &GoPlayGroundValidator{
		validate:  validator.New(),
		eventRepo: eventRepo,
	}

	v.validate.RegisterStructValidation(v.PageOptionStructLevelValidation, &domain.PageOption{})
	v.validate.RegisterStructValidation(v.EventCreateStructLevelValidation, &eventin.CreateInput{})

	return v
}

func (v *GoPlayGroundValidator) Validate(item interface{}) (err error) {
	return v.validate.Struct(item)
}
