package store

import "sync"

type DataPoint struct {
	Timestamp uint64
	Value     uint64
}

type DataStore struct {
	l sync.RWMutex
	u map[string][]DataPoint
}

func NewDataStore() *DataStore {
	return &DataStore{u: map[string][]DataPoint{}}
}

func (u *DataStore) Get(symbol string) ([]DataPoint, bool) {
	u.l.RLock()
	defer u.l.RUnlock()
	dataPoint, ok := u.u[symbol]
	return dataPoint, ok
}

func (u *DataStore) Add(symbol string, dataPoint DataPoint) {
	u.l.Lock()
	defer u.l.Unlock()
	u.u[symbol] = append(u.u[symbol], dataPoint)
}

func (u *DataStore) Delete(symbol string) {
	u.l.Lock()
	defer u.l.Unlock()
	delete(u.u, symbol)
}
