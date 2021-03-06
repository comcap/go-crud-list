package event

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"

	"meeting-room/app/view"
	"meeting-room/service/event/eventin"
)

// Delete godoc
// @Tags Companies
// @Summary Soft delete a company
// @Description Mark a company with a given company ID as deleted but keep its content
// @param company_id path string true "Company ID"
// @Param X-Authenticated-Userid header string true "User ID"
// @Produce  json
// @Success 200 {object} view.SuccessResp
// @Success 204 {object} view.SuccessResp
// @Success 401 {object} view.ErrResp
// @Success 403 {object} view.ErrResp
// @Success 422 {object} view.ErrResp
// @Success 500 {object} view.ErrResp
// @Success 503 {object} view.ErrResp
// @Router /companies/{company_id} [delete]
func (ctrl *Controller) Delete(c *gin.Context) {
	span, ctx := opentracing.StartSpanFromContextWithTracer(
		c.Request.Context(),
		opentracing.GlobalTracer(),
		"handler.event.Delete",
	)
	defer span.Finish()

	input := &eventin.DeleteInput{
		ID: c.Param("id"),
	}

	err := ctrl.service.Delete(ctx, input)
	if err != nil {
		view.MakeErrResp(c, err)
		return
	}

	view.MakeSuccessResp(c, http.StatusOK, nil)
}
