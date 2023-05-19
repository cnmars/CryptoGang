package apis

import (
	"fmt"
	"golddigger/common/wsserver"
	"golddigger/global"
	"golddigger/utils/jwtutil"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:   1024,
	WriteBufferSize:  1024,
	HandshakeTimeout: time.Second * 10,
	// CheckOrigin: func(r *http.Request) bool {
	// 	return true
	// },
}

// 此函数为结合Gin做路由的时候
func WsHandler(c *gin.Context) {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	//设置Sec-Websocket-Protocol
	upgrader.Subprotocols = []string{c.Request.Header.Get("Sec-Websocket-Protocol")}
	//升级get请求为webSocket协议
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}
	claims := jwtutil.GetUserInfo(c)
	client := wsserver.Client{
		Id: claims.Address,
		// Id:         wsserver.AutoId(),
		Connection: conn,
		IsConnect:  true,
		MessageCh:  make(chan wsserver.WsMessage),
	}
	// 把客户端加入列表
	global.GD_WS.AddClient(client)

	// defer conn.Close()

	fmt.Println("New Client is connected, total: ", len(global.GD_WS.Clients))
	//客户端命令处理函数
	go handleReciveMsg(client)
	//消息推送函数
	go global.GD_WS.PublishNew(client)

}

// 此函数为不通过gin路由 单独使用时调用
func WebsocketHandler(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := wsserver.Client{
		Id:         wsserver.AutoId(),
		Connection: conn,
		IsConnect:  true,
		MessageCh:  make(chan wsserver.WsMessage),
	}
	// 把客户端加入列表
	global.GD_WS.AddClient(client)

	// defer conn.Close()

	fmt.Println("New Client is connected, total: ", len(global.GD_WS.Clients))
	//客户端命令处理函数
	go handleReciveMsg(client)
	//消息推送函数
	go global.GD_WS.PublishNew(client)
}

func handleReciveMsg(client wsserver.Client) {
	for {
		messageType, p, err := client.Connection.ReadMessage()
		if err != nil {
			fmt.Println("Something went wrong", err)

			global.GD_WS.RemoveClient(client)
			fmt.Println("total clients and subscriptions ", len(global.GD_WS.Clients), len(global.GD_WS.Subscriptions))
			client.IsConnect = false
			return
		}

		global.GD_WS.HandleReceiveMessage(client, messageType, p)

	}
}
