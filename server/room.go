package server

import (
	"net"
)

type room struct {
	name    string
	members map[net.Addr]*client
}

func (room *room) broadcast(sender *client, msg string) {
	for addr, members := range room.members {
		if addr != sender.connection.RemoteAddr() {
			members.msg(msg)
		}
	}
}
