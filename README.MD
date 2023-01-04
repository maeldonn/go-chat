# TCP Chat

_(A very simple chat written in Go)_

## Running the server

To run the server you will need GNU make and Go SDK installed.

```sh
$ make run
```

## Interact with the server

To send messages, you will need a TCP client. I recommand using telnet.

```sh
$ telnet localhost 8888
```

## Useful commands

Change your nickname
```
/nick Bob
```

List all rooms
```
/rooms
```

Join a chatting room
```
/join #general
```

Send a message to the room
```
/msg Hello world!
```

Quit the chat
```
/quit
```

