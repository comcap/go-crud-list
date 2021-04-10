package implement

import (
	"context"

	"meeting-room/service/company/companyin"

	"github.com/opentracing/opentracing-go"
)

func (wrp *wrapper) Delete(ctx context.Context, input *companyin.DeleteInput) (err error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "service.company.Delete")
	defer sp.Finish()

	sp.LogKV("id", input.ID)

	err = wrp.service.Delete(ctx, input)

	sp.LogKV("err", err)

	return err
}
