// Package web provides a simple web interface for the gdb module.
//
// Every new handler spawns a new GDB instance which is not meant to serve more
// than one client.
//
// A WebSocket will be bound to NotificationsUrl, every notification sent by GDB
// will be delivered through this WebSocket as a JSON object.
//
// A WebSocket will be bound to TerminalUrl: data sent through it will be
// delivered to the target program's input whereas data read from it will
// reflect the the target program's output.
//
// POST requests are expected to CommandsUrl to send commands to GDB. Commands
// are represented by JSON arrays [command, arguments...]. Command responses are
// sent back as JSON objects in the response body.
//
// POST requests are expected to InterruptUrl to send interrupts to GDB, no
// parameters are allowed.
//
// HTTP methods return 200 on success and 500 on error.
package web
