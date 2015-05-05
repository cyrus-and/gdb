package web

import (
	"encoding/json"
	"github.com/cyrus-and/gdb"
	"golang.org/x/net/websocket"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

const (
	NotificationsUrl = "/notifications"
	TerminalUrl      = "/terminal"
	CommandsUrl      = "/send"
	InterruptUrl     = "/interrupt"
)

func die(err error, w http.ResponseWriter) {
	log.Print("### ", err)
	w.WriteHeader(http.StatusInternalServerError)
}

// NewHandler returns a new http.ServeMux which automatically spawn a new
// instance of GDB and offers HTTP mount points to interact with. It is assumed
// that only one client at a time access this multiplexer.
func NewHandler() (*http.ServeMux, error) {
	handler := http.NewServeMux()
	notificationsChan := make(chan []byte)

	// start a new GDB instance
	gdb, err := gdb.New(func(notification map[string]interface{}) {
		// notifications are converted to JSON and sent through the channel
		notificationText, err := json.Marshal(notification)
		if err != nil {
			log.Fatal(err)
		}
		notificationsChan <- notificationText
		log.Print(">>> ", string(notificationText))
	})
	if err != nil {
		return nil, err
	}

	// notifications WebSocket
	handler.Handle(NotificationsUrl, websocket.Handler(func(ws *websocket.Conn) {
		// deliver the notifications through the WebSocket
		log.Print("### notification WS opened")
		for notification := range notificationsChan {
			ws.Write(notification)
		}
		log.Print("### notification WS closed")
	}))

	// terminal WebSocket
	handler.Handle(TerminalUrl, websocket.Handler(func(ws *websocket.Conn) {
		// copy GDB to WS and WS to GDB
		log.Print("### terminal WS opened")
		var wg sync.WaitGroup
		wg.Add(2)
		go func() {
			io.Copy(gdb, ws)
			wg.Done()
			log.Print("### write side closed")
		}()
		go func() {
			io.Copy(ws, gdb)
			wg.Done()
			log.Print("### read side closed")
		}()
		wg.Wait()
		log.Print("### terminal WS closed")
	}))

	// send command action
	handler.HandleFunc(CommandsUrl, func(w http.ResponseWriter, req *http.Request) {
		if req.Method != "POST" {
			log.Print("### invalid method")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		data, err := ioutil.ReadAll(req.Body)
		if err != nil {
			die(err, w)
			return
		}
		log.Print("<<< ", string(data))
		command := []string{}
		err = json.Unmarshal(data, &command)
		if err != nil {
			die(err, w)
			return
		}
		result, err := gdb.Send(command[0], command[1:]...)
		if err != nil {
			die(err, w)
			return
		}
		reply, err := json.Marshal(result)
		if err != nil {
			die(err, w)
			return
		}
		io.WriteString(w, string(reply))
		log.Print(">>> ", string(reply))
	})

	// send interrupt action
	handler.HandleFunc(InterruptUrl, func(w http.ResponseWriter, req *http.Request) {
		if req.Method != "POST" {
			log.Print("### invalid method")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if err := gdb.Interrupt(); err != nil {
			die(err, w)
			return
		}
		log.Print("<<< interrupt")
	})

	return handler, nil
}
