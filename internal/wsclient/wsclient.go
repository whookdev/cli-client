package wsclient

import (
	"context"
	"fmt"
	"time"

	"github.com/gorilla/websocket"
	"github.com/whookdev/cli/internal/relay"
)

type WSClient struct {
	conn      *websocket.Conn
	done      chan struct{}
	messageCh chan []byte
}

func NewWSClient(url string) (*WSClient, error) {
	dialer := websocket.Dialer{
		HandshakeTimeout: 10 * time.Second,
	}

	conn, _, err := dialer.Dial(fmt.Sprintf("ws://%s", url), nil)
	if err != nil {
		return nil, fmt.Errorf("dialing websocket: %w", err)
	}

	client := &WSClient{
		conn:      conn,
		done:      make(chan struct{}),
		messageCh: make(chan []byte, 100),
	}

	go client.readPump()

	return client, nil
}

// Connect is a convenience function to get the websocket URL and connect in one step
func Connect(ctx context.Context, relay *relay.Relay) (*WSClient, error) {
	wsURL, err := relay.GetWSUrl()
	if err != nil {
		return nil, fmt.Errorf("getting websocket url: %w", err)
	}

	client, err := NewWSClient(wsURL)
	if err != nil {
		return nil, fmt.Errorf("creating websocket client: %s: %w", wsURL, err)
	}

	return client, nil
}

func (c *WSClient) ReadMessage() ([]byte, error) {
	select {
	case msg := <-c.messageCh:
		return msg, nil
	case <-c.done:
		return nil, fmt.Errorf("connection closed")
	}
}

func (c *WSClient) Close() error {
	close(c.done)
	return c.conn.Close()
}

func (c *WSClient) readPump() {
	defer func() {
		c.conn.Close()
		close(c.messageCh)
	}()

	for {
		select {
		case <-c.done:
			return
		default:
			_, _, err := c.conn.ReadMessage()
			if err != nil {
				if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
					return
				}
				// log other errors
			}
		}
	}
}
