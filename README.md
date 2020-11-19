# 赛福托管钱包API GO-SDK


### 安装SDK

#### mod安装
- 我们推荐mod安装 ,在项目的根目录创建[mod](https://studygolang.com/articles/20716)文件  
    ```go mod init project```
- 创建main.go,在代码里面导入api包
    ```go
    import sdk "github.com/chainlife-doc/safecustody_sdk_go"
    ```   
- 在go.mod目录下,命令行输入  
    ```go mod tidy```
   
- 具体操作  
  ```
  +---------------------------------------------------------------------------------------------+
  |~:cmd> mkdir project                                                                         |
  |~:cmd> cd project                                                                            |
  |./project:cmd> go mod init project                                                           |
  |./project:cmd> touch main.go                                                                 |
  |./project:cmd> echo "package main" > main.go                                                 |
  |./project:cmd> echo 'import sdk "github.com/chainlife-doc/safecustody_sdk_go"' >> main.go    |
  |./project:cmd> go mod tidy                                                                   |
  |./project:cmd>                                                                               |
  +──-------------------------------------------------------------------------------------------+ 
    ```        
#### 源码安装 
    
- 直接从GitHup下载源码,把整个`safecustody_sdk_go`项目放入您的项目目录中,  
    然后在代码里`import sdk "safecustody_sdk_go"`     
    
    ```
    Project //项目
      ├── main.go //代码中写入 import sdk "safecustody_sdk_go"
      ├── ...
      └── safecustody_sdk_go  //safecustody_sdk_go与main.go同级
          ├── sdk.go       
          └── ...   
    ```
# 例子

#### 创建sdkApi`mod方式`  
 ```
    package main
    
    import sdk "github.com/chainlife-doc/safecustody_sdk_go"
    
    api := new(sdk.Api)
    api.Host = "https://www.xxxx.com/"  //api域名
    api.SetUserInfo(
        "", //appid
        "", //salt
        "", //userid
     )       

``` 

#### 创建sdkApi`源码方式`  
 ```
    package main
    
    import sdk "safecustody_sdk_go"
    
    api := new(sdk.Api)
    api.Host = "https://www.xxxx.com/"  //api域名
    api.SetUserInfo(
        "", //appid
        "", //salt
        "", //userid
     )       

``` 

#### 单个币种查询
```go
r1, err := api.QueryCoinConf("btc")
```

#### 查询公共币种信息
```go
r2, _ := api.QueryCoins(0)
```

#### 查询余额
```go
var coins []sdk.Coins
coins = append(coins, sdk.Coins{Coin: "btc", Chain: "btc"})
r3, _ := api.QueryBalance(coins)
```

#### 获取充值地址
```go
var addrs []sdk.AddrCoins
addrs = append(addrs, sdk.AddrCoins{Coin: "btc", Chain: "btc", Subuserid: "1"})
r4, err := api.GetDepositAddr(addrs)
```

#### 获取充值记录
```go
r5, err := api.GetDepositHistory(sdk.History{
		Subuserid: "",
		Chain:     "eth",
		Coin:      "eth",
		Fromid:    0,
		Limit:     100,
	})
```

#### 内部地址查询
```go
r6, err := api.QueryIsInternalAddr(sdk.QueryIsInternalAddr{Coin: "eth", Chain: "eth", Addr: "3297f672db8afa3"})
```

#### 提交提币工单
```go
r7, err := api.SubmitWithdraw(sdk.SubmitWithdraw{
		Subuserid: "1",
		Chain:     "btc",
		Coin:      "btc",
		Addr:      "",
		Amount:    0.01,
		Memo:      "xxx",
		Usertags:  "123",
	})
```

#### 提币预校验
```go
err11 := api.ValidatorWithdraw(sdk.SubmitWithdraw{
		Subuserid: "1",
		Chain:     "btc",
		Coin:      "btc",
		Addr:      "dddddd",
		Amount:    0.01,
		Memo:      "xxx",
		Usertags:  "123",
	})
```

#### 查询工单状态
```go
r9, err := api.QueryWithdrawStatus(sdk.QueryWithdrawStatus{
		Coin:       "btc",
		Chain:      "btc",
		Withdrawid: 1,
	})
```

#### 查询历史提币记录
```go
r10, err := api.QueryWithdrawHistory(sdk.QueryWithdrawHistory{
		Chain:     "btc",
		Coin:      "btc",
		Subuserid: "1",
		Fromid:    1,
		Limit:     1,
	})
```

# Api接口
#### 单个币种查询  

- Response
    ```go
    //chain					string	链名
    //coin					string	币名
    //coin_precision			        int		币的精度,也就是该币支持多少位小数
    //min_deposit_amount		        string	最小充值数量
    //min_withdraw_amount		        string	最小提币数量
    //deposit_enabled			        int		充值是否启用: 1=启用,0=未启用
    //withdraw_enabled			        int		提币是否启用: 1=启用,0=未启用
    //deposit_confirm_count		        int		充值入账确认数
    //need_memo					int		充值是否需要备注: 1=充值需要备注,0=充值不需要备注
    type QueryCoinConfBody struct {
    	Chain               string `json:"chain"`
    	Coin                string `json:"coin"`
    	CoinPrecision       int    `json:"coin_precision"`
    	MinDepositAmount    string `json:"min_deposit_amount"`
    	MinWithdrawAmount   string `json:"min_withdraw_amount"`
    	DepositEnabled      int    `json:"deposit_enabled"`
    	WithdrawEnabled     int    `json:"withdraw_enabled"`
    	DepositConfirmCount int    `json:"deposit_confirm_count"`
    	NeedMemo            int    `json:"need_memo"`
    }
    ```

- Request
     ```
    //coin 		string 币名
    ```
- Function
    ```go
     func (a *Api) QueryCoinConf(coin string) ([]QueryCoinConfBody, error)
    ```  
#### 查询公共币种信息
- Response 
    ```go
    //chain					string	链名
    //coin					string	币名
    //coin_precision			        int	币的精度,也就是该币支持多少位小数
    //min_deposit_amount		        string	最小充值数量
    //min_withdraw_amount		        string	最小提币数量
    //deposit_enabled			        int	充值是否启用: 1=启用,0=未启用
    //withdraw_enabled			        int	提币是否启用: 1=启用,0=未启用
    //deposit_confirm_count		        int	充值入账确认数
    //need_memo					int	充值是否需要备注: 1=充值需要备注,0=充值不需要备注
    type QueryCoinConfBody struct {
        Chain               string `json:"chain"`
        Coin                string `json:"coin"`
        CoinPrecision       int    `json:"coin_precision"`
        MinDepositAmount    string `json:"min_deposit_amount"`
        MinWithdrawAmount   string `json:"min_withdraw_amount"`
        DepositEnabled      int    `json:"deposit_enabled"`
        WithdrawEnabled     int    `json:"withdraw_enabled"`
        DepositConfirmCount int    `json:"deposit_confirm_count"`
        NeedMemo            int    `json:"need_memo"`
        }
    ```
- Function  
    ```go
    func (a *Api) QueryCoins() ([]QueryCoinConfBody, error)
    ```

#### 查询余额
 - Response
     ```go
    //chain	string	主链
    //coin	string	币名
    //balance	string	余额数量
    //as_cny	string	余额以cnc为单位表示的数量
    type QueryBalanceBody struct {
        Chain   string `json:"chain"`
        Coin    string `json:"coin"`
        Balance string `json:"balance"`
        AsCny   string `json:"as_cny"`
    }
    ```
- Request
    ```go
    //Coin 币种
    //Chain 链名
    type Coins struct {
        Coin  string `json:"coin"`
        Chain string `json:"chain"`
    }
    ```
- Function
    ```go
    func (a *Api) QueryBalance(coins []Coins) ([]QueryBalanceBody, error)
    ````
#### 获取充值地址
 - Response
     ```go
    //chain	string	主链
    //coin	string	币名
    //subuserid	string	调用端子账号，字符串，平台不管其含义
    //addr	string	充币地址
    //needmemo	int	0:不需要，1需要，填写的memo字符串格式: "userid#subuserid#id"（id是申请的appid序号）
    type GetDepositAddrBody struct {
    	Chain     string `json:"chain"`
    	Coin      string `json:"coin"`
    	Subuserid string `json:"subuserid"`
    	Addr      string `json:"addr"`
    	NeedMemo  string `json:"needmemo"`
    }
    ```
- request
    ```go
    //chain	string	主链
    //coin	string	币名
    //subuserid	string	调用端子账号，字符串，平台不管其含义
    type AddrCoins struct {
    	Coin      string `json:"coin"`
    	Chain     string `json:"chain"`
    	Subuserid string `json:"subuserid"`
    }

    ```
- function
    ```go
    func (a *Api) GetDepositAddr(coins []AddrCoins) ([]GetDepositAddrBody, error)
    ```

#### 获取充值记录
- Response
    ```go
    //id	int	内部充值序号
    //subuserid	string	调用端子账号，字符串，平台不管其含义
    //chain	string	哪条主链上充值进来的
    //coin	string	币名
    //from_addr	string	订单发送地址
    //addr	string	订单接收地址
    //txid	string	交易ID
    //amount	string	充值数量
    //balance	string	充值后余额
    //time	string	订单生成时间
    type GetDepositHistoryBody struct {
        Id        int64  `json:"id"`
        Subuserid string `json:"subuserid"`
        Chain     string `json:"chain"`
        Coin      string `json:"coin"`
        FromAddr  string `json:"from_addr"`
        Addr      string `json:"addr"`
        Txid      string `json:"txid"`
        Amount    string `json:"amount"`
        Balance   string `json:"balance"`
        Time      string `json:"time"`
    }
    ```
- Request
    ```go
    //subuserid	string	子账号，平台不管其含义（空字符串默认不做筛选）
    //chain	string	主链 (空字符串默认不做筛选)
    //coin	string	币名 (空字符串默认不做筛选)
    //fromid	int	从哪个充值序号开始，值大于等于1,查询结果包含fromId对应的充值记录
    //limit	int	最多查询多少条记录，包含fromid这条记录
    type History struct {
        Subuserid string `json:"subuserid"`
        Chain     string `json:"chain"`
        Coin      string `json:"coin"`
        Fromid    int    `json:"fromid"`
        Limit     int    `json:"limit"`
    }
    ```  
- Function
    ```go
    func (a *Api) GetDepositHistory(h History) ([]GetDepositHistoryBody, error)
    ```

#### 内部地址查询
- Request
    ```go
    //chain	string	主链
    //coin	string	币名
    //addr	string	地址
    type QueryIsInternalAddr struct {
        Coin  string `json:"coin"`
        Chain string `json:"chain"`
        Addr  string `json:"addr"`
    }
    ```
- Function
    ```go
    //返回的第一个参数true代表是内部地址,否则非内部地址
    func (a *Api) QueryIsInternalAddr(param QueryIsInternalAddr)(bool, error)
    ```


#### 提交提币工单
- Response
    ```go
    //id		int		序号
    //subuserid		string	调用端子账号，字符串，平台不管其含义
    //chain		string	主链
    //coin		string	币名
    //from_addr		string	提币发送地址
    //addr		string	提币接收地址
    //amount		string	提币数量
    //amount_sent	string	实际发送的提币数量
    //memo		string	提币备注，比如用户ID之类的，可以是任意内容
    //status		int	提币状态: 1=准备发送,2=发送中,3=发送成功,4=发送失败,5=发送已取消
    //status_desc	string	状态描述
    //txid		string	链上的交易ID
    //fee_coin          string  手续费币种
    //fee_coin_chain 	string  手续费币种所在链
    //fee_amount    	string  手续费数量
    //usertags		string	用户标签
    //time		string	订单创建时间
    type SubmitWithdrawBody struct {
        Id            int64  `json:"id"`
        Subuserid     string `json:"subuserid"`
        Chain         string `json:"chain"`
        Coin          string `json:"coin"`
        FromAddr      string `json:"from_addr"`
        Addr          string `json:"addr"`
        Amount        string `json:"amount"`
        AmountSent    string `json:"amount_sent"`
        Memo          string `json:"memo"`
        Status        int    `json:"status"`
        StatusDesc    string `json:"status_desc"`
        FeeCoin       string `json:"fee_coin"`
        FeeCoinChain  string `json:"fee_coin_chain"`
        FeeAmount     string `json:"fee_amount"`
        Txid          string `json:"txid"`
        Usertags      string `json:"usertags"`
        Time          string `json:"time"`
    }
    ```
- Request  
    ````go
    //subuserid	string	调用端子账号，字符串，平台不管其含义
    //chain	string	主链
    //coin	string	币名
    //addr	int	提币目标地址
    //amount	float	提币数量
    //memo	string	用户备注,内容自定义（会记录到区块链上）
    //usertags	string	用户标签，内容自定义 （不会记录到区块链上）
    type SubmitWithdraw struct {
        Subuserid string  `json:"subuserid"`
        Chain     string  `json:"chain"`
        Coin      string  `json:"coin"`
        Addr      string  `json:"addr"`
        Amount    float64 `json:"amount"`
        Memo      string  `json:"memo"`
        Usertags  string  `json:"usertags"`
    }
    ````
- Function
    ```go
    func (a *Api) SubmitWithdraw(param SubmitWithdraw) (SubmitWithdrawBody, error)
    ```
#### 提币预校验接口
- Request
    ```go
    //subuserid	string	调用端子账号，字符串，平台不管其含义
    //chain	string	主链
    //coin	string	币名
    //addr	int	提币目标地址
    //amount	float	提币数量
    //memo	string	用户备注,内容自定义（会记录到区块链上）
    //usertags	string	用户标签，内容自定义 （不会记录到区块链上）
    type SubmitWithdraw struct {
        Subuserid string  `json:"subuserid"`
        Chain     string  `json:"chain"`
        Coin      string  `json:"coin"`
        Addr      string  `json:"addr"`
        Amount    float64 `json:"amount"`
        Memo      string  `json:"memo"`
        Usertags  string  `json:"usertags"`
    } 
    ```
- Function
    ```go
    func (a *Api) ValidateWithdraw(param SubmitWithdraw) error
    ```
#### 查询提币工单状态
- Response
    ```go
    //id		int	内部充值序号
    //subuserid		string	调用端子账号，字符串，平台不管其含义
    //chain		string	哪条主链上充值进来的
    //coin		string	币名
    //from_addr		string	提币发送地址
    //addr		string	提币接收地址
    //amount		string	充值数量
    //amount_sent	string	实际发送的提币数量
    //memo		string	提币备注，比如用户ID之类的，可以是任意内容
    //status		int	提币状态: 1=准备发送,2=发送中,3=发送成功,4=发送失败,5=发送已取消
    //status_desc	string	状态描述
    //txid		string	链上的交易ID
    //usertags		string	用户标签
    //time		string	订单创建时间
    type QueryWithdrawStatusBody struct {
        Id         int    `json:"id"`
        Subuserid  string `json:"subuserid"`
        Chain      string `json:"chain"`
        Coin       string `json:"coin"`
        FromAddr   string `json:"from_addr"`
        Addr       string `json:"addr"`
        Amount     string `json:"amount"`
        AmountSent string `json:"amount_sent"`
        SubBalance string `json:"sub_balance"`
        Memo       string `json:"memo"`
        Status     int    `json:"status"`
        StatusDesc string `json:"status_desc"`
        Txid       string `json:"txid"`
        Usertags   string `json:"usertags"`
        Time       string `json:"time"`
    }
    ```
- Request  
    ```go
    //chain			string	主链
    //coin			string	币名
    //withdrawid	        int	提币订单ID
    type QueryWithdrawStatus struct {
        Coin       string `json:"coin"`
        Chain      string `json:"chain"`
        Withdrawid int    `json:"withdrawid"`
    }    
    ```
- Function
    ```go
    func (a *Api) QueryWithdrawStatus(param QueryWithdrawStatus) (QueryWithdrawStatusBody, error)
    ```
#### 查询提币记录
- Response
    ```go
    //id		int	内部充值序号
    //subuserid		string	调用端子账号，字符串，平台不管其含义
    //chain		string	哪条主链上充值进来的
    //coin		string	币名
    //from_addr		string	提币发送地址
    //addr		string	提币接收地址
    //amount		string	充值数量
    //amount_sent	string	实际发送的提币数量
    //memo		string	提币备注，比如用户ID之类的，可以是任意内容
    //status		int	提币状态: 1=准备发送,2=发送中,3=发送成功,4=发送失败,5=发送已取消
    //status_desc	string	状态描述
    //txid		string	链上的交易ID
    //usertags		string	用户标签
    //time		string	订单创建时间
    type QueryWithdrawHistoryBody struct {
        Id         int    `json:"id"`
        Subuserid  string `json:"subuserid"`
        Chain      string `json:"chain"`
        Coin       string `json:"coin"`
        FromAddr   string `json:"from_addr"`
        Addr       string `json:"addr"`
        Amount     string `json:"amount"`
        AmountSent string `json:"amount_sent"`
        Memo       string `json:"memo"`
        Status     int    `json:"status"`
        StatusDesc string `json:"status_desc"`
        Txid       string `json:"txid"`
        Usertags   string `json:"usertags"`
        Time       string `json:"time"`
    }
    ```
- Request
    ```go
    //subuserid		string	子账号，平台不管其含义（空字符串默认不做筛选）
    //chain		string	主链 (空字符串默认不做筛选)
    //coin		string	币名 (空字符串默认不做筛选)
    //fromid		int	从哪个充值序号开始，值大于等于1,查询结果包含fromId对应的充值记录
    //limit		int	最多查询多少条记录，包含fromid这条记录
    type QueryWithdrawHistory struct {
        Subuserid string `json:"subuserid"`
        Chain     string `json:"chain"`
        Coin      string `json:"coin"`
        Fromid    int    `json:"fromid"`
        Limit     int    `json:"limit"`
    }
    ```
- Function
    ```go
    Func (a *Api) QueryWithdrawHistory(param QueryWithdrawHistory) (QueryWithdrawHistoryBody, error)
    ```
