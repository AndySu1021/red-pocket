package restful

import (
	"fmt"
	"net/http"
	"red-packet/pkg/model/dto"
	"red-packet/pkg/model/option"
	"time"

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

	activity := dto.Activity{
		Count:     params.Count,
		Amount:    params.Amount,
		CreatedBy: userId,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	result, err := h.redPacketSvc.Send(c.Request.Context(), &activity)
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

	redPacketOpt := option.RedPacketOption{
		RedPacket: dto.RedPacket{
			ActivityID: params.ActivityId,
			Status:     1,
		},
	}

	userOpt := option.UserOption{
		User: dto.User{
			ID: userId,
		},
	}

	amount, err := h.redPacketSvc.Grab(c.Request.Context(), &redPacketOpt, &userOpt)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(amount)
	c.Status(http.StatusNoContent)
	return
}
