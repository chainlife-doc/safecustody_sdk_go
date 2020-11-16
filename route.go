package safecustody_sdk_go

//返回币种的信息体
//chain						string	链名
//coin						string	币名
//coin_precision			int		币的精度,也就是该币支持多少位小数
//min_deposit_amount		string	最小充值数量
//min_withdraw_amount		string	最小提币数量
//deposit_enabled			int		充值是否启用: 1=启用,0=未启用
//withdraw_enabled			int		提币是否启用: 1=启用,0=未启用
//deposit_confirm_count		int		充值入账确认数
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

//查询公共币种信息
//https://github.com/chainlife-doc/wallet-api/blob/master/%E6%9F%A5%E8%AF%A2%E5%B8%81%E7%A7%8D%E4%BF%A1%E6%81%AF.md
func (a *Api) QueryCoins() ([]QueryCoinConfBody, error) {

	p := struct {
		Auth auth `json:"auth"`
	}{
		a.getAuth(),
	}

	d := a.buildParam(p)

	var arr []QueryCoinConfBody

	err := a.request("info.php", d, &arr)

	return arr, err
}

//单币种信息查询
//coin 		string 币名
//https://github.com/chainlife-doc/wallet-api/blob/master/%E5%8D%95%E5%B8%81%E7%A7%8D%E4%BF%A1%E6%81%AF%E6%9F%A5%E8%AF%A2.md
func (a *Api) QueryCoinConf(coin string) ([]QueryCoinConfBody, error) {

	p := struct {
		Coin string `json:"coin"`
		Auth auth   `json:"auth"`
	}{
		coin,
		a.getAuth(),
	}
	d := a.buildParam(p)

	var arr []QueryCoinConfBody

	err := a.request("coinconf.php", d, &arr)

	return arr, err
}

//返回查询余额的包体
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

//余额查询的请求参数
//Coin 币种
//Chain 链名
type Coins struct {
	Coin  string `json:"coin"`
	Chain string `json:"chain"`
}

//查询余额
//coin []Coins
//https://github.com/chainlife-doc/wallet-api/blob/master/%E6%9F%A5%E8%AF%A2%E4%BD%99%E9%A2%9D.md
func (a *Api) QueryBalance(coins []Coins) ([]QueryBalanceBody, error) {

	p := struct {
		Coin []Coins `json:"coins"`
		Auth auth    `json:"auth"`
	}{
		coins,
		a.getAuth(),
	}

	d := a.buildParam(p)

	var arr []QueryBalanceBody

	err := a.request("balance.php", d, &arr)
	return arr, err
}

//获取充值地址信息体
//chain		string	主链
//coin		string	币名
//subuserid	string	调用端子账号，字符串，平台不管其含义
//addr		string	充币地址
//needmemo	int		0:不需要，1需要，填写的memo字符串格式: "userid#subuserid#id"（id是申请的appid序号）
type GetDepositAddrBody struct {
	Chain     string `json:"chain"`
	Coin      string `json:"coin"`
	Subuserid string `json:"subuserid"`
	Addr      string `json:"addr"`
	NeedMemo  string `json:"needmemo"`
}

//获取提笔地址的请求参数
//chain		string	主链
//coin		string	币名
//subuserid	string	调用端子账号，字符串，平台不管其含义
type AddrCoins struct {
	Coin      string `json:"coin"`
	Chain     string `json:"chain"`
	Subuserid string `json:"subuserid"`
}

type Addrs []AddrCoins

//获取充值地址
//https://github.com/chainlife-doc/wallet-api/blob/master/deposit/%E8%8E%B7%E5%8F%96%E5%85%85%E5%80%BC%E5%9C%B0%E5%9D%80.md
func (a *Api) GetDepositAddr(coins []AddrCoins) ([]GetDepositAddrBody, error) {

	p := struct {
		Coins Addrs `json:"coins"`
		Auth  auth  `json:"auth"`
	}{
		coins,
		a.getAuth(),
	}

	d := a.buildParam(p)

	var arr []GetDepositAddrBody

	err := a.request("deposit/addr.php", d, &arr)
	return arr, err
}

//获取充值记录的响应包体
//id		int		内部充值序号
//subuserid	string	调用端子账号，字符串，平台不管其含义
//chain		string	哪条主链上充值进来的
//coin		string	币名
//from_addr	string	订单发送地址
//addr		string	订单接收地址
//txid		string	交易ID
//amount	string	充值数量
//balance	string	充值后余额
//time		string	订单生成时间
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

//获取充值记录的请求参数
//subuserid	string	子账号，平台不管其含义（空字符串默认不做筛选）
//chain		string	主链 (空字符串默认不做筛选)
//coin		string	币名 (空字符串默认不做筛选)
//fromid	int	从哪个充值序号开始，值大于等于1,查询结果包含fromId对应的充值记录
//limit		int	最多查询多少条记录，包含fromid这条记录
type History struct {
	Subuserid string `json:"subuserid"`
	Chain     string `json:"chain"`
	Coin      string `json:"coin"`
	Fromid    int    `json:"fromid"`
	Limit     int    `json:"limit"`
}

//获取充值记录
//https://github.com/chainlife-doc/wallet-api/blob/master/deposit/%E8%8E%B7%E5%8F%96%E5%85%85%E5%80%BC%E8%AE%B0%E5%BD%95.md
func (a *Api) GetDepositHistory(h History) ([]GetDepositHistoryBody, error) {

	p := struct {
		History
		Auth auth `json:"auth"`
	}{
		h,
		a.getAuth(),
	}

	d := a.buildParam(p)

	var arr []GetDepositHistoryBody
	err := a.request("deposit/history.php", d, &arr)
	return arr, err
}

//内部地址查询响应参数
//exist	int	1：是内部地址，0：非内部地址
type queryIsInternalAddrBody struct {
	Exist int `json:"exist"`
}

//内部地址查询请求参数
//chain	string	主链
//coin	string	币名
//addr	string	地址
type QueryIsInternalAddr struct {
	Coin  string `json:"coin"`
	Chain string `json:"chain"`
	Addr  string `json:"addr"`
}

//内部地址查询
//返回的第一个参数true代表是内部地址,否则非内部地址
//https://github.com/chainlife-doc/wallet-api/blob/master/internal-addr/%E5%86%85%E9%83%A8%E5%9C%B0%E5%9D%80%E6%9F%A5%E8%AF%A2.md
func (a *Api) QueryIsInternalAddr(param QueryIsInternalAddr) (bool, error) {
	p := struct {
		QueryIsInternalAddr
		Auth auth `json:"auth"`
	}{
		param,
		a.getAuth(),
	}

	d := a.buildParam(p)

	q := &queryIsInternalAddrBody{}
	err := a.request("internal-addr/query.php", d, &q)
	var ok bool
	if q.Exist == 1 {
		ok = true
	} else {
		ok = false
	}
	return ok, err
}

//提交提币工单响应包体
//id			int		序号
//subuserid		string	调用端子账号，字符串，平台不管其含义
//chain			string	主链
//coin			string	币名
//from_addr		string	提币发送地址
//addr			string	提币接收地址
//amount		string	提币数量
//amount_sent	string	实际发送的提币数量
//memo			string	提币备注，比如用户ID之类的，可以是任意内容
//status		int		提币状态: 1=准备发送,2=发送中,3=发送成功,4=发送失败,5=发送已取消
//status_desc	string	状态描述
//txid			string	链上的交易ID
//usertags		string	用户标签
//time			string	订单创建时间
type SubmitWithdrawBody struct {
	Id         int64  `json:"id"`
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

//提交提币工单请求参数
//subuserid	string	调用端子账号，字符串，平台不管其含义
//chain		string	主链
//coin		string	币名
//addr		int	提币目标地址
//amount	float	提币数量
//memo		string	用户备注,内容自定义（会记录到区块链上）
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

//提交提币工单
//https://github.com/chainlife-doc/wallet-api/blob/master/withdraw/%E6%8F%90%E4%BA%A4%E6%8F%90%E5%B8%81%E5%B7%A5%E5%8D%95.md
func (a *Api) SubmitWithdraw(param SubmitWithdraw) (SubmitWithdrawBody, error) {

	p := struct {
		SubmitWithdraw
		Auth auth   `json:"auth"`
		Sign string `json:"sign"`
	}{
		param,
		a.getAuth(),
		a.WithdrawSign(param.Addr, param.Memo, param.Usertags),
	}
	d := a.buildParam(p)

	q := &SubmitWithdrawBody{}
	err := a.request("withdraw/submit.php", d, &q)
	return *q, err
}

//提币预校验接口
//https://github.com/chainlife-doc/wallet-api/blob/master/withdraw/%E6%8F%90%E5%B8%81%E9%A2%84%E6%A0%A1%E9%AA%8C%E6%8E%A5%E5%8F%A3.md
func (a *Api) ValidateWithdraw(param SubmitWithdraw) error {
	p := struct {
		SubmitWithdraw
		Auth auth   `json:"auth"`
		Sign string `json:"sign"`
	}{
		param,
		a.getAuth(),
		a.WithdrawSign(param.Addr, param.Memo, param.Usertags),
	}

	d := a.buildParam(p)
	err := a.request("withdraw/validator.php", d, nil)
	return err
}

//查询提币工单状态响应参数
//id			int	内部充值序号
//subuserid		string	调用端子账号，字符串，平台不管其含义
//chain			string	哪条主链上充值进来的
//coin			string	币名
//from_addr		string	提币发送地址
//addr			string	提币接收地址
//amount		string	充值数量
//amount_sent	string	实际发送的提币数量
//memo			string	提币备注，比如用户ID之类的，可以是任意内容
//status		int	提币状态: 1=准备发送,2=发送中,3=发送成功,4=发送失败,5=发送已取消
//status_desc	string	状态描述
//txid			string	链上的交易ID
//usertags		string	用户标签
//time			string	订单创建时间
type QueryWithdrawStatusBody struct {
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

//查询提币工单状态请求参数
//chain			string	主链
//coin			string	币名
//withdrawid	int	提币订单ID
type QueryWithdrawStatus struct {
	Coin       string `json:"coin"`
	Chain      string `json:"chain"`
	Withdrawid int    `json:"withdrawid"`
}

//查询提币工单状态
//https://github.com/chainlife-doc/wallet-api/blob/master/withdraw/%E6%9F%A5%E8%AF%A2%E6%8F%90%E5%B8%81%E5%B7%A5%E5%8D%95%E7%8A%B6%E6%80%81.md
func (a *Api) QueryWithdrawStatus(param QueryWithdrawStatus) (QueryWithdrawStatusBody, error) {

	p := struct {
		QueryWithdrawStatus
		Auth auth `json:"auth"`
	}{
		param,
		a.getAuth(),
	}

	d := a.buildParam(p)

	r := &QueryWithdrawStatusBody{}
	err := a.request("withdraw/status.php", d, &r)
	return *r, err
}

//查询提币记录响应参数
//id			int	内部充值序号
//subuserid		string	调用端子账号，字符串，平台不管其含义
//chain			string	哪条主链上充值进来的
//coin			string	币名
//from_addr		string	提币发送地址
//addr			string	提币接收地址
//amount		string	充值数量
//amount_sent	string	实际发送的提币数量
//memo			string	提币备注，比如用户ID之类的，可以是任意内容
//status		int	提币状态: 1=准备发送,2=发送中,3=发送成功,4=发送失败,5=发送已取消
//status_desc	string	状态描述
//txid			string	链上的交易ID
//usertags		string	用户标签
//time			string	订单创建时间
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

//查询提币记录请求参数
//subuserid		string	子账号，平台不管其含义（空字符串默认不做筛选）
//chain			string	主链 (空字符串默认不做筛选)
//coin			string	币名 (空字符串默认不做筛选)
//fromid		int	从哪个充值序号开始，值大于等于1,查询结果包含fromId对应的充值记录
//limit			int	最多查询多少条记录，包含fromid这条记录
type QueryWithdrawHistory struct {
	Subuserid string `json:"subuserid"`
	Chain     string `json:"chain"`
	Coin      string `json:"coin"`
	Fromid    int    `json:"fromid"`
	Limit     int    `json:"limit"`
}

//查询提币记录
//https://github.com/chainlife-doc/wallet-api/blob/master/withdraw/%E6%9F%A5%E8%AF%A2%E6%8F%90%E5%B8%81%E8%AE%B0%E5%BD%95.md
func (a *Api) QueryWithdrawHistory(param QueryWithdrawHistory) ([]QueryWithdrawHistoryBody, error) {
	p := struct {
		QueryWithdrawHistory
		Auth auth `json:"auth"`
	}{
		param,
		a.getAuth(),
	}

	d := a.buildParam(p)

	var arr []QueryWithdrawHistoryBody
	err := a.request("withdraw/history.php", d, &arr)
	return arr, err
}
