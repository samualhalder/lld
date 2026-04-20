package lock

type Lock interface {
	SetLock(userId string, slotId string, seatId string) error
	RelaseLock(slotId string, seatId string) error
	IsExpired(slotId string, seatId string) bool
	IsLocked(slotId string, seatId string) bool
}
