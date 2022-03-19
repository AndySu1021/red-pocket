package red_packet

import (
	"context"
	"fmt"
	"math/rand"
	"red-packet/pkg/model/dto"
	"red-packet/pkg/model/option"
	"strconv"
	"time"
)

func (s *service) Send(userId uint64, count uint64, amount uint64) (bool, error) {
	activity := &dto.Activity{
		Count:     count,
		Amount:    amount,
		CreatedBy: userId,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// TODO: should use transaction to promise atomic
	err := s.repo.Create(context.Background(), nil, activity)
	if err != nil {
		return false, err
	}

	var redPacketIds []uint64

	remain := amount
	sum := uint64(0)
	for i := uint64(0); i < count; i++ {
		x := doubleAverage(count-i, remain)
		remain -= x
		sum += x

		tmp := &dto.RedPacket{
			ActivityID: activity.ID,
			UserID:     0,
			Amount:     x,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		}

		err := s.repo.Create(context.Background(), nil, tmp)
		if err != nil {
			return false, err
		}

		redPacketIds = append(redPacketIds, tmp.ID)
	}

	// TODO: need to promise atomic
	for _, v := range redPacketIds {
		total, err := s.cache.LPush(context.Background(), getRedisKey(activity.ID), v)
		if err != nil {
			return false, err
		}

		fmt.Println(total)
	}

	return true, nil
}

func (s *service) Grab(userId uint64, activityId uint64) (amount uint64, err error) {
	amount = uint64(10)

	result, err := s.cache.RPop(context.Background(), getRedisKey(activityId))
	if err != nil {
		return
	}

	redPacketId, err := strconv.ParseUint(result, 10, 64)
	if err != nil {
		return
	}

	var redPacket dto.RedPacket

	opt := &option.RedPacketOption{
		RedPacket: dto.RedPacket{
			ID: redPacketId,
		},
	}
	err = s.repo.Get(context.Background(), nil, &redPacket, opt)
	if err != nil {
		return
	}

	err = s.repo.Update(context.Background(), nil, opt, &option.RedPacketUpdateColumn{
		UserID:    userId,
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return
	}

	var user dto.User
	userOpt := &option.UserOption{
		User: dto.User{
			ID: userId,
		},
	}
	err = s.repo.Get(context.Background(), nil, &user, userOpt)
	if err != nil {
		return
	}

	err = s.repo.Update(context.Background(), nil, userOpt, &option.UserUpdateColumn{
		Balance:   user.Balance + redPacket.Amount,
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return
	}

	return redPacket.Amount, nil
}

const MinAmount = 1

func doubleAverage(count, amount uint64) uint64 {
	if count == 1 {
		return amount
	}

	//计算出最大可用金额
	max := amount - MinAmount*count

	//计算出最大可用平均值
	avg := max / count

	//二倍均值基础上再加上最小金额 防止出现金额为0
	avg2 := 2*avg + MinAmount

	//随机红包金额序列元素，把二倍均值作为随机的最大数
	rand.Seed(time.Now().UnixNano())
	x := rand.Int63n(int64(avg2)) + MinAmount

	return uint64(x)
}

func getRedisKey(activityId uint64) string {
	return "activity:" + strconv.FormatInt(int64(activityId), 10)
}
