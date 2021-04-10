package implement

import (
	"context"
	"meeting-room/service/event/eventin"

	"meeting-room/domain"
	"meeting-room/service/util"
)

func (impl *implementation) Update(ctx context.Context, input *eventin.UpdateInput) (err error) {
	err = impl.validator.Validate(input)
	if err != nil {
		return util.ValidationUpdateErr(err)
	}

	filters := makeEventIDFilters(input.ID)

	event := &domain.Event{}
	err = impl.repo.Read(ctx, filters, event)
	if err != nil {
		return util.RepoReadErr(err)
	}

	update := eventin.UpdateInputToEventDomain(input)
	update.CalendarId = event.CalendarId
	update.Created = event.Created
	err = impl.repo.Update(ctx, filters, update)
	if err != nil {
		return util.RepoUpdateErr(err)
	}

	return nil
}
