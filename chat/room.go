package chat

import (
	"net"
)

type room struct {
	name    string
	members map[net.Addr]*client
}

func (r *room) broadcast(sender *client, msg string) {
	for addr, members := range r.members {
		if addr != sender.conn.RemoteAddr() {
			members.msg(msg)
		}
	}
}
