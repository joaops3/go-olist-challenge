package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/joaops3/go-olist-challenge/internal/api/services"
)

type MessageType string

const BROADCAST MessageType = "BROADCAST"

type EventMessages struct {
	Type MessageType  `json:"type"`
	Data json.RawMessage `json:"data"`
}

type WSController struct {
	UserService services.UserServiceInterface
	Websocket websocket.Upgrader
	subscriptions map[string]map[*websocket.Conn]context.CancelFunc
	mu *sync.Mutex
}

func InitWSController(userService services.UserServiceInterface) *WSController {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	controller := &WSController{UserService: userService,
		Websocket: upgrader,
		mu:  &sync.Mutex{},
		subscriptions: make(map[string]map[*websocket.Conn]context.CancelFunc),
	}
	return controller
}

func(c *WSController) HandleWebsocket(ctx *gin.Context){
	conn, err := c.Websocket.Upgrade(ctx.Writer, ctx.Request, nil)

	id := ctx.Param("id")

	if err != nil {
		sendError(ctx, 400, "Failed to upgrade connection")
		return
	}

	defer conn.Close()

	contx, cancel := context.WithCancel(ctx)

	c.mu.Lock()

	if _, ok := c.subscriptions[id]; !ok {
		c.subscriptions[id] = make(map[*websocket.Conn]context.CancelFunc)
	}
	c.subscriptions[id][conn] = cancel
	c.mu.Unlock()
	

	go func(){
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					fmt.Printf("WebSocket closed unexpectedly: %v\n", err)
				}
				break
			}

			event := &EventMessages{}

			err = json.Unmarshal(message, event)
			if err != nil {
				fmt.Printf("Error parsing message: %v\n", err)
				continue
			}
			
			switch event.Type {
				case BROADCAST:
					c.broadCastToChannel(id, event.Data, conn)
				default:
					c.broadCastToChannel(id, message, conn)
			}
		
		}
		cancel()
	}()

	<- contx.Done()
	c.mu.Lock()
	delete(c.subscriptions, id)
	if len(c.subscriptions[id]) == 0 {
		delete(c.subscriptions, id)
	}
	defer c.mu.Unlock()
	return
}

func(c *WSController) broadCastToChannel(channelID string, message []byte, senderConn *websocket.Conn) {
	
	c.mu.Lock()
	defer c.mu.Unlock()
	
	for subscriberConn := range c.subscriptions[channelID] {
		if subscriberConn != senderConn {
			if err := subscriberConn.WriteMessage(websocket.TextMessage, message); err != nil {
				fmt.Printf("Error broadcasting message: %v\n", err)
			}
		}
	}

}

