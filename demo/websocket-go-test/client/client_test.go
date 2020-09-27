package client

import (
	"net/http"
	"sync"
	"testing"
	"time"
	"websocket-go-test/mock"

	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
)

func TestClient(t *testing.T) {
	assert := assert.New(t)
	t.Log("Start a mock websocket server")
	server := mock.NewWebsocketServer(t)
	defer server.Close()

	clientName := "client1"
	header := http.Header{"client-name": []string{clientName}}

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		RunWSClient(server.GetURL(), header)
		wg.Done()
	}()

	var conn *websocket.Conn
	for {
		conn = server.GetConn(clientName)
		if conn != nil {
			break
		}
		time.Sleep(1 * time.Second)
	}

	_, bytes, err := conn.ReadMessage()
	assert.NoError(err)
	assert.Equal("message 1", string(bytes))

	_, bytes, err = conn.ReadMessage()
	assert.NoError(err)
	assert.Equal("message 2", string(bytes))

	_, bytes, err = conn.ReadMessage()
	assert.NoError(err)
	assert.Equal("message 3", string(bytes))

	wg.Wait()

}
