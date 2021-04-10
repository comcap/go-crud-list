package implement

import (
	"context"

	"meeting-room/domain"
	"meeting-room/service/event/eventin"
	"meeting-room/service/event/out"
	"meeting-room/service/util"
)

func (impl *implementation) Read(ctx context.Context, input *eventin.ReadInput) (view *out.EventView, err error) {
	event := &domain.Event{}
	filters := makeEventIDFilters(input.EventID)

	err = impl.repo.Read(ctx, filters, event)
	if err != nil {
		return nil, util.RepoReadErr(err)
	}

	return out.EventToView(event), nil
}
