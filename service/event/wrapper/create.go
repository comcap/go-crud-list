package implement

import (
	"context"

	"meeting-room/service/company/companyin"

	"github.com/opentracing/opentracing-go"
)

func (wrp *wrapper) Create(ctx context.Context, input *companyin.CreateInput) (ID string, err error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "service.company.Create")
	defer sp.Finish()

	sp.LogKV("ID", input.ID)
	sp.LogKV("Name", input.Name)

	ID, err = wrp.service.Create(ctx, input)

	sp.LogKV("ID", ID)
	sp.LogKV("err", err)

	return ID, err
}
