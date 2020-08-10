package main

import (
	"database/sql"
	"math/rand"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func setup(t *testing.T) *sql.DB {
	rand.Seed(time.Now().UnixNano())
	t.Helper()
	db, err := ConnectDB(Config{
		Host:     "127.0.0.1",
		Port:     3306,
		User:     "test_user",
		Password: "testtest",
		Database: "test",
	})
	if err != nil {
		t.Fatalf("failed to connect DB: %v", err)
	}
	return db
}

func cleanup(t *testing.T, db *sql.DB, orderID int64) {
	t.Helper()
	db.Exec("DELETE FROM orders WHERE order_id = ?", orderID)
	db.Close()
}

func TestCreateOrder(t *testing.T) {
	assert := assert.New(t)

	t.Log("Given a DB connection and an order ID")
	db := setup(t)
	orderID := rand.Int63()
	userID := rand.Int63()
	defer cleanup(t, db, orderID)

	t.Log("When I call CreateOrder")
	ok, err := CreateOrder(db, orderID, userID)
	assert.NoError(err)
	t.Log("Then the order should be created successfully")
	assert.True(ok)

	t.Log("When I call CreateOrder again")
	ok, err = CreateOrder(db, orderID, userID)
	assert.NoError(err)
	t.Log("Then the order creation should fail")
	assert.False(ok)
}

func goroutineWithRandomDelay(f func()) {
	go func() {
		time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
		f()
	}()
}

func TestUpdateOrder(t *testing.T) {
	assert := assert.New(t)
	t.Log("Given a DB connection and an order")
	db := setup(t)
	orderID := rand.Int63n(100)
	ok, err := CreateOrder(db, orderID, rand.Int63n(100))
	assert.NoError(err)
	assert.True(ok)
	defer cleanup(t, db, orderID)

	t.Log("When I run two UpdateOrder in goroutine with random delays")
	var wg sync.WaitGroup
	wg.Add(2)
	var ok1, ok2 bool
	goroutineWithRandomDelay(func() {
		ok1, err = UpdateOrder(db, orderID, "status1", 0)
		assert.NoError(err)
		if ok1 {
			t.Log("Update to status1")
		}
		wg.Done()
	})
	goroutineWithRandomDelay(func() {
		ok2, err = UpdateOrder(db, orderID, "status2", 0)
		assert.NoError(err)
		if ok2 {
			t.Log("Update to status2")
		}
		wg.Done()
	})

	t.Log("Then there should be one and only one update success")
	wg.Wait()
	assert.False(ok1 && ok2)
}
