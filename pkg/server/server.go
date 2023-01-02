package server

import (
	"errors"
	"fmt"
	"log"
	"net"
	"strings"
)

type server struct {
	rooms    map[string]*room
	commands chan command
}

func NewServer() *server {
	return &server{
		rooms:    make(map[string]*room),
		commands: make(chan command),
	}
}

func (server *server) Run() {
	for cmd := range server.commands {
		switch cmd.id {
		case CMD_NICK:
			server.nick(cmd.client, cmd.args)
		case CMD_JOIN:
			server.join(cmd.client, cmd.args)
		case CMD_ROOMS:
			server.listRooms(cmd.client, cmd.args)
		case CMD_MSG:
			server.msg(cmd.client, cmd.args)
		case CMD_QUIT:
			server.quit(cmd.client, cmd.args)
		}
	}
}

func (server *server) NewClient(connection net.Conn) *client {
	log.Printf("New client gas connected: %s", connection.RemoteAddr().String())
	return &client{
		connection: connection,
		nickname:   "Anonymous",
		commands:   server.commands,
	}
}

func (server *server) nick(client *client, args []string) {
	name := args[1]
	client.nickname = name
	client.msg(fmt.Sprintf("All right, I will call you %s", name))
}

func (server *server) join(c *client, args []string) {
	roomName := args[1]

	r, ok := server.rooms[roomName]
	if !ok {
		r = &room{
			name:    roomName,
			members: make(map[net.Addr]*client),
		}
		server.rooms[roomName] = r
	}

	r.members[c.connection.RemoteAddr()] = c

	server.quitCurrentRoom(c)

	c.room = r

	r.broadcast(c, fmt.Sprintf("%s has joined the room", c.nickname))
	c.msg(fmt.Sprintf("Welcome to %s", r.name))
}

func (server *server) listRooms(client *client, args []string) {
	var rooms []string
	for name := range server.rooms {
		rooms = append(rooms, name)
	}

	client.msg(fmt.Sprintf("Available rooms are: %s", strings.Join(rooms, ", ")))
}

func (server *server) msg(client *client, args []string) {
	if client.room == nil {
		client.err(errors.New("You must join a room first"))
		return
	}

	client.room.broadcast(client, client.nickname+": "+strings.Join(args[1:], " "))
}

func (server *server) quit(client *client, args []string) {
	log.Printf("Client has disconnected: %s", client.connection.RemoteAddr())

	server.quitCurrentRoom(client)

	client.msg("Sad to see you go :(")
	client.connection.Close()
}

func (server *server) quitCurrentRoom(client *client) {
	if client.room != nil {
		delete(client.room.members, client.connection.RemoteAddr())
		client.room.broadcast(client, fmt.Sprintf("%s has left the room", client.nickname))
	}
}
