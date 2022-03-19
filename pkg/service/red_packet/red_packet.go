package red_packet

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"math/rand"
	"red-packet/pkg/model/dto"
	"red-packet/pkg/model/option"
	"strconv"
	"time"
)

func (s *service) Send(ctx context.Context, activity *dto.Activity) (bool, error) {
	if err := s.repo.Transaction(ctx, func(tx *gorm.DB) (err error) {
		err = s.repo.Create(ctx, nil, activity)
		if err != nil {
			return
		}

		redPackets := genRedPacket(activity.ID, activity.Count, activity.Amount)

		err = s.repo.Create(ctx, nil, &redPackets)
		if err != nil {
			return
		}

		return nil
	}); err != nil {
		return false, err
	}

	return true, nil
}

func (s *service) Grab(ctx context.Context, redPacketOpt *option.RedPacketOption, userOpt *option.UserOption) (uint64, error) {
	lock, err := s.redisLock.Obtain(ctx, getRedisKey(redPacketOpt.RedPacket.ID), 5*time.Second, nil)
	if err != nil {
		return 0, err
	}

	var amount uint64

	if err = s.repo.Transaction(ctx, func(tx *gorm.DB) (err error) {
		var redPackets []dto.RedPacket
		_, err = s.repo.Get(ctx, nil, &redPackets, redPacketOpt)
		if err != nil {
			return
		}

		if len(redPackets) < 1 {
			return errors.New("red packet not enough")
		}

		idx := rand.Int63n(int64(len(redPackets)))
		redPacket := redPackets[idx]
		amount = redPacket.Amount

		opt2 := &option.RedPacketOption{
			RedPacket: dto.RedPacket{
				ID: redPacket.ID,
			},
		}

		err = s.repo.Update(ctx, nil, opt2, &option.RedPacketUpdateColumn{
			UserID:    userOpt.User.ID,
			Status:    2,
			UpdatedAt: time.Now(),
		})
		if err != nil {
			return
		}

		var user dto.User
		err = s.repo.GetOne(ctx, nil, &user, userOpt)
		if err != nil {
			return
		}

		err = s.repo.Update(ctx, nil, userOpt, &option.UserUpdateColumn{
			Balance:   user.Balance + redPacket.Amount,
			UpdatedAt: time.Now(),
		})
		if err != nil {
			return
		}

		return nil
	}); err != nil {
		return 0, nil
	}

	err = lock.Release(ctx)
	if err != nil {
		return 0, err
	}

	return amount, nil
}

const MinAmount = 1

func genRedPacket(activityId uint64, count uint64, amount uint64) (result []dto.RedPacket) {
	remain := amount
	sum := uint64(0)
	for i := uint64(0); i < count; i++ {
		x := doubleAverage(count-i, remain)
		remain -= x
		sum += x

		tmp := dto.RedPacket{
			ActivityID: activityId,
			UserID:     0,
			Amount:     x,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		}

		result = append(result, tmp)
	}

	return
}

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
