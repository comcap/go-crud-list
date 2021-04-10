package event

import (
	"context"
	"meeting-room/domain"
	"meeting-room/service/event/eventin"
	"meeting-room/service/event/out"
)

//go:generate mockery --name=Service
type Service interface {
	List(ctx context.Context, opt *domain.PageOption) (total int, items []*out.EventView, err error)
	Create(ctx context.Context, input *eventin.CreateInput) (ID string, err error)
	Read(ctx context.Context, input *eventin.ReadInput) (event *out.EventView, err error)
	Update(ctx context.Context, input *eventin.UpdateInput) (err error)
	Delete(ctx context.Context, input *eventin.DeleteInput) (err error)
}
