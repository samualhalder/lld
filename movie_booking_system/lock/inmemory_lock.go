package lock

import (
	"fmt"
	"sync"
	"time"
)

type InMemeoryLock struct {
	keys         map[string]string
	exTime       map[string]time.Time
	expriesAfter time.Duration
	mutex        sync.RWMutex
}

func (i *InMemeoryLock) SetLock(userId string, slotId string, seatId string) error {
	defer i.mutex.Unlock()
	key := i.createKey(seatId, seatId)
	i.mutex.Lock()
	if i.IsLocked(seatId, slotId) {
		return fmt.Errorf("sear already locked")
	}
	expiresAt := time.Now().Add(i.expriesAfter)
	i.keys[key] = userId
	i.exTime[key] = expiresAt
	return nil
}
func (i *InMemeoryLock) RelaseLock(slotId string, seatId string) error {
	defer i.mutex.Unlock()
	key := i.createKey(seatId, seatId)
	i.mutex.Lock()
	delete(i.keys, key)
	delete(i.exTime, key)
	return nil
}
func (i *InMemeoryLock) IsExpired(slotId string, seatId string) bool {
	defer i.mutex.Unlock()
	key := i.createKey(seatId, seatId)
	i.mutex.Lock()
	exTime := i.exTime[key]
	if exTime.Before(time.Now()) {
		return false
	}
	return true
}
func (i *InMemeoryLock) IsLocked(slotId string, seatId string) bool {
	defer i.mutex.Unlock()
	key := i.createKey(seatId, seatId)
	i.mutex.Lock()
	_, ok := i.keys[key]
	if !ok {
		return false
	}
	if i.IsExpired(slotId, seatId) {
		return false
	}
	return true
}

func (i *InMemeoryLock) createKey(slotId string, seatId string) string {
	return fmt.Sprint(slotId + "#" + seatId)
}

func (i *InMemeoryLock) LockedBy(slotId string, seatid string) string {
	key := i.createKey(slotId, seatid)
	return i.keys[key]
}
