package wsserver

import (
	// "fmt"
	// "time"

	// "context"
	// "time"

	uuid "github.com/satori/go.uuid"
)

func AutoId() string {
	//go mod edit -replace=github.com/satori/go.uuid@v1.2.0=github.com/satori/go.uuid@master
	//or go get github.com/satori/go.uuid@master
	return uuid.Must(uuid.NewV4()).String()
}

// 数据传输函数，把三方数据变为通用数据写入客户端数据通道，没有用这个函数
func DataTransfer(from <-chan SerilizeData, client *Client) {
	for client.IsConnect {
		select {
		case f := <-from:
			client.MessageCh <- f.ToCommonData()
		}
	}
}

// 分发三方数据到客户端
func (pb *PubSub) Distribute() {
	for !pb.IsStop {
		select {
		case v := <-pb.ThirdPartData:
			ty := v.GetType()
			for _, sub := range pb.Subscriptions {
				if ty == sub.Topic {
					sub.Client.MessageCh <- v
				}
			}
		}
	}
}

// func (pb *PubSub) Start() {
// 	pb.ThirdPartData = make(chan WsMessage, 100)
// 	go pb.FetchPending()
// 	go pb.FetchGas()
// 	go pb.Distribute()
// 	pb.IsStop = false

// }


