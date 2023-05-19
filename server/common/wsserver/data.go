package wsserver

import (
	// "fmt"
	"encoding/json"
)

// 没有用这个接口
type SerilizeData interface {
	Marshal() ([]byte, error)
	// Unmarshal(data []byte) error
	ToCommonData() WsMessage
	GetType() int32
}

var (
	PENDING = "pending"
	GAS     = "gas"
	COIN    = "coin"
)

type WsMessage struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

// pending数据通道
type WsPendingData struct {
	ProjectName string `json:"projectName"`
	Icon        string `json:"icon"`
}

// gas数据通道
type WsGasData struct {
	Data string `json:"data"`
	// Level1 string `json:"level1"`
}

// coin price数据通道
type WsPriceData struct {
	Data *map[string]map[string]float32 `json:"data"`
}

// 转换为向客户端消息通道写入的数据
func (wd *WsPendingData) ToCommonData() (ws WsMessage) {
	ws.Data = *wd
	ws.Type = PENDING
	return
}
func (wd *WsPendingData) Marshal() ([]byte, error) {
	return json.Marshal(wd)
}
func (wp *WsPriceData) ToCommonData() (ws WsMessage) {
	ws.Data = *wp
	ws.Type = COIN
	return
}
func (wp *WsPriceData) Marshal() ([]byte, error) {
	return json.Marshal(wp)
}

func (wg *WsGasData) ToCommonData() (ws WsMessage) {
	ws.Data = wg
	ws.Type = GAS
	return
}
func (wg *WsGasData) Marshal() ([]byte, error) {
	return json.Marshal(wg)
}

func (m *WsMessage) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

func (m *WsMessage) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}

// GetType 获取状态信息
func (m *WsMessage) GetType() string {
	return m.Type
}
