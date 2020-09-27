package mock

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/websocket"
)

// WebsocketServer defines mock websocket server
// it expose connection to allow users manipulate message by themselves
type WebsocketServer struct {
	t       *testing.T
	server  *httptest.Server
	connMap map[string]*websocket.Conn
}

// GetConn gets connection by its client name in header
func (ws *WebsocketServer) GetConn(clientName string) *websocket.Conn {
	return ws.connMap[clientName]
}

// GetURL returns ws server url
func (ws *WebsocketServer) GetURL() string {
	return "ws" + strings.TrimPrefix(ws.server.URL, "http")
}

// TestHandler defines a simple test handler for ws communication
func (ws *WebsocketServer) TestHandler(w http.ResponseWriter, r *http.Request) {
	clientName := r.Header.Get("client-name")
	if clientName == "" {
		ws.t.Fatal("invalid header")
	}

	upgrader := &websocket.Upgrader{}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		ws.t.Fatal("failed to upgrade")

		return
	}

	ws.connMap[clientName] = conn
	defer func() {
		delete(ws.connMap, clientName)
		conn.Close()
	}()
	select {}
}

// Close closes test server
func (ws *WebsocketServer) Close() {
	ws.server.Close()
}

// NewWebsocketServer is the constructor of WebsocketServer
func NewWebsocketServer(t *testing.T) *WebsocketServer {
	mockWS := &WebsocketServer{
		t:       t,
		connMap: map[string]*websocket.Conn{},
	}

	mockWS.server = httptest.NewServer(http.HandlerFunc(mockWS.TestHandler))
	return mockWS
}
