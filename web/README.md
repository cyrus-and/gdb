# web

Package `web` provides a simple web interface for the `gdb` module.

Every new handler spawns a new GDB instance which is not meant to serve more than one client.

Several URLs are available:

- a WebSocket will be bound to `NotificationsUrl`, every notification sent by GDB will be delivered through this WebSocket as a JSON object;

- a WebSocket will be bound to `TerminalUrl`: data sent through it will be delivered to the target program's input whereas data read from it will reflect the the target program's output;

- POST requests are expected to `CommandsUrl` to send commands to GDB. Commands are represented by JSON arrays `[command, arguments...]`. Command responses are sent back as JSON objects in the response body;

- POST requests are expected to `InterruptUrl` to send interrupts to GDB, no parameters are allowed.

HTTP methods return 200 on success and 500 on error.

## Example

A dummy example can be found in the [example](example/) directory.

## Installation

This package requires the `websocket` library from the additional networking packages:

```
go get golang.org/x/net/websocket
```
