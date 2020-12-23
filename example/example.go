package main

import (
	"fmt"
	"log"
	sdk "safecustody_sdk_go"
)

//这里是sdk参考使用方法
//这里示例代码有些错误的处理比较简单,仅供参考!!!
//嵌入到业务上面请严格判断错误和异常
//使用sdk,提币不需要主动签名(sign字段)加密,内部已做处理
//使用sdk,验证身份(token字段)不需要主动签名加密,内部已做处理
//使用案例请认真阅读开发文档,因为有些字段是选填的,案例中并没有体现出来
func main() {

	//创建api
	api := new(sdk.Api)

	//TODO 请向微信群里面的官方人员获取
	api.Host = ""

	//api访问公钥
	api.ApiKey = ""

	//设置用户信息
	api.SetUserInfo(
		"", //对应商户后台的APPID
		"", //对应商户后台的SECRETKEY
		"", //userid 对应商户后台的商户ID
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
		Fromid:    0,   //从哪个充值序号开始，值大于等于1,查询结果包含fromId对应的充值记录
		Limit:     100, //最多查询多少条记录，包含Fromid这条记录
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
		Memo:      "xxx", //该字段主要提供给链上支持备注的币种，内容会更新到链上
		Usertags:  "123", //用户标签, 自定义内容，一般作为订单备注使用,辅助说明
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
		Memo:      "xxx", //该字段主要提供给链上支持备注的币种，内容会更新到链上
		Usertags:  "123", //用户标签, 自定义内容，一般作为订单备注使用,辅助说明
	})
	fmt.Println(err11)

	//查询工单状态
	r9, err := api.QueryWithdrawStatus(sdk.QueryWithdrawStatus{
		Coin:       "trx",
		Chain:      "trx",
		Withdrawid: 1186, //提币订单ID
	})
	fmt.Println(r9)
	fmt.Println(err)

	//查询历史提币记录
	r10, err := api.QueryWithdrawHistory(sdk.QueryWithdrawHistory{
		Chain:     "trx",
		Coin:      "trx",
		Subuserid: "26",
		Fromid:    1, //从哪个充值序号开始，值大于等于1,查询结果包含fromId对应的充值记录
		Limit:     1, //最多查询多少条记录，包含Fromid这条记录
	})
	fmt.Println(r10)
	fmt.Println(err)

	//取消提币工单
	err = api.WithdrawCancel(sdk.WithdrawCancel{
		Subuserid:  "",
		Chain:      "",
		Coin:       "",
		Withdrawid: 0,
	})
	if err != nil {
		log.Fatalln("取消失败" + err.Error())
	}
}
