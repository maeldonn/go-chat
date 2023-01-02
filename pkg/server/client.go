package server

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

type client struct {
	connection net.Conn
	nickname   string
	room       *room
	commands   chan<- command
}

func (client *client) ReadInput() {
	for {
		msg, err := bufio.NewReader(client.connection).ReadString('\n')
		if err != nil {
			return
		}

		msg = strings.Trim(msg, "\r\n")

		args := strings.Split(msg, " ")
		cmd := strings.TrimSpace(args[0])

		// TODO: Gerer le cas ou pas d'argument

		switch cmd {
		case "/nick":
			client.commands <- command{
				id:     CMD_NICK,
				client: client,
				args:   args,
			}
		case "/join":
			client.commands <- command{
				id:     CMD_JOIN,
				client: client,
				args:   args,
			}
		case "/rooms":
			client.commands <- command{
				id:     CMD_ROOMS,
				client: client,
				args:   args,
			}
		case "/msg":
			client.commands <- command{
				id:     CMD_MSG,
				client: client,
				args:   args,
			}
		case "/quit":
			client.commands <- command{
				id:     CMD_QUIT,
				client: client,
				args:   args,
			}
		default:
			client.err(fmt.Errorf("Unknown command: %s", cmd))
		}
	}
}

func (client *client) msg(msg string) {
	client.connection.Write([]byte("> " + msg + "\n"))
}

func (client *client) err(err error) {
	client.connection.Write([]byte("ERR: " + err.Error() + "\n"))
}
