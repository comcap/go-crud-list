package event

import (
	"meeting-room/service/event"
)

type Controller struct {
	service event.Service
}

func New(service event.Service) (ctrl *Controller) {
	return &Controller{service}
}
