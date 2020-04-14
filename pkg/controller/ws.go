package controller

import (
	"github.com/google/uuid"
	"golang.org/x/net/websocket"
	"log"
)

var (
	manager clientManager
)

type clientManager struct {
	clients    map[*client]bool
	broadcast  chan interface{}
	register   chan *client
	unregister chan *client
}

func (manager *clientManager) start() {
	for {
		select {
		case conn := <-manager.register:
			manager.clients[conn] = true
		case conn := <-manager.unregister:
			if _, ok := manager.clients[conn]; ok {
				close(conn.send)
				delete(manager.clients, conn)
			}
		case message := <-manager.broadcast:
			log.Printf("Broadasting to %d clients: %+v", len(manager.clients), message)
			for conn := range manager.clients {
				log.Printf("Broadasting message to client %s", conn.id)
				select {
				case conn.send <- message:
				default:
					close(conn.send)
					delete(manager.clients, conn)
				}
			}
		}
	}
}

func (manager *clientManager) send(message interface{}) {
	for conn := range manager.clients {
		conn.send <- message
	}
}

type client struct {
	id     string
	socket *websocket.Conn
	send   chan interface{}
}

func (c *client) write() {
	defer func() {
		c.socket.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:
			log.Printf("On send %v - %v", message, ok)
			if !ok {
				log.Println("Unable to sent")
				return
			}
			websocket.JSON.Send(c.socket, message)
		}
	}
}

func init() {
	manager = clientManager{
		broadcast:  make(chan interface{}, 100),
		register:   make(chan *client),
		unregister: make(chan *client),
		clients:    make(map[*client]bool),
	}
	go manager.start()
}

func (c *Controller) WSHandler(ws *websocket.Conn) {
	log.Println("WS connection...")

	client := &client{
		id:     uuid.New().String(),
		socket: ws,
		send:   make(chan interface{}),
	}
	manager.register <- client
	go reader(ws)
	client.write()
}

type message struct {
	Yaml string `json:"yaml"`
	Mode string `json:"mode"`
	Target string `json:"target"`
}

func reader(conn *websocket.Conn) {
	//b := make([]byte, 0, 1024*10) // TODO: set higher?
	//for {
	//	// read in a message
	//	n, err := conn.Read(b)
	//	if err != nil {
	//		log.Println(n, err)
	//		return
	//	}
	//	if n > 0 {
	//		fmt.Println("read ", n, string(b))
	//	}
	//}

	//

	for {
		// allocate our container struct
		var m message
		// receive a message using the codec
		if err := websocket.JSON.Receive(conn, &m); err != nil {
			log.Println(err)
			continue
		}
		log.Println("Received message:", m.Yaml)
		//// send a response
		//m2 := message{"Thanks for the message!"}
		//if err := websocket.JSON.Send(ws, m2); err != nil {
		//	log.Println(err)
		//	break
		//}
	}
}
