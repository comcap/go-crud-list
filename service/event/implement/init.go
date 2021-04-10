package implement

import (
	"meeting-room/service/event"
	"meeting-room/service/util"
	"meeting-room/service/validator"
)

type implementation struct {
	validator validator.Validator
	repo      util.Repository
	uuid      util.UUID
}

func New(validator validator.Validator, repo util.Repository, uuid util.UUID) (service event.Service) {
	return &implementation{validator, repo, uuid}
}
