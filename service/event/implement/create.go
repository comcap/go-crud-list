package implement

import (
	"context"
	"meeting-room/service/event/eventin"
	"meeting-room/service/util"
)

func (impl *implementation) Create(ctx context.Context, input *eventin.CreateInput) (ID string, err error) {
	err = impl.validator.Validate(input)
	if err != nil {
		return "", util.ValidationCreateErr(err)
	}
	input.ID = impl.uuid.Generate()
	event := eventin.CreateInputToEventDomain(input)

	_, err = impl.repo.Create(ctx, event)
	if err != nil {
		return "", util.RepoCreateErr(err)
	}

	return event.ID, nil
}
