package domain

import "strconv"

type Paginator struct {
	Total   int
	CurPage int
	PerPage int
	HasMore bool
}

type PageOption struct {
	Page         int      `form:"page" validate:"min=0"`
	PerPage      int      `form:"per_page" validate:"min=0"`
	Filters      []string `form:"filters"`
	Sorts        []string `form:"sorts"`
	DurationFrom string   `form:"duration_from"`
	DurationTo   string   `form:"duration_to"`
}

type SetOpParam struct {
	Filters      []string
	SetFieldName string
	Item         interface{}
}

func (opt *PageOption) GetDurationTime() (durationFrom, durationTo int64, err error) {
	durationFrom, err = strconv.ParseInt(opt.DurationFrom, 10, 64)
	if err != nil {
		return 0, 0, err
	}

	durationTo, err = strconv.ParseInt(opt.DurationTo, 10, 64)
	if err != nil {
		return 0, 0, err
	}

	return durationFrom, durationTo, err
}
