package iface

type IRedPacketService interface {
	Send(userId uint64, count uint64, amount uint64) (bool, error)
	Grab(userId uint64, activityId uint64) (uint64, error)
}
