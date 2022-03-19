package restful

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type sendParams struct {
	Count  uint64 `json:"count" binding:"required"`
	Amount uint64 `json:"amount" binding:"required"`
}

func (h *handler) Send(c *gin.Context) {
	var params sendParams
	if err := c.BindJSON(&params); err != nil {
		fmt.Println(err)
		return
	}
	userId := uint64(1)
	result, err := h.redPacketSvc.Send(userId, params.Count, params.Amount)
	if !result || err != nil {
		fmt.Println(err)
	}
	c.Status(http.StatusNoContent)
	return
}

type grabParams struct {
	ActivityId uint64 `json:"activity_id" binding:"required"`
}

func (h *handler) Grab(c *gin.Context) {
	var params grabParams
	if err := c.BindJSON(&params); err != nil {
		fmt.Println(err)
		return
	}
	userId := uint64(1)
	amount, err := h.redPacketSvc.Grab(userId, params.ActivityId)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(amount)
	c.Status(http.StatusNoContent)
	return
}
