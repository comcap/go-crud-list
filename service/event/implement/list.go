package implement

import (
	"context"
	"meeting-room/domain"
	"meeting-room/service/event/out"
	"meeting-room/service/util"
)

func (impl *implementation) List(ctx context.Context, opt *domain.PageOption) (total int, items []*out.EventView, err error) {
	if opt.DurationFrom != "" && opt.DurationTo != "" {
		durationFrom, durationTo, err := opt.GetDurationTime()
		if err != nil {
			return 0, nil, util.RepoListErr(err)
		}
		opt.Filters = append(opt.Filters, MakeDurationFromFilterString(durationFrom))
		opt.Filters = append(opt.Filters, MakeDurationToFilterString(durationTo))
	}
	total, records, err := impl.repo.List(ctx, opt, &domain.Event{})
	if err != nil {
		return 0, nil, util.RepoListErr(err)
	}

	items = make([]*out.EventView, len(records))
	for i, record := range records {
		items[i] = out.EventToView(record.(*domain.Event))
	}

	return total, items, nil
}
