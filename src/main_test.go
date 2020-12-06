package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestConnectionsAreEstablishedAndMessageIsRead(t *testing.T) {
	t.Run("returns ", func(t *testing.T) {
		s := httptest.NewServer(http.HandlerFunc(handleConnections))
		defer s.Close()
		u := "ws" + strings.TrimPrefix(s.URL, "http")

		ws, response, err := websocket.DefaultDialer.Dial(u, nil)
		if err != nil {
			t.Fatalf("%v", err)
		}
		assert.Equal(t, response.StatusCode, http.StatusSwitchingProtocols, fmt.Sprintf("resp.StatusCode = %q, want %q", response.StatusCode, http.StatusSwitchingProtocols))

		b := []byte("hello server")
		ws.WriteMessage(websocket.TextMessage,b)
		msgType, message, err:= ws.ReadMessage()
		assert.NoError(t, err)
		assert.Equal(t, websocket.TextMessage, msgType)
		assert.Equal(t, b, message)
		defer ws.Close()
		s.Close()
	})
}
func TestConnectionsAreEstablishedAndBroadCastedInChannel(t *testing.T) {
	t.Run("returns ", func(t *testing.T) {
		s := httptest.NewServer(http.HandlerFunc(handleConnections))
		defer s.Close()
		u := "ws" + strings.TrimPrefix(s.URL, "http")

		ws, response, err := websocket.DefaultDialer.Dial(u, nil)
		if err != nil {
			t.Fatalf("%v", err)
		}
		assert.Equal(t, response.StatusCode, http.StatusSwitchingProtocols, fmt.Sprintf("resp.StatusCode = %q, want %q", response.StatusCode, http.StatusSwitchingProtocols))

		b := []byte("hello server")
		ws.WriteMessage(websocket.TextMessage,b)
		msgType, message, err:= ws.ReadMessage()
		assert.NoError(t, err)
		assert.Equal(t, websocket.TextMessage, msgType)
		assert.Equal(t, "hello server", message)
		handleMessages()
		assert.Equal(t, b, string(<-broadcast))
		defer ws.Close()
		s.Close()
	})
}