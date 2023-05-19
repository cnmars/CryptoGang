package wsserver

/* import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	// "github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}



// var ps = &PubSub{}



// 此函数为单独使用没有结合Gin做路由的时候
func websocketHandler(w http.ResponseWriter, r *http.Request) {

	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true

	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	// ctx := r.Context()
	client := Client{
		Id:         AutoId(),
		Connection: conn,
	}
	// 把客户端加入列表
	ps.AddClient(client)

	fmt.Println("New Client is connected, total: ", len(ps.Clients))

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println("Something went wrong", err)

			ps.RemoveClient(client)
			log.Println("total clients and subscriptions ", len(ps.Clients), len(ps.Subscriptions))

			return
		}

		ps.HandleReceiveMessage(client, messageType, p)

	}

}

// 此函数为单独使用的时候
func New(port string, path string) *PubSub {

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

	// 	http.ServeFile(w, r, "static")

	// })
	http.HandleFunc(path, websocketHandler)

	// http.HandleFunc("/ws", websocketHandler)
	go http.ListenAndServe(":"+port, nil)
	// go http.ListenAndServe(":"+port, nil)

	fmt.Println("Server is running: http://localhost:" + port)

	return ps

}

// 测试
func main() {

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

	// 	http.ServeFile(w, r, "static")

	// })
	// http.HandleFunc("/ws", websocketHandler)
	server := New("3004", "/ws")
	go func() {
		for {
			server.PublishOld("pending", []byte("hhhhhh"), nil)
			time.Sleep(time.Second * 2)
		}
	}()
	go func() {
		for {
			server.PublishOld("aaa", []byte("aaaa"), nil)
			time.Sleep(time.Second * 2)
		}
	}()
	go func() {
		for {
			server.PublishOld("bbb", []byte("bbbbbb"), nil)
			time.Sleep(time.Second * 2)
		}
	}()

	// http.ListenAndServe(":3003", nil)
	// fmt.Println("Server is running: http://localhost:3000")

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan

} */
