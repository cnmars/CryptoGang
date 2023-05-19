package jobs

import (
	"golddigger/common/wsserver"
	"golddigger/global"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"go.uber.org/zap"
)

func FetchGas() {
	var client *http.Client
	if global.GD_CONFIG.System.Env == "develop" {
		u, _ := url.Parse("htttp://127.0.0.1:7890")
		t := &http.Transport{
			MaxIdleConns:    10,
			MaxConnsPerHost: 10,
			IdleConnTimeout: time.Duration(10) * time.Second,
			//Proxy: http.ProxyURL(url),
			Proxy: http.ProxyURL(u),
		}
		client = &http.Client{
			Transport: t,
			Timeout:   time.Duration(10) * time.Second,
		}

	} else {
		client = &http.Client{}
	}
	//生成要访问的url
	url := "https://api.blocknative.com/gasprices/blockprices"
	//提交请求
	reqest, err := http.NewRequest("GET", url, nil)
	reqest.Header.Add("Content-Type", "application/json")
	//增加header选项
	// reqest.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:107.0) Gecko/20100101 Firefox/107.0")
	reqest.Header.Add("Authorization", "f41e1267-e2ee-4ff4-bcff-324rdsfs")

	if err != nil {
		panic(err)
	}
	//处理返回结果
	for !global.GD_WS.IsStop {
		response, err := client.Do(reqest)
		// fmt.Println(response)
		if err == nil {
			defer response.Body.Close()
			body, err := ioutil.ReadAll(response.Body)

			str := string(body)
			if err != nil {
				global.GD_LOG.Error("获取Gas失败", zap.Error(err))
			}
			m := wsserver.WsGasData{
				Data: str,
			}
			global.GD_WS.ThirdPartData <- m.ToCommonData()
			time.Sleep(time.Second * 10)
		}
	}

}
