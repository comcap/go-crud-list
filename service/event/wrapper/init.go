package implement

import (
	"meeting-room/service/company"
)

type wrapper struct {
	service company.Service
}

func _(service company.Service) company.Service {
	return &wrapper{
		service: service,
	}
}
