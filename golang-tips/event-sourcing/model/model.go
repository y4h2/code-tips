package model

import (
	"time"
)

type OpenEventContent struct {
	Balance int
}

type TransferEventContent struct {
	Amount int
	FromID string
	ToID   string
}

type StoreEventContent struct {
	Amount int
}

type WithdrawEventContent struct {
	Amount int
}

type Event struct {
	ID       int
	Type     string
	Entity   string
	EntityID string
	Content  interface{}
	Time     time.Time
}

type Snapshot struct {
	Type            string
	Entity          string
	EntityID        string
	EntityVersion   int
	SnapshotType    string
	SnapshotContent string
}
