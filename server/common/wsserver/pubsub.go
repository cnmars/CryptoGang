package wsserver

import (
	"encoding/json"
	"fmt"
"sync"
	"github.com/gorilla/websocket"
)

const (
	PUBLISH     = "publish"
	SUBSCRIBE   = "subscribe"
	UNSUBSCRIBE = "unsubscribe"
)

type PubSub struct {
	Clients       []Client
	lock_Clients sync.Mutex
	Subscriptions []Subscription
	ThirdPartData chan WsMessage
	IsStop        bool
	//key为tpoic，value为客户端id数组
	// TopicSummary map[string]*HashSet
}

type Client struct {
	Id         string
	Connection *websocket.Conn
	MessageCh  chan WsMessage
	Topic      HashSet
	IsConnect  bool
	// Mux        sync.Mutex
}

// 来自客户端的消息
type Message struct {
	Action  string          `json:"action"`
	Topic   string          `json:"topic"`
	Message json.RawMessage `json:"message"`
}

type Subscription struct {
	Topic  string
	Client *Client
}

type HashSet []string

var SupportedSub HashSet = HashSet{PENDING, GAS, COIN}

func (sa HashSet) Contains(value string) bool {
	for _, a := range sa {
		if a == value {
			return true
		}
	}
	return false
}
func (sa *HashSet) Add(value string) bool {
	for _, a := range *sa {
		if a == value {
			return false
		}
	}
	*sa = append(*sa, value)
	return true
}
func (sa *HashSet) Remove(value string) bool {
	s := *sa
	if sa.Contains(value) {
		for id, sub := range s {
			if sub == value {
				// found this subscription from client and we do need remove it
				*sa = append(s[:id], s[id+1:]...)
				return true
			}
		}
	}
	return false
}
func (ps *PubSub) GetSubscribedTopics() []string {
	var topics HashSet
	for _, sub := range ps.Subscriptions {
		topics.Add(sub.Topic)
	}
	return topics
}
func (ps *PubSub) AddClient(client Client) *PubSub {
	var index = -1
	ps.lock_Clients.Lock()
	for i,v := range ps.Clients{
		if v.Id == client.Id{
			index = i
		}
	}
	if index >-1{//假如原来已经存在连接，则替换
		ps.Clients[index].IsConnect = false
		ps.Clients[index].Connection.Close()
		ps.Clients[index] = client
	}else{
		ps.Clients = append(ps.Clients, client)
	}
	ps.lock_Clients.Unlock()
	//fmt.Println("adding new client to the list", client.Id, len(ps.Clients))

	payload := []byte("Hello Client ID:" +
		client.Id)

	client.Connection.WriteMessage(1, payload)

	return ps

}

func (ps *PubSub) RemoveClient(client Client) *PubSub {

	// first remove all subscriptions by this client

	for index, sub := range ps.Subscriptions {

		if client.Id == sub.Client.Id {
			if index >= (len(ps.Subscriptions) - 1) {
				ps.Subscriptions = ps.Subscriptions[:len(ps.Subscriptions)-1]
			} else {
				ps.Subscriptions = append(ps.Subscriptions[:index], ps.Subscriptions[index+1:]...)
			}
		}
	}

	// remove client from the list
	ps.lock_Clients.Lock()
	for index, c := range ps.Clients {

		if c.Id == client.Id {
			if index >= (len(ps.Subscriptions) - 1) {
				ps.Clients = ps.Clients[:len(ps.Clients)-1]
			} else {
				ps.Clients = append(ps.Clients[:index], ps.Clients[index+1:]...)
			}
		}

	}
	ps.lock_Clients.Unlock()
	return ps
}

func (ps *PubSub) GetSubscriptions(topic string, client *Client) []Subscription {

	var subscriptionList []Subscription

	for _, subscription := range ps.Subscriptions {

		if client != nil {

			if subscription.Client.Id == client.Id && subscription.Topic == topic {
				subscriptionList = append(subscriptionList, subscription)
			}
		} else {

			if subscription.Topic == topic {
				subscriptionList = append(subscriptionList, subscription)
			}
		}
	}
	return subscriptionList
}

func (ps *PubSub) GetSubsTopics() HashSet {
	set := new(HashSet)

	for _, subscription := range ps.Subscriptions {
		set.Add(subscription.Topic)
	}
	return *set
}

func (ps *PubSub) Subscribe(client *Client, topic string) *PubSub {

	if !(SupportedSub.Contains(topic)) {
		return ps
	}
	clientSubs := ps.GetSubscriptions(topic, client)

	if len(clientSubs) > 0 {

		// client is subscribed this topic before

		return ps
	}

	newSubscription := Subscription{
		Topic:  topic,
		Client: client,
	}
	client.Topic.Add(topic)
	ps.Subscriptions = append(ps.Subscriptions, newSubscription)
	return ps
}

func (ps *PubSub) PublishOld(topic string, message []byte, excludeClient *Client) {

	subscriptions := ps.GetSubscriptions(topic, nil)

	for _, sub := range subscriptions {

		fmt.Printf("Sending to client id %s message is %s \n", sub.Client.Id, message)
		//sub.Client.Connection.WriteMessage(1, message)
		sub.Client.Send(message)
	}

}
func (ps *PubSub) PublishNew(client Client) {
	for {
		select {
		case m := <-client.MessageCh:
			client.Connection.WriteJSON(m)
			// client.Connection.WriteMessage(1, mm)
		}
	}
}

func (client *Client) Send(message []byte) error {

	return client.Connection.WriteMessage(1, message)

}

func (ps *PubSub) Unsubscribe(client *Client, topic string) *PubSub {

	//clientSubscriptions := ps.GetSubscriptions(topic, client)
	for index, sub := range ps.Subscriptions {

		if sub.Client.Id == client.Id && sub.Topic == topic {
			// found this subscription from client and we do need remove it
			ps.Subscriptions = append(ps.Subscriptions[:index], ps.Subscriptions[index+1:]...)
		}
	}
	client.Topic.Remove(topic)
	if len(client.Topic) < 1 {
		client.IsConnect = false
	}
	return ps

}
func (ps *PubSub) HandleReceiveMessage(client Client, messageType int, payload []byte) *PubSub {
	m := Message{}

	err := json.Unmarshal(payload, &m)
	if err != nil {
		fmt.Println("This is not correct message payload")
		return ps
	}
	switch m.Action {
	case PUBLISH:
		fmt.Println("This is publish new message")
		ps.PublishOld(m.Topic, m.Message, nil)

	case SUBSCRIBE:
		ps.Subscribe(&client, m.Topic)
		fmt.Println("new subscriber to topic", m.Topic, len(ps.Subscriptions), client.Id)

	case UNSUBSCRIBE:
		fmt.Println("Client unsubscribe the topic", m.Topic, client.Id)
		ps.Unsubscribe(&client, m.Topic)

	default:
		break
	}

	return ps
}