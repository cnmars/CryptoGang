package jobs

import (
	"golddigger/common/wsserver"
	"golddigger/global"
	"net/http"
	"net/url"
	"time"

	coingecko "github.com/superoo7/go-gecko/v3"
	"go.uber.org/zap"
)
type CoinPrice struct{
	Coin string `json:"coin"`
	Price float32 `json:"price"`
}
func FetchEthPriceOnce()(*CoinPrice,error){
	var httpClient *http.Client
	if global.GD_CONFIG.System.Env == "develop" {
		u, _ := url.Parse("htttp://127.0.0.1:7890")
		t := &http.Transport{
			MaxIdleConns:    10,
			MaxConnsPerHost: 10,
			IdleConnTimeout: time.Duration(10) * time.Second,
			//Proxy: http.ProxyURL(url),
			Proxy: http.ProxyURL(u),
		}
		httpClient = &http.Client{
			Transport: t,
			Timeout:   time.Duration(10) * time.Second,
		}
	} else {
		httpClient = &http.Client{
			Timeout: time.Second * 100,
		}
	}
	CG := coingecko.NewClient(httpClient)
	ids := []string{"ethereum"}
	vc := []string{"usd"}
	sp, err := CG.SimplePrice(ids, vc)
	if err != nil {
		global.GD_LOG.Error("获取币价失败", zap.Error(err))
	}
	data := &CoinPrice{
		Coin:"eth",
		Price:(*sp)["ethereum"]["usd"],
	}
	return data,err
}
func FetchPrice() {
	var httpClient *http.Client
	if global.GD_CONFIG.System.Env == "develop" {
		u, _ := url.Parse("htttp://127.0.0.1:7890")
		t := &http.Transport{
			MaxIdleConns:    10,
			MaxConnsPerHost: 10,
			IdleConnTimeout: time.Duration(10) * time.Second,
			//Proxy: http.ProxyURL(url),
			Proxy: http.ProxyURL(u),
		}
		httpClient = &http.Client{
			Transport: t,
			Timeout:   time.Duration(10) * time.Second,
		}
	} else {
		httpClient = &http.Client{
			Timeout: time.Second * 100,
		}
	}
	CG := coingecko.NewClient(httpClient)
	// list, _ := CG.CoinsList()
	// coin, err := CG.CoinsID("dogecoin", false, false, false, false, false, false)
	ids := []string{"bitcoin", "ethereum"}
	vc := []string{"usd", "cny"}
	// fmt.Println(sp)
	for !global.GD_WS.IsStop {
		sp, err := CG.SimplePrice(ids, vc)
		if err != nil {
			global.GD_LOG.Error("获取币价失败", zap.Error(err))
		}
		m := wsserver.WsPriceData{
			Data: sp,
		}
		global.GD_WS.ThirdPartData <- m.ToCommonData()
		time.Sleep(time.Second * 60)
	}
}
