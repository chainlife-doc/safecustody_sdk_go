package main

import (
	"fmt"
	sdk "safecustody_sdk_go"
)

//这里是sdk参考使用方法
//这里示例代码有些错误的处理比较简单,仅供参考!!!
//嵌入到业务上面请严格判断错误和异常
func main() {
	api := new(sdk.Api)
	api.Host = ""
	api.SetUserInfo(
		"",   //appid
		"",   //salt
		"26", //userid
	)

	//单个币种查询
	r1, err := api.QueryCoinConf("usdt")
	fmt.Println(r1)
	fmt.Println(err)

	//查询公共币种信息
	r2, _ := api.QueryCoins()
	fmt.Print(r2)

	//查询余额
	var coins []sdk.Coins
	coins = append(coins, sdk.Coins{Coin: "usdt", Chain: "eth"})
	r3, _ := api.QueryBalance(coins)
	fmt.Println(r3)
	//
	//获取充值地址
	var addrs []sdk.AddrCoins
	addrs = append(addrs, sdk.AddrCoins{Coin: "btc", Chain: "btc", Subuserid: "1"})
	r4, err := api.GetDepositAddr(addrs)
	if err != nil {
		return
	}
	fmt.Print(r4)

	//获取充值记录
	r5, err := api.GetDepositHistory(sdk.History{
		Subuserid: "",
		Chain:     "",
		Coin:      "",
		Fromid:    0,
		Limit:     100,
	})
	fmt.Println(err)
	fmt.Println(r5)

	//内部地址查询
	r6, err := api.QueryIsInternalAddr(sdk.QueryIsInternalAddr{Coin: "trx", Chain: "trx", Addr: ""})
	fmt.Println(r6)
	fmt.Println(err)

	//提交提币工单
	r7, err := api.SubmitWithdraw(sdk.SubmitWithdraw{
		Subuserid: "26",
		Chain:     "trx",
		Coin:      "trx",
		Addr:      "",
		Amount:    0.01,
		Memo:      "xxx",
		Usertags:  "123",
	})
	fmt.Println(r7)
	fmt.Println(err)

	//提币预校验
	err11 := api.ValidateWithdraw(sdk.SubmitWithdraw{
		Subuserid: "26",
		Chain:     "trx",
		Coin:      "trx",
		Addr:      "",
		Amount:    0,
		Memo:      "xxx",
		Usertags:  "123",
	})
	fmt.Println(err11)

	//查询工单状态
	r9, err := api.QueryWithdrawStatus(sdk.QueryWithdrawStatus{
		Coin:       "trx",
		Chain:      "trx",
		Withdrawid: 1186,
	})
	fmt.Println(r9)
	fmt.Println(err)

	//查询历史提币记录
	r10, err := api.QueryWithdrawHistory(sdk.QueryWithdrawHistory{
		Chain:     "trx",
		Coin:      "trx",
		Subuserid: "26",
		Fromid:    1,
		Limit:     1,
	})
	fmt.Println(r10)
	fmt.Println(err)
}
