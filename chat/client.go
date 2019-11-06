package main

import (
	"time"

	"github.com/gorilla/websocket"
)

type client struct {
	// socket은 이 클라이언트의 웹 소켓이다.
	socket *websocket.Conn
	// send는 메시지가 전송되는 채널이다.
	send chan *message
	// room은 클라이언트가 채팅하는 방이다.
	room *room
	// userData 는 사용자에 대한 정보를 보유한다.
	userData map[string]interface{}
}

func (c *client) read() {
	defer c.socket.Close()

	for {
		var msg *message
		if err := c.socket.ReadJSON(&msg); err != nil {
			return
		}

		msg.When = time.Now()
		msg.Name = c.userData["name"].(string)
		if avatarURL, ok := c.userData["avatar_url"]; ok {
			msg.AvatarURL = avatarURL.(string)
		}
		c.room.forward <- msg
	}
}

func (c *client) write() {
	defer c.socket.Close()

	for msg := range c.send {
		err := c.socket.WriteJSON(msg)
		if err != nil {
			break
		}
	}
}
