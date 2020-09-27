package server

import (
	"fmt"
	"testing"

	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
)

func TestServer(t *testing.T) {
	assert := assert.New(t)

	t.Log("Start websocket server")
	app := NewApp(0)
	err := app.Start()
	assert.NoError(err)

	defer app.Shutdown()

	t.Log("Connect ws server")
	wssURL := fmt.Sprintf("ws://127.0.0.1:%d/ws", app.GetPort())
	conn, _, err := websocket.DefaultDialer.Dial(wssURL, nil)
	assert.NoError(err)

	t.Log("When I send a message to the server")
	message1 := "message 1"
	conn.WriteMessage(websocket.TextMessage, []byte(message1))
	_, returnedBytes, err := conn.ReadMessage()
	assert.NoError(err)
	assert.Equal(message1, string(returnedBytes))

	// ref: https://github.com/gorilla/websocket/blob/master/examples/echo/client.go#L70
	err = conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	assert.NoError(err)
	conn.Close()
}
