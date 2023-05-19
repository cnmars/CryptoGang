import { useAppStore } from '@/pinia/modules/app'
// import { httpProvider } from "../common/httpProvider";


export class GdWebsocket{
   websocket :  WebSocket
   PendingDataArr: Array<Object>
   GasDataObj:Object
   CoinPriceObj: Object
	constructor() {
		this.websocket = null;	// websocket对象
		this.PendingDataArr = new Array<Object>()
		this.CoinPriceObj =new Object()
		this.GasDataObj ={ethereum:{cny:0,usd:0}}

	}

	public init = (url:string)=> {
		if (!window.WebSocket) {
            return console.log('您的浏览器不支持WebSocket');
          }

          try {
			let token=localStorage.getItem("token")
            this.websocket = new WebSocket(url,[token]);
            
          } catch (error) {
            console.log(error);
            
          }
		if (this.websocket) {
			this.websocket.onopen = this.WSonOpen
			this.websocket.onmessage = this.WSonMessage
			this.websocket.onerror = this.WSonError
			this.websocket.onclose = this.WSonClosed
			// this.websocket.send = this.sendMsg
			this.websocket.close = this.closeSocket
			
		}else{
            console.log('连接异常');
        }
	}

	// websocket连接成功后的回调函数。
	public WSonOpen = () =>{
		this.Subscribe('gas')
		this.Subscribe('coin')
		console.log('ws连接成功');
	}

	// websocket从服务器接受到信息时的回调函数。
	public WSonMessage = (event) =>{
		try {
			const{FillCoinPriceData,FillGasData} =useAppStore()
			let data = JSON.parse(event.data)
			switch (data.type) {
				case 'coin':
					if(data.data.data){
						console.log(data.data);
						this.CoinPriceObj = data.data.data
						FillCoinPriceData(data.data.data)
					}
					break;
				case 'gas':
					if(data.data.data){
						console.log(JSON.parse(data.data.data));
						let d = this.dealGasData(JSON.parse(data.data.data))
						FillGasData(d)
						// this.GasDataObj = data.data.data
					}
					break;
				case 'pending':
					
					break;
			
				default:
					break;
			}
		} catch (error) {
			console.log(event.data);
		}
	}

	// websocket连接失败后的回调函数。
	public WSonError=(event)=> {
			console.log('websocket error ', event);
	}

	// websocket连接关闭后的回调函数。
	public WSonClosed=(event)=> {
			console.log('websocket closed ', event);

	}

	/**
	 * WebSocket的属性值readyState，取下面4个状态值：
	 * CONNECTING: 0 连接中 、OPEN: 1 已连接 、CLOSING: 2 关闭中 、CLOSED: 3 已关闭
	 * 直接判断WebSocket对象的属性值readyState，如果是1的话就表示连接可用。
	 */
	public Subscribe =(topic:string)=> {
		if (this.websocket.readyState === 1) {
			switch (topic) {
				case "coin":
					this.websocket.send('{"action":"subscribe","topic":"coin"}')
					break;
				case "gas":
					this.websocket.send('{"action":"subscribe","topic":"gas"}')
					break;
				case "pending":
					this.websocket.send('{"action":"subscribe","topic":"pending"}')
					break;
			
				default:
					break;
			}
		}else{
			console.log('ws连接异常');
		}
	}
    public UnSubscribe=(topic:string) =>{
		if (this.websocket.readyState === 1) {
			switch (topic) {
				case "coin":
					this.websocket.send('{"action":"unsubscribe","topic":"coin"}')
					break;
				case "gas":
					this.websocket.send('{"action":"unsubscribe","topic":"gas"}')
					break;
				case "pending":
					this.websocket.send('{"action":"unsubscribe","topic":"pending"}')
					break;
			
				default:
					break;
			}
		}else{
			console.log('ws连接异常');
		}
	}
	private dealGasData =(GasData:any)=>{
		let data ={}
		if(GasData.blockPrices){
			let bp = GasData.blockPrices[0]
			let a =  (GasData as any).blockPrices[0].baseFeePerGas.toString().slice(0,2)
			if(a.indexOf('.')>-1){
			a =  a.slice(0,1)
			}
			data['baseFee'] = a
			let epArr =  bp.estimatedPrices
			data['fast'] = epArr[0]
			data['medium'] = epArr[1]
			data['slow'] = epArr[2]
		}else{
			data['baseFee'] ='13'
		}
		return data
	  }




    // 主动关闭websocket连接
	closeSocket() {
		this.websocket.close()
		this.websocket = null
		console.log("websocket closed");
	}

}