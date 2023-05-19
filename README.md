# CryptoGang

> 一个包含前后端的链上数据监控和玩链上生态的工具集，开发初期起名叫GoldDigger，所以项目中有些地方名字还是叫这个没有去改

## 技术架构

* 前端：

**vite + vue3 + ts+pinia + elementui-plus + axios**

* 后端

**golang + viper + gorm + redis + jwt + casbin + gorllia websocket + cobra**

## 部分运行截图

![image-20230519104307310](https://raw.githubusercontent.com/selfmakeit/resource/main/image-20230519104307310.png)

![image-20230519104344336](https://raw.githubusercontent.com/selfmakeit/resource/main/image-20230519104344336.png)

![image-20230519104442166](https://raw.githubusercontent.com/selfmakeit/resource/main/image-20230519104442166.png)





## 功能涵盖

**NFT板块功能包含：**

1. Live Mint(链上实时Mint的pending数据)
2. Hash Mint：根据别人的交易hash快速跟单
3. Quick Sell：快速挂单出货
4. Quick Offer：快速批量发Offer
5. Sweep Floor：快速扫地板
6. Follow Mint：Mint自动跟单
7. 后续会不断推出ERC20相关新功能

**ERC20板块功能包含：**

1. Hot Token：当前热门代币
2. Big Flow：大资金流向
3. Token Analysis：代币分析，包含代币基本信息、持币分布、近期大量交易等
4. Wallet Analysis：钱包地址分析，包含资产分布、历史盈利，NFT持有情况等
5. 后续会不断推出NFT相关新功能

**常用工具**：

1. 抢跑工具：按照预先设定抢跑交易
2. 资产管理：NFT、代币的分发、归集

