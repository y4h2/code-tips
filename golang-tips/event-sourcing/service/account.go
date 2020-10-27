package service

import (
	"errors"
	"time"

	"event-sourcing/enum"
	"event-sourcing/eventstore"
	"event-sourcing/model"
)

type Account struct {
	Balance int
}

type AccountEntity struct {
	eventStore *eventstore.EventStore
}

func (a *AccountEntity) Open(ID string, balance int) error {
	openEvent := model.Event{
		Type:     enum.EventType.Open,
		Entity:   enum.EventEntity.Account,
		EntityID: ID,
		Content: model.OpenEventContent{
			Balance: balance,
		},
		Time: time.Now().UTC(),
	}

	// store event
	return a.eventStore.StoreEvent([]model.Event{openEvent})
}

func (a *AccountEntity) Transfer(fromID string, toID string, amount int) error {
	balance, err := a.GetBalance(fromID)
	if err != nil {
		return err
	}
	if balance < amount {
		return errors.New("not enough balance")
	}

	curTime := time.Now().UTC()
	transferEvents := []model.Event{
		{
			Type:     enum.EventType.Transfer,
			Entity:   enum.EventEntity.Account,
			EntityID: fromID,
			Content: model.TransferEventContent{
				Amount: amount,
				FromID: fromID,
				ToID:   toID,
			},
			Time: curTime,
		},
		{
			Type:     enum.EventType.Transfer,
			Entity:   enum.EventEntity.Account,
			EntityID: toID,
			Content: model.TransferEventContent{
				Amount: amount,
				FromID: fromID,
				ToID:   toID,
			},
			Time: curTime,
		},
	}
	return a.eventStore.StoreEvent(transferEvents)
}

func (a *AccountEntity) Store(ID string, amount int) error {
	storeEvent := model.Event{
		Type:     enum.EventType.Open,
		Entity:   enum.EventEntity.Account,
		EntityID: ID,
		Content: model.StoreEventContent{
			Amount: amount,
		},
		Time: time.Now().UTC(),
	}

	// store event
	return a.eventStore.StoreEvent([]model.Event{storeEvent})
}

func (a *AccountEntity) Withdraw(ID string, amount int) error {
	balance, err := a.GetBalance(ID)
	if err != nil {
		return err
	}
	if balance < amount {
		return errors.New("not enough balance")
	}

	withdrawEvent := model.Event{
		Type:     enum.EventType.Open,
		Entity:   enum.EventEntity.Account,
		EntityID: ID,
		Content: model.WithdrawEventContent{
			Amount: amount,
		},
		Time: time.Now().UTC(),
	}

	return a.eventStore.StoreEvent([]model.Event{withdrawEvent})
}

func (a *AccountEntity) HandleEvent(accountID string, evt model.Event) (int, error) {
	switch evt.Type {
	case enum.EventType.Open:
		content, ok := (evt.Content).(model.OpenEventContent)
		if !ok {
			return 0, errors.New("cannot assert to open event")
		}
		return content.Balance, nil
	case enum.EventType.Store:
		content, ok := (evt.Content).(model.StoreEventContent)
		if !ok {
			return 0, errors.New("cannot assert to store event")
		}

		return content.Amount, nil

	case enum.EventType.Withdraw:
		content, ok := (evt.Content).(model.WithdrawEventContent)
		if !ok {
			return 0, errors.New("cannot assert to withdraw event")
		}

		return -1 * content.Amount, nil

	case enum.EventType.Transfer:
		content, ok := (evt.Content).(model.TransferEventContent)
		if !ok {
			return 0, errors.New("cannot assert to transfer event")
		}
		if accountID == content.FromID {
			return -1 * content.Amount, nil
		} else if accountID == content.ToID {
			return content.Amount, nil
		}

		return 0, errors.New("invalid event")

	default:
		return 0, errors.New("invalid event type")
	}
}

func (a *AccountEntity) GetBalance(ID string) (int, error) {
	balance := 0
	events := a.eventStore.ListEvents(enum.EventEntity.Account, ID)
	for _, evt := range events {
		incr, err := a.HandleEvent(ID, evt)
		if err != nil {
			return 0, err
		}
		balance += incr
	}

	return balance, nil
}

type SnapshotStore struct {
}
