package implement

import (
	"context"
	"meeting-room/domain"
	"meeting-room/service/event/eventin"
	"meeting-room/service/util"
)

func (impl *implementation) Delete(ctx context.Context, input *eventin.DeleteInput) (err error) {
	event := &domain.Event{}
	filters := makeEventIDFilters(input.ID)

	err = impl.repo.Read(ctx, filters, event)
	if err != nil {
		return util.RepoReadErr(err)
	}

	err = impl.repo.Delete(ctx, filters)
	if err != nil {
		return util.RepoDeleteErr(err)
	}

	return nil
}
