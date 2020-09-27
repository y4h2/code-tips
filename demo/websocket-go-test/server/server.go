package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

// App defines app server
type App struct {
	server   *http.Server
	listener net.Listener
}

// Start starts server
func (app *App) Start() error {
	listener, err := net.Listen("tcp", app.server.Addr)
	if err != nil {
		return err
	}
	app.listener = listener

	go func() {
		if err := app.server.Serve(app.listener); err != nil {
			log.Fatalf("failed to launch app :%v", err)
		}
	}()

	return nil
}

// Shutdown shuts down the server
func (app *App) Shutdown() {
	app.server.Shutdown(context.Background())
}

// GetPort returns http server's binding port
func (app *App) GetPort() int {
	return app.listener.Addr().(*net.TCPAddr).Port
}

// NewApp is the constructor of App
func NewApp(port int) *App {
	router := mux.NewRouter()
	router.HandleFunc("/ping", pingHandler)
	router.HandleFunc("/ws", wsHandler)

	app := &App{
		server: &http.Server{
			Addr:    fmt.Sprintf(":%d", port),
			Handler: router,
		},
	}
	return app
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  4096,
		WriteBufferSize: 4096,
		CheckOrigin: func(*http.Request) bool {
			return true
		},
	}

	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("failed to upgrade HTTP connection")
		return
	}

	defer c.Close()

	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}
